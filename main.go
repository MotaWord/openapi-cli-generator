package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	"github.com/motaword/openapi-cli-generator/shorthand"
)

//go:generate go-bindata ./templates/...

// OpenAPI Extensions
const (
	ExtAliases     = "x-cli-aliases"
	ExtDescription = "x-cli-description"
	ExtIgnore      = "x-cli-ignore"
	ExtHidden      = "x-cli-hidden"
	ExtName        = "x-cli-name"
	ExtCmdGroups   = "x-cli-cmd-groups"
	ExtCmdGroup    = "x-cli-cmd-group"
	ExtWaiters     = "x-cli-waiters"
)

// Param describes an OpenAPI parameter (path, query, header, etc)
type Param struct {
	Name        string
	CLIName     string
	GoName      string
	Description string
	In          string
	Required    bool
	Type        string
	TypeNil     string
	Style       string
	Explode     bool
}

// Operation describes an OpenAPI operation (GET/POST/PUT/PATCH/DELETE)
type Operation struct {
	HandlerName    string
	GoName         string
	Use            string
	Aliases        []string
	Short          string
	Long           string
	Method         string
	CanHaveBody    bool
	ReturnType     string
	Path           string
	AllParams      []*Param
	RequiredParams []*Param
	OptionalParams []*Param
	MediaType      string
	Examples       []string
	Hidden         bool
	NeedsResponse  bool
	Group          *OperationGroup
	Waiters        []*WaiterParams
}

// OperationGroup describes a grouping of operations to be nested under a same level.
type OperationGroup struct {
	GoName string
	Name   string
	Short  string
	Long   string
	Parent *OperationGroup
}

// Waiter describes a special command that blocks until a condition has been
// met, after which it exits.
type Waiter struct {
	CLIName     string
	GoName      string
	Use         string
	Aliases     []string
	Short       string
	Long        string
	Delay       int
	Attempts    int
	OperationID string `json:"operationId"`
	Operation   *Operation
	Matchers    []*Matcher
	After       map[string]map[string]string
}

// Matcher describes a condition to match for a waiter.
type Matcher struct {
	Select   string
	Test     string
	Expected json.RawMessage
	State    string
}

// WaiterParams links a waiter with param selector querires to perform wait
// operations after a command has run.
type WaiterParams struct {
	Waiter *Waiter
	Args   []string
	Params map[string]string
}

// Server describes an OpenAPI server endpoint
type Server struct {
	Description string
	URL         string
	// TODO: handle server parameters
}

// Imports describe optional imports based on features in use.
type Imports struct {
	Fmt     bool
	Strings bool
	Time    bool
}

// OpenAPI describes an API
type OpenAPI struct {
	GoPackage    string
	Imports      Imports
	Name         string
	GoName       string
	PublicGoName string
	Title        string
	Description  string
	Servers      []*Server
	Operations   []*Operation
	Groups       []*OperationGroup
	Waiters      []*Waiter
}

// ProcessAPI returns the API description to be used with the commands template
// for a loaded and dereferenced OpenAPI 3 document.
func ProcessAPI(shortName string, api *openapi3.Swagger) *OpenAPI {
	apiName := shortName
	if api.Info.Extensions[ExtName] != nil {
		apiName = extStr(api.Info.Extensions[ExtName])
	}

	apiDescription := api.Info.Description
	if api.Info.Extensions[ExtDescription] != nil {
		apiDescription = extStr(api.Info.Extensions[ExtDescription])
	}

	result := &OpenAPI{
		GoPackage:    "main",
		Name:         apiName,
		GoName:       toGoName(shortName, false),
		PublicGoName: toGoName(shortName, true),
		Title:        api.Info.Title,
		Description:  escapeString(apiDescription),
	}

	for _, s := range api.Servers {
		result.Servers = append(result.Servers, &Server{
			Description: s.Description,
			URL:         s.URL,
		})
	}

	if api.Extensions[ExtCmdGroups] != nil {
		result.Groups = make([]*OperationGroup, 0)

		groups := make(map[string]struct {
			Name   string
			Short  string
			Long   string
			Parent string
		})
		if err := json.Unmarshal(api.Extensions[ExtCmdGroups].(json.RawMessage), &groups); err != nil {
			panic(err)
		}

		// Sort groups alphabetically
		names := func() []string {
			gn := make([]string, 0)
			for name := range groups {
				gn = append(gn, name)
			}
			return sort.StringSlice(gn)
		}()

		for _, name := range names {
			result.Groups = append(result.Groups, &OperationGroup{
				GoName: toGoName(name, true),
				Name:   name,
				Short:  groups[name].Short,
				Long:   groups[name].Long,
			})
		}

		// Link groups to their parent
		for _, group := range result.Groups {
			if parentName := groups[group.Name].Parent; parentName != "" {
				if parentName == group.Name {
					panic(fmt.Errorf("cyclic relationship detected: group %q is parent to itself", parentName))
				}

				// Look for referenced parent in known groups
				for i := range result.Groups {
					if result.Groups[i].Name == parentName {
						group.Parent = result.Groups[i]
						break
					}
				}
				if group.Parent == nil {
					panic(fmt.Errorf("unknown parent %q for group %q", parentName, group.Name))
				}
			}
		}
	}

	// Convenience map for operation ID -> operation
	operationMap := make(map[string]*Operation)

	var paths []string
	for path := range api.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		pathItem := api.Paths[path]

		if pathItem.Extensions[ExtIgnore] != nil {
			// Ignore this path.
			continue
		}

		pathHidden := false
		if pathItem.Extensions[ExtHidden] != nil {
			json.Unmarshal(pathItem.Extensions[ExtHidden].(json.RawMessage), &pathHidden)
		}

		// Group operations per path for later path-level sorting, produces a more stable result
		// over multiple generations.
		pathOperations := make([]*Operation, 0)

		for method, operation := range pathItem.Operations() {
			if operation.Extensions[ExtIgnore] != nil {
				// Ignore this operation.
				continue
			}

			name := operation.OperationID
			if operation.Extensions[ExtName] != nil {
				name = extStr(operation.Extensions[ExtName])
			}

			var aliases []string
			if operation.Extensions[ExtAliases] != nil {
				// We need to decode the raw extension value into our string slice.
				json.Unmarshal(operation.Extensions[ExtAliases].(json.RawMessage), &aliases)
			}

			params := getParams(pathItem, method)
			requiredParams := getRequiredParams(params)
			optionalParams := getOptionalParams(params)
			short := operation.Summary
			if short == "" {
				short = name
			}

			use := usage(name, requiredParams)

			description := operation.Description
			if operation.Extensions[ExtDescription] != nil {
				description = extStr(operation.Extensions[ExtDescription])
			}

			reqMt, reqSchema, reqExamples := getRequestInfo(operation)

			var examples []string
			if len(reqExamples) > 0 {
				wroteHeader := false
				for _, ex := range reqExamples {
					if _, ok := ex.(string); !ok {
						// Not a string, so it's structured data. Let's marshal it to the
						// shorthand syntax if we can.
						if m, ok := ex.(map[string]interface{}); ok {
							ex = shorthand.Get(m)
							examples = append(examples, ex.(string))
							continue
						}

						b, _ := json.Marshal(ex)

						if !wroteHeader {
							description += "\n## Input Example\n\n"
							wroteHeader = true
						}

						description += "\n" + string(b) + "\n"
						continue
					}

					if !wroteHeader {
						description += "\n## Input Example\n\n"
						wroteHeader = true
					}

					description += "\n" + ex.(string) + "\n"
				}
			}

			if reqSchema != "" {
				description += "\n## Request Schema (" + reqMt + ")\n\n" + reqSchema
			}

			method := strings.Title(strings.ToLower(method))

			hidden := pathHidden
			if operation.Extensions[ExtHidden] != nil {
				json.Unmarshal(operation.Extensions[ExtHidden].(json.RawMessage), &hidden)
			}

			groupName := ""
			var group *OperationGroup
			if operation.Extensions[ExtCmdGroup] != nil {
				json.Unmarshal(operation.Extensions[ExtCmdGroup].(json.RawMessage), &groupName)
				if groupName != "" {
					for _, g := range result.Groups {
						if g.Name == groupName {
							result.Imports.Strings = true
							group = g
							break
						}
					}
				}
			}

			returnType := "interface{}"
		returnTypeLoop:
			for code, ref := range operation.Responses {
				if num, err := strconv.Atoi(code); err != nil || num < 200 || num >= 300 {
					// Skip invalid responses
					continue
				}

				if ref.Value != nil {
					for _, content := range ref.Value.Content {
						if _, ok := content.Example.(map[string]interface{}); ok {
							returnType = "map[string]interface{}"
							break returnTypeLoop
						}

						if content.Schema != nil && content.Schema.Value != nil {
							if content.Schema.Value.Type == "object" || len(content.Schema.Value.Properties) != 0 {
								returnType = "map[string]interface{}"
								break returnTypeLoop
							}
						}
					}
				}
			}

			o := &Operation{
				HandlerName: Slug(name),
				GoName:      toGoName(name, true),
				Use:         use,
				Aliases:     aliases,
				Short:       short,
				Long:        escapeString(description),
				Method:      method,
				CanHaveBody: method == "Post" || method == "Put" || method == "Patch",
				ReturnType:  returnType,
				Path:        path,
				AllParams:   params,
				RequiredParams: requiredParams,
				OptionalParams: optionalParams,
				MediaType:      reqMt,
				Examples:       examples,
				Hidden:         hidden,
				Group:          group,
			}

			operationMap[operation.OperationID] = o

			pathOperations = append(pathOperations, o)

			for _, p := range params {
				if p.In == "path" {
					result.Imports.Strings = true
				}
			}

			for _, p := range optionalParams {
				if p.In == "query" || p.In == "header" {
					result.Imports.Fmt = true
				}
			}
		}

		sort.Slice(pathOperations, func(i, j int) bool {
			return pathOperations[i].HandlerName < pathOperations[j].HandlerName
		})
		result.Operations = append(result.Operations, pathOperations...)
	}

	if api.Extensions[ExtWaiters] != nil {
		var waiters map[string]*Waiter

		if err := json.Unmarshal(api.Extensions[ExtWaiters].(json.RawMessage), &waiters); err != nil {
			panic(err)
		}

		for name, waiter := range waiters {
			waiter.CLIName = Slug(name)
			waiter.GoName = toGoName(name+"-waiter", true)
			waiter.Operation = operationMap[waiter.OperationID]
			waiter.Use = usage(name, waiter.Operation.RequiredParams)

			for _, matcher := range waiter.Matchers {
				if matcher.Test == "" {
					matcher.Test = "equal"
				}
			}

			for operationID, waitOpParams := range waiter.After {
				op := operationMap[operationID]
				if op == nil {
					panic(fmt.Errorf("Unknown waiter operation %s", operationID))
				}

				var args []string
				for _, p := range op.RequiredParams {
					selector := waitOpParams[p.Name]
					if selector == "" {
						panic(fmt.Errorf("Missing required parameter %s", p.Name))
					}
					delete(waitOpParams, p.Name)

					args = append(args, selector)

					result.Imports.Fmt = true
					op.NeedsResponse = true
				}

				// Transform from OpenAPI param names to CLI names
				wParams := make(map[string]string)
				for p, s := range waitOpParams {
					found := false
					for _, optional := range op.OptionalParams {
						if optional.Name == p {
							wParams[optional.CLIName] = s
							found = true
							break
						}
					}
					if !found {
						panic(fmt.Errorf("Unknown parameter %s for waiter %s", p, name))
					}
				}

				op.Waiters = append(op.Waiters, &WaiterParams{
					Waiter: waiter,
					Args:   args,
					Params: wParams,
				})
			}

			result.Waiters = append(result.Waiters, waiter)
		}

		if len(waiters) > 0 {
			result.Imports.Time = true
		}
	}

	return result
}

// extStr returns the string value of an OpenAPI extension stored as a JSON
// raw message.
func extStr(i interface{}) (decoded string) {
	if err := json.Unmarshal(i.(json.RawMessage), &decoded); err != nil {
		panic(err)
	}

	return
}

func toGoName(input string, public bool) string {
	transformed := strings.Replace(input, "-", " ", -1)
	transformed = strings.Replace(transformed, "_", " ", -1)
	transformed = strings.Title(transformed)
	transformed = strings.Replace(transformed, " ", "", -1)

	if !public {
		transformed = strings.ToLower(string(transformed[0])) + transformed[1:]
	}

	return transformed
}

func escapeString(value string) string {
	transformed := strings.Replace(value, "\n", "\\n", -1)
	transformed = strings.Replace(transformed, "\"", "\\\"", -1)
	return transformed
}

func Slug(operationID string) string {
	re, _ := regexp.Compile("([a-z])([A-Z])")
	transformed := re.ReplaceAllString(operationID, "$1-$2")
	transformed = strings.ToLower(transformed)
	transformed = strings.Replace(transformed, "_", "-", -1)
	transformed = strings.Replace(transformed, " ", "-", -1)
	return transformed
}

func usage(name string, requiredParams []*Param) string {
	usage := Slug(name)

	for _, p := range requiredParams {
		usage += " " + Slug(p.Name)
	}

	return usage
}

func getParams(path *openapi3.PathItem, httpMethod string) []*Param {
	operation := path.Operations()[httpMethod]
	allParams := make([]*Param, 0, len(path.Parameters))

	var total openapi3.Parameters
	total = append(total, path.Parameters...)
	total = append(total, operation.Parameters...)

	for _, p := range total {
		if p.Value != nil && p.Value.Extensions["x-cli-ignore"] == nil {
			t := "string"
			tn := "\"\""
			if p.Value.Schema != nil && p.Value.Schema.Value != nil && p.Value.Schema.Value.Type != "" {
				switch p.Value.Schema.Value.Type {
				case "boolean":
					t = "bool"
					tn = "false"
				case "integer":
					t = "int64"
					tn = "0"
				case "number":
					t = "float64"
					tn = "0.0"
				}
			}

			cliName := Slug(p.Value.Name)
			if p.Value.Extensions[ExtName] != nil {
				cliName = extStr(p.Value.Extensions[ExtName])
			}

			description := p.Value.Description
			if p.Value.Extensions[ExtDescription] != nil {
				description = extStr(p.Value.Extensions[ExtDescription])
			}

			allParams = append(allParams, &Param{
				Name:        p.Value.Name,
				CLIName:     cliName,
				GoName:      toGoName("param "+cliName, false),
				Description: description,
				In:          p.Value.In,
				Required:    p.Value.Required,
				Type:        t,
				TypeNil:     tn,
			})
		}
	}

	return allParams
}

func getRequiredParams(allParams []*Param) []*Param {
	required := make([]*Param, 0)

	for _, param := range allParams {
		if param.Required || param.In == "path" {
			required = append(required, param)
		}
	}

	return required
}

func getOptionalParams(allParams []*Param) []*Param {
	optional := make([]*Param, 0)

	for _, param := range allParams {
		if !param.Required && param.In != "path" {
			optional = append(optional, param)
		}
	}

	return optional
}

func getRequestInfo(op *openapi3.Operation) (string, string, []interface{}) {
	mts := make(map[string][]interface{})

	if op.RequestBody != nil && op.RequestBody.Value != nil {
		for mt, item := range op.RequestBody.Value.Content {
			var schema string
			var examples []interface{}

			if item.Schema != nil && item.Schema.Value != nil {
				// Let's make this a bit more concise. Since it has special JSON
				// marshalling functions, we do a dance to get it into plain JSON before
				// converting to YAML.
				data, err := json.Marshal(item.Schema.Value)
				if err != nil {
					continue
				}

				var unmarshalled interface{}
				json.Unmarshal(data, &unmarshalled)

				data, err = yaml.Marshal(unmarshalled)
				if err == nil {
					schema = string(data)
				}
			}

			if item.Example != nil {
				examples = append(examples, item.Example)
			} else {
				for _, ex := range item.Examples {
					if ex.Value != nil {
						examples = append(examples, ex.Value.Value)
						break
					}
				}
			}

			mts[mt] = []interface{}{schema, examples}
		}
	}

	// Prefer JSON.
	for mt, item := range mts {
		if strings.Contains(mt, "json") {
			return mt, item[0].(string), item[1].([]interface{})
		}
	}

	// Fall back to YAML next.
	for mt, item := range mts {
		if strings.Contains(mt, "yaml") {
			return mt, item[0].(string), item[1].([]interface{})
		}
	}

	// Last resort: return the first we find!
	for mt, item := range mts {
		return mt, item[0].(string), item[1].([]interface{})
	}

	return "", "", nil
}

func writeFormattedFile(filename string, data []byte) {
	formatted, errFormat := format.Source(data)
	if errFormat != nil {
		formatted = data
	}

	err := ioutil.WriteFile(filename, formatted, 0600)
	if errFormat != nil {
		panic(errFormat)
	} else if err != nil {
		panic(err)
	}
}

func initCmd(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("main.go"); err == nil {
		fmt.Println("Refusing to overwrite existing main.go")
		return
	}

	data, _ := Asset("templates/main.tmpl")
	tmpl, err := template.New("cli").Parse(string(data))
	if err != nil {
		panic(err)
	}

	templateData := map[string]string{
		"Name":    args[0],
		"NameEnv": strings.Replace(strings.ToUpper(args[0]), "-", "_", -1),
	}

	var sb strings.Builder
	err = tmpl.Execute(&sb, templateData)
	if err != nil {
		panic(err)
	}

	writeFormattedFile("main.go", []byte(sb.String()))
}

func generate(cmd *cobra.Command, args []string) {
	data, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	// Load the OpenAPI document.
	loader := openapi3.NewSwaggerLoader()
	var swagger *openapi3.Swagger
	swagger, err = loader.LoadSwaggerFromData(data)
	if err != nil {
		log.Fatal(err)
	}

	funcs := template.FuncMap{
		"escapeStr": escapeString,
		"Slug":      Slug,
		"title":     strings.Title,
	}

	data, _ = Asset("templates/commands.tmpl")
	tmpl, err := template.New("cli").Funcs(funcs).Parse(string(data))
	if err != nil {
		panic(err)
	}

	name := strings.TrimSuffix(path.Base(args[0]), path.Ext(args[0]))
	if v, _ := cmd.Flags().GetString("name"); v != "" {
		name = v
	}

	templateData := ProcessAPI(name, swagger)

	if v, _ := cmd.Flags().GetString("go-package"); v != "" {
		templateData.GoPackage = v
	}

	var sb strings.Builder
	err = tmpl.Execute(&sb, templateData)
	if err != nil {
		panic(err)
	}

	output := name + ".go"
	if v, _ := cmd.Flags().GetString("output"); v != "" {
		output = v
	}

	writeFormattedFile(output, []byte(sb.String()))
}

func main() {
	root := &cobra.Command{}

	root.AddCommand(&cobra.Command{
		Use:   "init <app-name>",
		Short: "Initialize and generate a `main.go` file for your project",
		Args:  cobra.ExactArgs(1),
		Run:   initCmd,
	})

	generateCmd := &cobra.Command{
		Use:   "generate <api-spec>",
		Short: "Generate a `commands.go` file from an OpenAPI spec. Modified by MotaWord.",
		Args:  cobra.ExactArgs(1),
		Run:   generate,
	}
	generateCmd.Flags().StringP("name", "n", "",
		"specify a name to prefix symbols with in generated code (default <input file base name>)")
	generateCmd.Flags().StringP("go-package", "p", "main",
		"specify a Go package name for the generated code file")
	generateCmd.Flags().StringP("output", "o", "",
		"specify an alternative output file for generated code (default <input file base name>.go)")
	root.AddCommand(generateCmd)

	root.Execute()
}

// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/commands.tmpl
// templates/main.tmpl

package main


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataTemplatesCommandstmpl = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xdd\x6e\xe3\x36\x13\xbd\xa6\x9e\x62\x3e\x61\xbf\x40\x6a\x1d\x09" +
	"\x8b\xde\xb9\xf0\xc5\x36\xe9\x66\x03\x6c\x7e\xea\x24\xbd\x31\x02\x84\x96\xc6\x32\xbb\x12\xa9\x50\x54\x5a\x43\xd0" +
	"\xbb\x17\x43\x4a\x96\x9c\x28\xbb\x09\xd0\x00\x49\x6c\xce\x70\xce\x99\x33\x3f\x52\x1c\xc3\x89\x4a\x11\x32\x94\xa8" +
	"\xb9\xc1\x14\xd6\x3b\x50\x25\x4a\x5e\x8a\xe3\x24\x17\xc7\x9d\x41\xe9\x08\x4e\xaf\xe0\xf2\xea\x16\x7e\x3f\x3d\xbf" +
	"\x8d\xbc\x38\x86\x1b\x44\xd8\x1a\x53\x56\xf3\x38\xce\x84\xd9\xd6\xeb\x28\x51\x45\x9c\x72\x29\x30\xcf\x0c\xdf\xe5" +
	"\x4a\xc7\x93\xb1\x3c\xaf\xe4\xc9\x37\x9e\x21\x14\x5c\x48\xcf\x13\x45\xa9\xb4\x81\xc0\x63\xfe\x7a\x67\xb0\xf2\x3d" +
	"\xe6\xa3\x4c\x54\x2a\x64\x16\xff\x55\x29\x49\x07\x12\x4d\x4c\x70\xf4\xb9\x32\x5a\xc8\xac\xf2\x3d\x8f\xf9\xef\x83" +
	"\x8e\x93\x5c\xf8\x87\xb7\xaa\x72\xf3\xf1\x97\x38\x51\x6b\xcd\x27\x2d\x4f\xa2\x44\xed\x7b\xa1\xe7\x35\x0d\x7c\x90" +
	"\xbc\x40\x98\x2f\x20\xba\xa4\x0f\x6d\x6b\x0f\x79\x29\xec\xd9\x99\xea\x4f\xbd\x4d\x2d\x13\xe8\x6d\x6d\x7b\x83\xfa" +
	"\x09\x75\x15\x84\xb0\xba\x2f\x78\xb9\x72\x19\xdc\xbb\x7f\xd0\x78\x4c\xa3\xa9\xb5\x9c\xb2\x36\x1e\x63\x4d\x03\x9a" +
	"\xcb\x0c\xe1\x43\x65\x03\x59\xb4\x2e\x26\xc1\x31\xc6\x26\xef\x31\xe6\xa7\x58\x25\x5a\x94\x46\x28\xe9\xcf\xc1\x27" +
	"\x4a\x2e\x46\x74\x3a\x58\xa0\x6d\xfd\x99\xf3\xaf\x75\xfe\xcc\xef\x6e\xf9\x75\x6f\x6f\x67\x8e\x0d\xca\xd4\xe2\xb6" +
	"\xde\xcb\x5c\x97\x98\x89\xca\xa0\x0e\xaa\x7a\x9d\xa8\xa2\xe0\x32\x85\xb5\x52\x79\x68\xf3\x54\xca\x10\xfb\x24\x17" +
	"\xd1\x52\x29\xe3\x79\x4c\x6c\x60\xe4\x49\xb4\xad\xd3\x02\x8e\x6c\x51\xa2\x13\x67\xb1\xf9\xdc\x55\xd8\x91\x93\x4e" +
	"\x69\x47\xeb\x66\xab\xb4\x71\x86\xe8\x56\x98\x7c\xb0\x7c\x55\x32\x9b\x5b\xb4\x0b\xae\xbf\xa5\xea\x6f\x19\x58\xaf" +
	"\x67\xc9\x87\xe4\x4c\xf9\x00\xe6\x15\x5a\x12\x3d\xc3\xc8\x06\x87\xc5\xb3\xe8\x63\x0f\x02\x81\xc5\x1b\x50\x48\x30" +
	"\x6f\x54\x4d\x55\x52\x5b\x92\x95\x0a\x7a\xd5\x7f\xeb\x6a\x6a\x75\x0d\xac\x6c\xa3\x16\x88\xae\x6c\x3c\x9e\x5f\x73" +
	"\xcd\x8b\xbe\xfc\xec\x89\x6b\xaa\xc1\xd0\x83\xf6\xdb\xed\xae\xc4\xce\x63\xa8\x1a\x7d\x4b\x8a\x94\x30\x27\x24\x1e" +
	"\x34\x8e\xee\xaa\x41\xc8\x03\x8d\x9d\x24\x7b\xd3\x6b\x22\x5b\x5d\x7a\x75\x19\xfb\xa4\xb3\x6a\x0e\x0e\xf1\x42\x48" +
	"\x51\xd4\xc5\x25\x9d\x05\x4d\x03\x39\x4a\x88\x96\xf8\x58\x0b\x8d\xe9\x3e\xb3\xee\xe2\xb2\x96\x73\x20\x35\x02\xa2" +
	"\xfd\xd3\x01\xe7\x19\x70\x9d\x55\xb0\xea\xba\xbe\x53\x8b\xb1\x61\x50\xec\x00\x47\x67\x68\x6e\xac\x47\xe0\x3b\x13" +
	"\x55\x83\x7e\xa8\xfd\x9c\xef\x62\x01\xbe\xdf\xdf\xef\x03\x2c\xa6\xa6\x78\xb5\x8f\x79\x2e\x4d\x1f\xf0\x58\xc8\x14" +
	"\xff\xf1\xc3\xfb\x95\x1d\xa2\x7b\x17\xc7\xc9\xcd\x58\xad\x73\x22\xe3\x5c\x7f\xb6\xf2\x5c\x73\xb3\xb5\xad\xe4\x3c" +
	"\x86\xb6\x10\x33\xf8\x50\x92\x06\xb6\x2d\x5e\xa8\xd2\x11\x6c\x1a\x10\x1b\xc0\xc7\xce\x37\x3a\x97\xe0\x97\xdc\x6c" +
	"\xfd\xc1\xc5\xa2\x2e\xa0\xdb\x97\xd1\x12\xcb\x9c\x27\x18\xd4\x3a\x9f\x51\x21\x1f\x9a\x87\xb6\xa5\xf4\x5c\x80\xae" +
	"\x71\x9a\xe6\xa1\x7d\xa0\xda\x5a\x65\x57\x64\xa7\xe4\xef\x67\xf0\x31\x1c\xa0\xfb\x0d\xc0\x5e\xb4\x16\x63\x1a\x1f" +
	"\xfb\x19\x3f\xc9\x05\x4a\x13\x51\xba\x17\x68\xb6\x8a\xbc\x82\x90\x76\x0a\xb1\x08\xff\xc3\xd4\x1f\x6b\xd4\xbb\x71" +
	"\xee\xc4\x62\x01\x1a\x1f\xa3\x4f\x69\xfa\x07\x59\x5d\x57\x5e\xf6\x9b\xe3\x59\x7e\xe3\xe4\x68\x09\xbc\x40\xd8\x22" +
	"\x4f\x51\xbf\x0a\xf1\xc5\x9a\xdf\x8e\xf1\x1d\x01\x7f\x34\xec\xb6\x6b\x0f\xa7\xfd\x7f\x8b\xfd\xc0\x5f\x8a\xdc\xce" +
	"\x7f\xcf\x72\xaf\xd6\xb4\x4e\x6f\x10\x6a\x53\x98\xe8\xa6\xd4\x42\x9a\x4d\xe0\xff\xff\xc9\x9f\x1d\x82\x87\xe1\x08" +
	"\x6a\xa4\xdd\x2b\xaa\xbd\x45\xb6\x77\x21\x8e\x94\x64\xaf\x2b\x2a\x36\x10\x9d\x70\xf9\x85\x3f\xe1\x6f\x2a\xdd\x0d" +
	"\x57\xd6\x2a\xdd\xcd\x00\xb5\xee\x9b\xf6\x0c\x0d\x79\x38\x4e\x17\x98\x0a\xde\xed\xd1\x51\x3d\x27\x16\x56\xdb\xce" +
	"\xf7\x05\x26\x01\xb4\xa6\xaa\x48\x91\x0f\xa5\x28\xb9\x14\x49\x80\x5a\x87\x7b\xba\xc3\x05\xe2\x41\x37\x46\x4b\x68" +
	"\x5a\xaa\x13\x25\x0d\x4a\x73\x4c\xac\xfc\x19\xbc\xa4\x19\x46\xc4\xbf\xdb\x77\x14\x36\xfc\xbe\x3a\x1a\xab\x72\xaf" +
	"\x00\x81\x9d\xaa\x60\x58\x8f\x13\x89\x3c\xcf\xa3\x0f\x44\x4f\xa1\x14\x13\x95\x62\x0a\x42\x1a\xd4\x1b\x9e\x60\xd3" +
	"\x1e\x84\xea\x44\xbe\x93\x05\xd7\xd5\x96\xe7\x81\x43\x3f\xea\xee\x85\xbf\xbe\x07\x90\x22\x7d\x56\xba\xe0\xc6\xa0" +
	"\xee\x3e\x05\x7d\x24\xeb\xd2\xba\x17\x17\xfa\x43\xef\x15\xa4\x63\xf7\xe8\xa0\xa7\x89\xdb\x40\x3f\x9a\xb7\xa4\x48" +
	"\xa3\xcf\x39\xcf\xaa\x20\x8c\x9c\xaa\x7f\x72\x7d\x1d\x1c\x1d\xf4\xa4\x2b\x44\x95\xd7\xd9\xb8\x91\xfd\xfe\x77\xfa" +
	"\x7d\xe0\xb0\x81\x5b\x52\x7d\x38\x68\xbd\x7f\x03\x00\x00\xff\xff\xf1\xda\x4d\xed\xa1\x0b\x00\x00")

func bindataTemplatesCommandstmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesCommandstmpl,
		"templates/commands.tmpl",
	)
}



func bindataTemplatesCommandstmpl() (*asset, error) {
	bytes, err := bindataTemplatesCommandstmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/commands.tmpl",
		size: 2977,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1538449585, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesMaintmpl = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x41\x4b\xc3\x40\x10\x85\xcf\x3b\xbf\x62\xc8\x41\x12\xb0\x9b\x7a" +
	"\xed\xad\x68\x0e\x5e\xac\x88\x78\x5f\x37\x93\xed\x60\x76\x66\xd9\x6c\x4a\x25\xe4\xbf\x4b\x2a\xe2\xed\xbd\xf7\xf1" +
	"\x1e\x2f\x39\xff\xe5\x02\x61\x74\x2c\x00\x1c\x93\xe6\x82\x35\x98\x2a\x70\x39\xcf\x9f\xd6\x6b\x6c\x7b\x27\x4c\x63" +
	"\x28\xee\x7b\xd4\xdc\x6a\x22\x71\x89\x77\x7e\xe4\x5d\x20\xa1\xec\x8a\xe6\xd6\x8f\x5c\x41\x03\x30\xcc\xe2\x6f\x63" +
	"\x75\x83\x0b\x18\x3f\xb2\x7d\x16\x2e\xf5\xdd\xa6\x1e\x55\x06\x0e\x0b\x18\x73\x4c\xe9\xc5\x45\x3a\x20\x62\xb5\x2c" +
	"\x68\x37\x83\xeb\x5a\xdd\x83\x31\x9d\x5c\x5e\x33\x0d\x7c\x3d\xfc\xb3\x4e\x2e\x7f\xf8\x83\xf2\xc4\x2a\xb7\xea\x83" +
	"\xdd\xdb\xfd\x96\xae\x0d\x80\x69\x5b\x7c\x3f\x3d\x9d\x0e\x78\xec\x7b\xcc\x14\x78\x2a\x94\xd1\x6b\x8c\x4e\xfa\x09" +
	"\xcf\x94\xc9\xc2\xef\xa7\x37\xd5\x62\xbb\x2b\xf9\xb9\x50\xdd\xc0\x0a\x3f\x01\x00\x00\xff\xff\xd7\x90\x9c\xb4\x08" +
	"\x01\x00\x00")

func bindataTemplatesMaintmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesMaintmpl,
		"templates/main.tmpl",
	)
}



func bindataTemplatesMaintmpl() (*asset, error) {
	bytes, err := bindataTemplatesMaintmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/main.tmpl",
		size: 264,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1538448749, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"templates/commands.tmpl": bindataTemplatesCommandstmpl,
	"templates/main.tmpl":     bindataTemplatesMaintmpl,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}


type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"templates": {Func: nil, Children: map[string]*bintree{
		"commands.tmpl": {Func: bindataTemplatesCommandstmpl, Children: map[string]*bintree{}},
		"main.tmpl": {Func: bindataTemplatesMaintmpl, Children: map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

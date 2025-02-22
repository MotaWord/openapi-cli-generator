# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 1.1.0 - 2020-10-14

### Changed

- JSON output is not HTML-escaped anymore
- If `--raw` is specified, the JSON output is not prettified (no colors/indented formatting)


## 1.1.1 - 2021-05-17

**Forked from github.com/exoscale/openapi-cli-generator**

## 1.0.0 - 2020-09-23

**Forked from github.com/danielgtaylor/openapi-cli-generator**

### Added

- New `x-cli-cmd-groups` extension enabling [CLI commands grouping](https://github.com/exoscale/openapi-cli-generator/#commands-grouping).
- New `--output|-o` option to the `generate` command.
- New `--name|-n` option to the `generate` command.
- New `--go-package|-p` option to the `generate` command.

### Changed

- The `generate` command now accepts OpenAPI spec files with any extension, not only `.yaml`.
- Generated CLI configuration directory is not created automatically by default.
- Generated CLI temporary data caching is now optional, and disabled by default.


## 2020-02-27

- Add [enhanced JMESPath](https://github.com/danielgtaylor/go-jmespath-plus) support.


## 2020-01-17

- Add support for multi-auth, where two different auth schemes can be used. This
  adds a new `cli.UseAuth(typeName, AuthHandler)` method that can be used to
  register new auth types by name. This is *backward compatible* and the
  existing (but now deprecated) credentials calls continue to work, but cannot
  be used in conjuction with the new multi-auth system.


## 2020-01-10

- Add support for OAuth2 Authorization Code with PKCE https://oauth.net/2/pkce/


## 2019-03-10

- Better rendering of error messages.
- Switch to Go modules.


## 2019-03-09

- Generate methods of the format `{{ API Name }}{{ Operation Name }}(...)` for each API operation. These can be used by custom code as if you were invoking the CLI, but it returns rather than formats the response.
- Decouple CLI commands from the command path used to register middleware. Each API operation now has one and only one command path regardless of which CLI command calls it.
- Add support for waiters through the `x-cli-waiters` extension.


## 2018-09-29

- Initial release

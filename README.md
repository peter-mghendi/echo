# echo

echo is a WebAssembly guest component that acts as a sample handler for [dispatch](https://github.com/peter-mghendi/dispatch). It implements the handle function defined in this repository's WIT, which the dispatch runtime calls with a Request and expects a Response back. This sample implementation echoes the request's headers and body back with a 200 status.

## Prerequisites
- [Go](https://go.dev/dl/) 1.25.5 or later
- [componentize-go](https://github.com/bytecodealliance/componentize-go) (latest version)
- (Optional) [wasm-tools CLI](https://github.com/bytecodealliance/wasm-tools) or [WIT IDL for VSCode](https://github.com/bytecodealliance/vscode-wit)

> [!NOTE]
> Only edit `export_dispatch_guest_handler/echo.go`. Everything else in this repository is generated or defines the contract, and should not be modified by hand.

## Build

```bash
componentize-go build --output ./echo.wasm
```

This provides an `echo.wasm` file in the specified directory that can be uploaded to dispatch. 
Optionally, verify this with `wasm-tools` or using the WIT IDL for VSCode extension.

```bash
wasm-tools component wit echo.wasm
```

You should see the generated world exporting `dispatch:guest/handler@1.0.0` with a handle method.

## Tests

`export_dispatch_guest_handler/echo_test.go` contains unit tests for the default echo behavior. Run them with:

```bash
go test ./export_dispatch_guest_handler/
```

If you change what Handle does, update or replace these tests to match. 
The tests run natively with the standard Go toolchain, so they don't require building the component.

## Making edits

Edit the `export_dispatch_guest_handler/echo.go` file to update the hhandler logic, and update any tests. 
It may be renamed or deleted, but a handler MUST be implemented in `package export_dispatch_guest_handler` that exports a function with the signature:

```go
func Handle(req dispatch_guest_types.Request) dispatch_guest_types.Response
```

## Regenerating WIT bindings

> [!NOTE]
> This only needs to be done if the WIT definitions change

- Replace the contents of the `/wit` directory with the new file, and increment the version.
- Regenerate the bindings:
and 
```bash
componentize-go --world dispatch bindings
go mod tidy
```

- Update `export_dispatch_guest_handler/echo.go` so that it compiles against the updated bindings.
- Follow the instructions above to [Build](#build).


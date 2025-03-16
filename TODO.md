
# toolchain

# toolchain/test.go

- [ ] Build all tests
- [ ] Build a separate standalone test server `main.go`
- [ ] Include assets via `go:embed` of `/public/*` folder
- [ ] Parse out methods of all packages in `gooey-components/pkg` with `Test` prefix and add it to a generated `main.go` entry point
- [ ] Build the generated `main.go` as a `/public/main.wasm` file
- [ ] Start the `main.go` server
- [ ] Start playwright and navigate to `http://localhost:3000`
- [ ] Get the console output to verify whether tests ran through or not



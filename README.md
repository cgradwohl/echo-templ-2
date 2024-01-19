## Notes
- `go run .` - compiles and execute all go files
- `go run server.go` - compiles and executes server.go as well as any imported modules

- modules only export Capitalized entities, while lowercase entities are private
- you need to compile the templ templates in order to satisfy the compile, because server.go cannot import the .templ file, it needs to import the _templ.go file.

debug:
- abrir desde el vscode la carpeta primero
- esa carpeta está configurada para que se puede debugear el archivo hello.go
- hay que ir a la carpeta hello (con el mismo vscode) y ejecutar /home/sebastian/go/bin/dlv-dap dap
- luego lanzar el debugger

```
go mod init example/hello
go run .
go mod tidy
go mod edit -replace example.com/greetings=../greetings
go test
go test -v
go env
go build
    go list -f '{{.Target}}'
    go env -w GOBIN=/path/to/your/bin
        go env -w GOBIN=/home/sebastian/go/bin/
    go install // queda instalado en el GOBIN

```
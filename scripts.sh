go mod init  webapp


go get github.com/go-chi/chi/v5

go test ./...

go run ./cmd/web

GOOS=windows GOARCH=amd64 go build -o myprogram.exe

go test ./...

go test -cover ./...

go test -v ./...


go test -run function_name

go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

go get github.com/alexedwards/scs/v2
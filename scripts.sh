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

# for sessions
go get github.com/alexedwards/scs/v2  


go get github.com/jackc/pgx/v4

go get -u github.com/ory/dockertest/v3

go test -v -tags=integration

go get github.com/golang-jwt/jwt/v4
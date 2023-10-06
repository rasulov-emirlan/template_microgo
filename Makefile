dev:
	go build -o bin/apiserver cmd/apiserver/main.go
	bin/apiserver --config=.env

migrate_new:
	goose -dir internal/storage/postgresql/migrations create $(name) sql
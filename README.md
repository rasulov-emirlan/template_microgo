# template_microgo
This is how i currently start my microservice 

## domain
Here we will have a package for each entity we have. This package should have following files: dto.go, errors.go, entity.go, service.go. It should contain all the buisness logic about the entity.

## storage
Here we will have everything related to storage. Repositories and caching solutions will be stored here.
```
├── Makefile
├── README.md
├── cmd
│   └── api
│       └── main.go
├── config
│   └── config.go
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   ├── storage
│   │   └── postgres
│   └── transport
│       └── httprest
└── pkg
    └── logger
        └── logger.go
```
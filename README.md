# Hexagonal architecture bank account example

Esse repositório contém um pouco dos meus estudos sobre arquitetura hexagonal.

Go Lang é a linguagem utilizada.

## Estrutura

```
├── cmd
│   ├── cli
│   │   └── main.go
│   └── httpserver
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   ├── account.go
│   │   ├── account_test.go
│   │   └── health.go
│   ├── handlers
│   │   ├── bank.go
│   │   └── health.go
│   ├── repositories
│   │   ├── memory-db.go
│   │   └── memory-db_test.go
│   └── services
│       ├── bank-service.go
│       ├── bank-service_test.go
│       └── health-service.go
```

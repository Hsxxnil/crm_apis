SCRIPTS = $(shell cd /d %cd% && echo %cd%\scripts)
PROJECT = $(shell cd /d %cd% && echo %cd%)
GO = go
OUTPUTS = $(shell cd /d %cd% && echo %cd%\deploy)
TAG ?= debug

## 首次使用專案模版時, 必要執行一次
setup:
	copy $(PROJECT)\config\config.go.example $(PROJECT)\config\debug_config.go
	rem copy $(PROJECT)\config\config.go.example $(PROJECT)\config\production_config.go
	copy $(PROJECT)\air.example.windows $(PROJECT).air.toml

## 映射遠端Ports至本地端Ports
ssh:
	go run -tags $(TAG) $(PROJECT)\tools\ssh\ssh.go

## 開發中
air:
	air

migration:
	go run -tags $(TAG) $(PROJECT)\tools\migration\migration.go

## 以下由CI/CD人員維護!!!
update_lib:
	rem brew install golang-migrate
	rem brew install golangci-lint
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest
	go get -u
	rem go get -u ...
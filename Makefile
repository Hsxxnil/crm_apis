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

## by Fleet
format:
	goimports -w $(PROJECT)

## 以下由CI/CD人員維護!!!
authority:
	set GOOS=linux&& set GOARCH=amd64&& $(GO) build -tags $(TAG) -o $(OUTPUTS)/authority cmd/lambda/authority.go
	zip -D -j -r $(OUTPUTS)/authority.zip $(OUTPUTS)/authority

clean:
	del /f /q $(OUTPUTS)

task:
	make clean
	make authority

changeLog:
	git-chglog > $(PROJECT)\changeLog.md

update_lib:
	rem brew install golang-migrate
	rem brew install golangci-lint
	rem go install github.com/swaggo/swag/cmd/swag@latest
	rem go install github.com/cosmtrek/air@latest
	go get -u
	go get -u ...
SCRIPTS = $(shell cd /d %cd% && echo %cd%\scripts)
PROJECT = $(shell cd /d %cd% && echo %cd%)
GO = go
OUTPUTS = $(shell cd /d %cd% && echo %cd%\deploy)
TAG ?= debug

authority:
	set GOOS=linux&& set GOARCH=amd64&& $(GO) build -tags $(TAG) -o $(OUTPUTS)/authority cmd/lambda/authority.go
	zip -D -j -r $(OUTPUTS)/authority.zip $(OUTPUTS)/authority

clean:
	del /f /q $(OUTPUTS)

task:
	make clean
	make authority

setup:
	copy $(PROJECT)\config\config.go.example $(PROJECT)\config\debug_config.go
	rem copy $(PROJECT)\config\config.go.example $(PROJECT)\config\production_config.go
	copy $(PROJECT)\air.example.windows $(PROJECT).air.toml

air:
	air

migration:
	go run -tags $(TAG) $(PROJECT)\tools\migration\migration.go

ssh:
	go run -tags $(TAG) $(PROJECT)\tools\ssh\ssh.go

format:
	goimports -w $(PROJECT)

changeLog:
	git-chglog > $(PROJECT)\changeLog.md

update_lib:
	rem brew install golang-migrate
	rem brew install golangci-lint
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest
	rem go install golang.org/x/tools/cmd/goimports@latest
	go get -u
	rem go get -u ...
	go get -u all

autoMigrate:
	set CGO_ENABLED=1
	go run -tags $(TAG) $(PROJECT)\tools\autoMigrate\main.go
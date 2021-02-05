PHONY: run-docker-mysql8
run-docker-mysql8:
	@docker run --name mysql8 -p3306:3306 -e MYSQL_ROOT_PASSWORD=pass4root -e MYSQL_DATABASE=translation -d mysql:8

.PHONY: vendor
vendor:
	@go mod vendor

PHONY: full-deploy
full-deploy: run-docker-mysql8 vendor
	@go run .\cmd\translation\main.go

PHONY: run
run:
	@go run .\cmd\translation\main.go

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m
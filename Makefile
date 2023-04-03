all: test run

test:
	go test ./... -cover -race -count=1

run:
	go run cmd/omdb/app.go


remockgen:
	@echo "===Remockgen==="
	@echo "==REPOSITORY=="
	@mockgen -source=./internal/repository/mysql.go -destination=./testfile/repository/mock_mysql.go
	@mockgen -source=./internal/repository/external_api.go -destination=./testfile/repository/mock_ext_api.go
	@echo "==REPOSITORY DONE=="
	@echo "==USECASE=="
	@mockgen -source=./internal/usecase/omdb_client.go -destination=./testfile/usecase/mock_omdb_client.go
	@echo "====USECASE DONE==="
	@echo "====Done==="
install:
	glide install

unit_tests:
	go test -v ./server ./domain ./service/tesco -tags=unit

integration_tests:
	go test -v ./service/tesco -tags=integration

system_tests:
	go test -v ./systemtest -tags=system

docker_tests:
	docker-compose -f docker/docker-compose-tests.yml -p groceries-tesco rm -v -f
	docker-compose -f docker/docker-compose-tests.yml -p groceries-tesco up --abort-on-container-exit

prod_build:
	GOOS=linux GOARCH=386 go build -o bin/groceries-tesco cmd/main.go

prod_image:
	make docker_tests && make prod_build && docker build -t groceries-tesco -f docker/Dockerfile .

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

docker_tests:
	docker-compose -f docker/docker-compose-tests.yml -p groceries-tesco rm -v -f
	docker-compose -f docker/docker-compose-tests.yml -p groceries-tesco up --abort-on-container-exit
		

benchmark:
	docker-compose up -d
	@echo "Waiting the database to be ready"
	@sleep 10
	go test --bench=. --benchmem  -v .
	@make clean
	
clean:
	docker-compose down

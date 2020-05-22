run:
	docker-compose up -d

benchmark:
	go test --bench=. --benchmem  -v .

engine:
	go build -o engine ./main/main.go
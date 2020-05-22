run:
	docker-compose up -d

benchmark:
	go test --bench=. --benchmem  -v .

auto-increment:
	go build -o aiengine ./autoincrementid/main/main.go
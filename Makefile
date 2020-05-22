run:
	docker-compose up -d

benchmark:
	go test --bench=. --benchmem  -v .

.PHONY: auto-increment
auto-increment:
	go build -o aiengine ./autoincrementid/main/main.go

.PHONY: offset-limit
offset-limit:
	go build -o olengine ./offsetlimit/main/main.go

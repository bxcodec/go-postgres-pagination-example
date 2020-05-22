run:
	docker-compose up -d

benchmark:
	go test --bench=. --benchmem  -v .

auto-increment:
	go build -o aiengine ./autoincrementid/main/main.go

offset-limit:
	go build -o olengine ./offsetlimit/main/main.go
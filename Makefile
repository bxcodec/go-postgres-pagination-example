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

.PHONY: uuid-created-time
uuid-created-time:
	go build -o ucengine ./uuidcreatedtime/main/main.go

.PHONY: page-number
page-number:
	go build -o pnengine ./pagenumber/main/main.go

swagger:
	docker run --rm -it -e GOPATH=/go -e GO111MODULE=off -v $(PWD):/go/src/expenses -w /go/src/expenses quay.io/goswagger/swagger validate api/v1/swagger.json
	docker run --rm -it -e GOPATH=/go -e GO111MODULE=off -v $(PWD):/go/src/expenses -w /go/src/expenses quay.io/goswagger/swagger generate server -s api/v1 -m api/v1/models -f api/v1/swagger.json --exclude-main

compile:
	go build .

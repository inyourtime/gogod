dev:
	nodemon --exec go run app/main.go --signal SIGTERM
test:
	go test ./... -cover -coverprofile=test_result/coverage.out
	go tool cover -html=test_result/coverage.out -o test_result/cover.html
build:
	go build -o server ./app
run:
	./server
clean:
	rm -f server
build-local:
	docker build -t gogod . 
build-dev:
	docker buildx build --push --tag inyourtime/ecommerce-be:dev --platform=linux/amd64 .	
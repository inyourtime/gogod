dev:
	nodemon --exec env-cmd -f .env go run app/main.go --signal SIGTERM
test:
	grc go test -v ./...
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
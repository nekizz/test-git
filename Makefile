dev:
	rm -f ./bin/go-test-service
	go build -o ./bin/go-test-service .
	docker build -t localhost:5000/go-test-service .
	docker push localhost:5000/go-test-service
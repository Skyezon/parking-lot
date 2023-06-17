build:
	go build -o ./bin/ParkingLot

run:
	go build -o ./bin/ParkingLot
	./bin/ParkingLot

test:
	go test ./...

test-functional:
	go run functional_testing/functional_testing.go -url=http://localhost:8080 -case=1
	go run functional_testing/functional_testing.go -url=http://localhost:8080 -case=2

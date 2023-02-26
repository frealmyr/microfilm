export WATCH_DIR=/home/fredrick/Repositories/frealmyr/microfilm/assets/pictures

run:
	go run ./cmd/microfilm

build:
	go build -o bin/microfilm ./cmd/microfilm

bench:
	go build -o bin/microfilm ./cmd/microfilm
	time ./bin/microfilm

startredis:
	docker pull redis && docker run --name redis-test-instance -p 6379:6379 -d redis

stopredis:
	docker stop $$(docker container ls -aq --filter name=redis*)
	docker rm $$(docker container ls -aq --filter name=redis*)

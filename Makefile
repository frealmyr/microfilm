all:
	WATCH_DIR=/home/fredrick/Repositories/frealmyr/microfilm/assets/pictures

build:
	go build -o bin/microfilm ./cmd/microfilm

run:
	WATCH_DIR=/home/fredrick/Repositories/frealmyr/microfilm/assets/pictures go run ./cmd/microfilm

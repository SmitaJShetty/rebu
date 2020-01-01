GO=go
GOBUILD= go build
APPNAME=carttrip
APP_BUILD_FOLDER=build/carttrip
DBNAME=carttripdb
DB_DOCKER_NAME=carttripdb

.PHONEY: build 
bld: 
		cd cmd/ && $(GOBUILD) -o ../$(APP_BUILD_FOLDER)

build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 $(GOBUILD) -a -installsuffix cgo -o $(APP_BUILD_FOLDER)
	redis-server &
	./$(APPNAME)


clean: 
	rm $(APP_BUILD_FOLDER)

redis-start:
	docker pull redis
	docker run --name redis-test-instance -p 6379:6379 -d redis

db-upbd:
	docker-compose up --build --detach

db-stop:
	docker-compose stop
	docker rm $(DB_DOCKER_NAME)

run: bld
	$(APP_BUILD_FOLDER)

test: bld
	go test ./...
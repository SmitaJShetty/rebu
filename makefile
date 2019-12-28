GO=go
GOBUILD= go build
APPNAME=carttrip
APP_BUILD_FOLDER=build/carttrip

.PHONEY: build 
bld: 
		cd cmd/ && $(GOBUILD) -o ../$(APP_BUILD_FOLDER)

build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 $(GOBUILD) -a -installsuffix cgo -o $(APP_BUILD_FOLDER)
	redis-server &
	./$(APPNAME)


clean: ## Remove all the temporary and build files
	go clean

redis-start:
	docker pull redis
	docker run --name redis-test-instance -p 6379:6379 -d redis

run: bld
	$(APP_BUILD_FOLDER)
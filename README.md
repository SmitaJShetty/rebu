# rebu

An golang backend application that fetches data from mysql db and uses caching. Redis cache is used. 

A makefile has been created with items for setting up redis, mysql docker containers. Application itself is not containerized and conncects to redis and mysql db containers. 

Arguments that are to be provided for Get are pickupdate, comma separated list of medallions, fresh value. When fresh is set to true (case insensitive), data is fetched from mysql db; 

Gomod is used in this implementation. I've implemented and tested with version go1.12.5. Any version that supports gomod can be used to test this. Go has to be available locally. 

**Setup instructions:**

Install go

Set up environment parameters: 
```
export APP_LISTENER_ADDR=localhost:8090
export DB_USERNAME=testuser
export DB_PASSWORD=testpass
export REDIS_LISTENER_ADDR=localhost:6379
```
This command sets up redis cache container:

``` make redis-start ```

This command sets up mysql db container, executes the script file by inserting initial setup data:

``` make db-upbd ```

Following command builds and and runs the golang application:

``` make run ``` 

``` make test``` will execute tests

------------------------------------------------------------------------------

*Console scripts to execute http endpoints:*
```
curl "http://localhost:8090/cartrip/2013-12-01?medallionlist=D7D598CD99978BD012A87A76A7C891B7,42D815590CE3A33F3A23DBF145EE66E3&fresh=true"

curl -X POST  "http://localhost:8090/cartrip/clearcache"

curl "http://localhost:8090/cartrip/2013-12-01?medallionlist=D7D598CD99978BD012A87A76A7C891B7,42D815590CE3A33F3A23DBF145EE66E3,B672154F0FD3D6B5277580C3B7CBBF8E&fresh=true"

curl "http://localhost:8090/cartrip/2013-12-01?medallionlist=D7D598CD99978BD012A87A76A7C891B7,42D815590CE3A33F3A23DBF145EE66E3,B672154F0FD3D6B5277580C3B7CBBF8E"

```
-------------------------------------------------------------------------------






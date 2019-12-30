# rebu

An golang backend application that fetches data from mysql db and uses caching. Redis cache is used. 

A makefile has been created with items for setting up redis, mysql docker containers. Application itself is not containerized and conncects to redis and mysql db. 

Arguments that are to be provided for Get are pickupdate, comma separated list of medallions, isFresh value. When isFresh is set to true (case insensitive), data is fetched from mysql db; 

Setup instructions:

Set up environment parameters: 
```
export APP_LISTENER_PORT=localhost:8090
export DB_USERNAME=testuser
export DB_PASSWORD=testpass
export REDIS_LISTENER_PORT=localhost:6379
```
This command sets up redis cache container:

``` make redis-start ```

This command sets up mysql db container, executes the script file by inserting initial setup data:

``` make db-updb ```

Following command builds and and runs the golang application:

``` make run ``` 

------------------------------------------------------------------------------

*Console scripts to execute http endpoints:*
```
curl "http://localhost:8090/cartrip/2013-12-01?medallionlist=D7D598CD99978BD012A87A76A7C891B7,42D815590CE3A33F3A23DBF145EE66E3&isFresh=true"

curl "http://localhost:8090/cartrip/clearcache" METHOD=Post  

curl "http://localhost:8090/cartrip/2013-12-01?medallionlist=D7D598CD99978BD012A87A76A7C891B7,42D815590CE3A33F3A23DBF145EE66E3,B672154F0FD3D6B5277580C3B7CBBF8E&isFresh=true"

curl "http://localhost:8090/cartrip/2013-12-01?medallionlist=D7D598CD99978BD012A87A76A7C891B7,42D815590CE3A33F3A23DBF145EE66E3,B672154F0FD3D6B5277580C3B7CBBF8E"

```
-------------------------------------------------------------------------------






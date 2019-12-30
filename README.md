# rebu

An golang backend application that fetches data from mysql db and uses caching. Redis cache is used currently. 

A makefile has been created with items for setting up redis, mysql docker containers. Application itself is not containerized and conncects to redis and mysql db. 

Setup instructions:

Following command builds and and runs the golang application:

``` make run ``` 

This command sets up mysql db container, executes the script file by inserting initial setup data:

``` make db-updb ```

This command sets up redis cache container:

``` make redis-start ```


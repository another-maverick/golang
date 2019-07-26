
This is a webservice that accepts a GET, POST, PUT or DELETE request and make changes to the backend posts table in postgres

1. Create a postgres user
create user <webservice>;

2. Create a postgres database
create database <webservice>;

postgres=# create user webservice;
CREATE ROLE
postgres=# create database webservice;
CREATE DATABASE


Create a SQL file that will be used to create a table


cat webservice_posts.sql
create table posts (
    id serial primary key,
        content text,
            author varchar(255)
            
);

Execute a psql command to create the table

psql -U webservice -f webservice_posts.sql -d webservice
CREATE TABLE

######

Now the database is created. Now start the go webservice which is listening on port:12345 (port hard coded in the code)

check the table rows before making changes

webservice=# select * from posts;
 id | content | author
 ----+---------+--------
 (0 rows)

 Execute a POST request via curl command to add an entry in the DB

 curl -k -i -v -X POST -H "Content-Type: application/json" -d '{"content": "Random Post Content", "author": "Maverick"}' http://127.0.0.1:12345/post/
 Note: Unnecessary use of -X or --request, POST is already inferred.
 *   Trying 127.0.0.1...
 *   * TCP_NODELAY set
 *   * Connected to 127.0.0.1 (127.0.0.1) port 12345 (#0)
 *   > POST /post/ HTTP/1.1
 *   > Host: 127.0.0.1:12345
 *   > User-Agent: curl/7.54.0
 *   > Accept: */*
 *   > Content-Type: application/json
 *   > Content-Length: 56
 *   >
 *   * upload completely sent off: 56 out of 56 bytes
 *   < HTTP/1.1 200 OK
 *   HTTP/1.1 200 OK
 *   < Date: Fri, 26 Jul 2019 22:41:08 GMT
 *   Date: Fri, 26 Jul 2019 22:41:08 GMT
 *   < Content-Length: 0
 *   Content-Length: 0
 *

 Now check the DB again
 webservice=# select * from posts;
  id |       content       |  author
  ----+---------------------+----------
    1 | Random Post Content | Maverick



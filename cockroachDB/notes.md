Reading material


Spinning up one node cluster
` cockroach start-single-node --insecure --listen-addr=localhost:26257 --http-addr=localhost:8080`

`--insecure` - no authentication, no encryption


### starting the terminal
```
cockroach workload init movr
cocoroach sql --insecure
```

```
SHOW databases;
SHOW TABLES FROM movr;
SELECT * FROM movr.users LIMIT 10;
```

Primary Key

composite primary key - Multiple column primary key

```
create database <db_name>;

set database = <db_name>;

create table students(id uuid primary key default gen_random_uuid(), name string);

# to show schema
show create students;

create table courses (sys_id uuid default gen_random_uuid(), course_id int, name string, primary key(sys_id, course_id));


alter table courses add column schedule string;

create table products(id uuid default gen_random_uuid(), name string, quantity integer, price decimal);
```

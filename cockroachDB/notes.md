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



concepts to understand

1. Two phase commit
4. Truetime
2. Importance of Clock Synchronization in distributed systems.
3. External consistency
5. MVCC
6. CAP theorem - understand what C - consistency, A - availability and P -
Partion mean

### Cluster Fundamentals
The Keyspace, Ranges and Replicas

Keyspace - Basic model of cluster data

Ranges  - Keyspaces divided into ranges
Replicas - copies of ranges
Replicas distributed among nodes of cluster
Replication factor is 3 by default but it can be increased to any odd
number. eg: 5 or 7 etc

The Keyspace is divided into Ranges of 64MB each (by default). The
cluster creates multiple copies of each, call Replicas, which it
distributes among the nodes of a cluster.

### The raft protocol
- Cockroach DB uses Raft protocol to perform writes in a distributed and
durable manner.

- Raft is an alogrithm that allows a distributed set of servers to agree
  on any values without losing the record of that value, even in the
face of node failure.


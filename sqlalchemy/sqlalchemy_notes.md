### Introduction to SQLAlchemy (Mike Bayer)
source: https://www.youtube.com/watch?v=woKYyhLCcnU

#### Note: Most of the text is from the slides in the video.

SQLAlchemy consists of the Core and the ORM.
https://www.dropbox.com/s/riajyi0jr681xpn/Screenshot%202017-10-20%2013.10.15.png?dl=0

SQLAlchemy Core consists of:
* Schemes/Types
* SQL Expression Language
* Engine
  * Connection Pooling
  * Dialect


SQLAlchemy ORM:
  Built on top of the Core. ORM doesn't talk to DB directly. It goes
  through CORE apis


SQLAlchemy Core:
* Engine: a registry which provide connectivity to a particular DB
  server.

* Dialect: interprets generic SQL and db commands in terms of a specific
  DBAPI and database backend.
  eg: It knows how to product query for mysql and for oracle.

* Connection Pool: holds a collection of db connections in memory for
  fast-reuse.

* SQL Expression Language: allows SQL statements to be written using
  Python expressions.

* Schema/Types: uses Python objects to represent tables, columns, and
  datatypes.


SQLAlchemy ORM:
* Allows construction of Python objects which can be mapped to
  relational db tables.
* Transparently persists objects into their corresponding db tables
  using the unit of work pattern.
* Provides a query system which loads objects and attributes using SQL
  generated from mappings.

SQLAlchemy is like an Onion
  https://www.dropbox.com/s/nej9h2tyvgakxrj/Screenshot%202017-10-20%2013.19.26.png?dl=0



## Engine, Connection, Transactions

The job of Engine is to deal with he Python DBAPI (PEP-0249).
PEP - Python Enhancement Protocol

The first layer in SQL Alchemy is known as the Engine, which is the
object that maintains the classical DBAPI interaction.

slide_repl

```python
from sqlalchemy import create_engine
engine = create_engine("sqlite:///some.db")
result = engine.execute("select emp_id, emp_name from employee where
emp_id=:emp_id", emp_id=3)
row = result.fetchone()
print(row)
result.close()
result.fetchone() # will throw an error
```

With insert, update and delete operations SQLAlchemy auto-commits though
DBAPI doesn't have that in PEP spec. DBAPI wants to do it in a
transaction block.

Engine is not just a single connection, it is a connection pool. When we
execute a statement, engine connects to db, runs the query, gives the
result and close the connection or put it back in the connection pool.

To get a connection object
```python
conn = engine.connect()
result = conn.execute("select * from employee")
result.fetchall()
conn.close()
```

To do a transaction
```python
conn = engine.connect()
trans = conn.begin()
conn.execute("insert into employee (emp_name) values (:emp_name)",
emp_name="gp")
conn.execute("update employee_of_the_month set emp_name = :emp_name,
emp_name="gp")
trans.commit()
conn.close()
```


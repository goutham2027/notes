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



## Level1: Engine, Connection, Transactions

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
we need connection when we do transaction because the transaction
happens in one connecton at a time.

The short hand version of this is to use context manager.
```python
with engine.begin() as conn:
  conn.execute("<query>")
  conn.execute("<query>")
```

using bound parameters is the best practice. Prevents sql injection
attacks, and there is efficiency using bound practice.

#### Engine facts
* Executing via the Engine directly is call connectionless execution -
  the Engine connects and disconnects for us.
* Using a `Connection` is called explicit execution. We control the span
  of a connection in use.
* Engine usually uses a connection pool, which means "disconnecting"
  often means the connection is just returned to the pool.
* The SQL we send to engine.execute() as a string is not modified, is
  consumed by the DBAPI verbatim (in exactly the same words as used
  originally)

## Level2: Table, Metadata, Reflection, DDL
Metadata: Describes the structure of the db, i.e. tables, columns,
constraints interms of data structures in Python.
  Serves as the basis for SQL generation and object relational mapping.
  Can generate to a schema.
  Can be generate from a schema.

```python
from sqlalchemy import MetaData, Table, Column, Integer, String, ForeignKey
metadata = MetaData()
user_table = Table(
  'user', metadata,
  Column('id', Integer, primary_key=True),
  Column('name', String),
  Column('fullname', String),
)
# to create table in the db
metadata.create_all(engine)

# Foreign key
address_table = Table(
  'address', metadata,
  Column('id', Integer, primary_key=True),
  Column('email_address', String),
  Column('user_id', Integer, ForeignKey('user.id')),
)
```

### Reflection
Go check the db and fetch the columns from the DB.
```python
metdadata = Metadata()
user_reflected = Table('user', metadata, autoload=True, autoload_with=engine)
# output
print(user_reflected)
Table('user', MetaData(bind=None), Column('id', INTEGER(), table=<user>, primary_key=True, nullable=False), Column('name', VARCHAR(), table=<user>), Column('fullname', VARCHAR(), table=<user>), schema=None)
```

### Inspector object - if we want more specific information
```python
from sqlalchemy import inspect
inspector = inspect(engine)
inspector.get_table_names()
inspector.get_columns('address')
```

Basic Types
Integer - INT
String - VARCHAR
Unicode - VARCHAR, NVARCHAR
Boolean - BOOLEAN, INT, TINYINT
DateTime - DATETIME or TIMESTAMP
Float - floating point values
Numeric - precision numbers using Python Decimal

```python
# metadata create_all and drop_all will sort the correct hierarchy
# incase of foreign keys. The otherway to do is drop all constraints,
# but not all DBs support that.
metadata.create_all(engine, checkfirst=<True|False>)
table.create(engine, checkfirst=<True|False>)
metadata.drop_all(engine, checkfirst=<True|False>)
table.drop(engine, checkfirst=<True|False>)
```

### SQL Expressions
The SQL Expression system build upon Table Metadata in order to compose
SQL statements in Python.

start from 1:04

Google BigQuery
try_gcloud @.in

bq uses capacitor columnar data format. It offers the standard db
concepts of tables, partitions, columns and rows.
bq uses columnar storage that supports semistructured data - nested and
repeated fields.


loading data into BQ storage
- batch loads
- streaming


operations on bq
- copying tables
- querying tables using sql
- modifying data using sql dml
- exporting data

bq also supports querying data that's not in BQ's storage.

A bq slot is a unit of computational capacity required to execute SQL
queries. bq automatically calculates how many slots are required by each
query, depending on query size and complexity.

- run interactive queries
- run batch queries
- create a view, which is a virtual table defined by a sql query
- use partition tables to query a subset of the data.

bq table structure

bq jobs:
  jobs are actions that bq executes: load data, export data, query data
  or copy data.

  to run a bigquery job programmatically using the REST APIs or client
  libraries
  - call the jobs.insert method using a unique job id that is generated
    by client code
  - periodically request the jobs resource and examine the status
    property to learn when the job is complete



`pip install --upgrade google-cloud-bigquery`

resources:
https://www.reddit.com/r/bigquery/
https://www.kaggle.com/dansbecker/getting-started-with-sql-and-bigquery
https://www.kaggle.com/dansbecker/getting-started-with-sql-and-bigquery

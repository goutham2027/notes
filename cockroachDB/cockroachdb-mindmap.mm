<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1588651559274" ID="ID_1317463009" MODIFIED="1588651835321" TEXT="Distributed SQL">
<node CREATED="1588651574225" ID="ID_1911481265" MODIFIED="1588652561708" POSITION="right" TEXT="Legacy SQL">
<node CREATED="1588651966128" ID="ID_227417191" MODIFIED="1588652655627">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="4">Pros </font>
    </p>
    <p>
      <font size="4">- Consistency </font>
    </p>
    <p>
      <font size="4">- SQL - ACID transactions</font>
    </p>
    <p>
      <font size="4">Cons </font>
    </p>
    <p>
      <font size="4">- Resiliency is not smooth </font>
    </p>
    <p>
      <font size="4">- Scale is not easy</font>
    </p>
  </body>
</html>
</richcontent>
<node CREATED="1588652676140" ID="ID_558599886" MODIFIED="1588652954067">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <b><font size="3">ACID Transactions </font></b>
    </p>
    <p>
      
    </p>
    <p>
      <font size="3">Atomicity: transactions are all or nothin. If any part of the transaction fails, the entire transaction is aborted and the database is left unchanged. </font>
    </p>
    <p>
      
    </p>
    <p>
      <font size="3">Consistency: enforces rules on data. Any data written into DB will follow rules of SQL. </font>
    </p>
    <p>
      
    </p>
    <p>
      <font size="3">Isolation: ensures that concurrent transactions will leave the database in the same state just like if the transactions were executed serially, one after another </font>
    </p>
    <p>
      
    </p>
    <p>
      <font size="3">Durable: guarantees that once a transaction is committed, it will remain so even in the event of power loss or node failure or crashes.</font>
    </p>
  </body>
</html>
</richcontent>
</node>
</node>
<node CREATED="1588652336418" ID="ID_605535641" LINK="https://www.seas.upenn.edu/~zives/03f/cis550/codd.pdf" MODIFIED="1588652425187" TEXT="- Paper on Relational model of data for large shared data banks by Codd. https://www.seas.upenn.edu/~zives/03f/cis550/codd.pdf&#xa;- This has laid the foundation of RDBMS software"/>
</node>
<node CREATED="1588651939580" ID="ID_1662696774" MODIFIED="1588651959179" POSITION="right" TEXT="NoSQL">
<node CREATED="1588652970580" ID="ID_1565644638" MODIFIED="1588653002366">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="3">Pros </font>
    </p>
    <p>
      <font size="3">Scale </font>
    </p>
    <p>
      <font size="3">Resiliency </font>
    </p>
    <p>
      
    </p>
    <p>
      <font size="3">Cons </font>
    </p>
    <p>
      <font size="3">Consistency. Eventual consistency happens </font>
    </p>
    <p>
      <font size="3">SQL - ACID properties</font>
    </p>
  </body>
</html>
</richcontent>
</node>
</node>
<node CREATED="1588652094393" ID="ID_1898589251" MODIFIED="1588673218803" POSITION="right" TEXT="Examples: CockroachDB, Spanner &#xa;- Can handle OLTP workloads &#xa;- Not  legacy SQL &#xa;- Not NoSQL &#xa;- Combines the best of both worlds"/>
<node CREATED="1588665115200" ID="ID_982535423" MODIFIED="1588665130255" POSITION="right" TEXT="CockroachDB">
<node CREATED="1588673136700" ID="ID_323921976" MODIFIED="1588673203298">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      - Google published Spanner paper in 2012
    </p>
    <p>
      
    </p>
    <p>
      - CockroachDB started open source development in 2014
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588665144074" ID="ID_801561291" MODIFIED="1588665258511">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="3">Serializable Isolation </font>
    </p>
    <p>
      <font size="3">Every transaction behaves as if it had exclusive </font>
    </p>
    <p>
      <font size="3">use of the entire cluster as it performed transaction</font>
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588665263699" ID="ID_295678719" MODIFIED="1588665302843">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      The PostgreSQL wire protocol
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588665304172" ID="ID_641834341" MODIFIED="1588665359693">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="3">Geo-Replication </font>
    </p>
    <p>
      <font size="3">Ability to control where data resides, </font>
    </p>
    <p>
      <font size="3">even if it's a globally distributed cluster.</font>
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588665391700" ID="ID_89537091" MODIFIED="1588665473984">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="3">Multi-Cloud </font>
    </p>
    <p>
      <font size="3">Not tied to specific cloud provider </font>
    </p>
    <p>
      <font size="3">Using Spanner you will be locked into Google Cloud</font>
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588665484521" ID="ID_357714280" MODIFIED="1588665612665">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="3">Advanced Admin tools </font>
    </p>
    <p>
      <font size="3">Optimization </font>
    </p>
    <p>
      <font size="3">Baked in security: user authorization, authentication, audit logging</font>
    </p>
  </body>
</html>
</richcontent>
</node>
</node>
<node CREATED="1588665642677" ID="ID_551412747" MODIFIED="1588665844839" POSITION="right" TEXT="Notes">
<node CREATED="1588665649833" ID="ID_1561635174" LINK="https://www.cockroachlabs.com/blog/what-is-distributed-sql/" MODIFIED="1588665691186">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      What is distributed SQL
    </p>
    <p>
      https://www.cockroachlabs.com/blog/what-is-distributed-sql/
    </p>
  </body>
</html></richcontent>
</node>
<node CREATED="1588665695068" ID="ID_1734960813" LINK="https://www.cockroachlabs.com/webinars/why-distributedsql" MODIFIED="1588665745071">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="3">Webinar: Why Distributed SQL? </font>
    </p>
    <p>
      <a charset="utf-8" href="https://www.cockroachlabs.com/webinars/why-distributedsql">https://www.cockroachlabs.com/webinars/why-distributedsql</a>
    </p>
  </body>
</html></richcontent>
</node>
<node CREATED="1588665757471" ID="ID_1283836757" LINK="https://dl.acm.org/doi/10.1145/2168836.2168853" MODIFIED="1588665775001">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      Snapshot Isolation
    </p>
    <p>
      <a charset="utf-8" href="https://dl.acm.org/doi/10.1145/2168836.2168853">https://dl.acm.org/doi/10.1145/2168836.2168853</a>
    </p>
  </body>
</html></richcontent>
</node>
<node CREATED="1588665785412" ID="ID_331573651" LINK="https://www.youtube.com/watch?v=6OFeuNy39Qg" MODIFIED="1588672732061">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      The how's and why's of a&#160;Distributed SQL DB
    </p>
    <p>
      <a charset="utf-8" href="https://www.youtube.com/watch?v=6OFeuNy39Qg">https://www.youtube.com/watch?v=6OFeuNy39Qg</a>
    </p>
  </body>
</html></richcontent>
<node CREATED="1588666180543" ID="ID_444749586" MODIFIED="1588666465010">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      <font size="3">1960s - The first DB </font>
    </p>
    <p>
      <font size="3">- Records stored in pointers to each other in a hierarchy or network. </font>
    </p>
    <p>
      <font size="3">- Queries had to process one tuple at a time, traversing the structure </font>
    </p>
    <p>
      <font size="3">- To optimize queries the entire data has to be reloaded. </font>
    </p>
    <p>
      <font size="3">- For programmers it is tough to use</font>
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588666467263" ID="ID_357547992" MODIFIED="1588666594914">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      1970s - SQL/RDBMS
    </p>
    <p>
      - Codd - mathematician - IBM
    </p>
    <p>
      - Codd came up with relational model
    </p>
    <p>
      - Developer friendly
    </p>
    <p>
      - Queries are independent of underlying physical storage.
    </p>
    <p>
      - Designed to run on a single machine.
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588666611212" ID="ID_1263667520" MODIFIED="1588672589430">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      1980s - 1990s - Continued growth of SQL/RDBMS
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588672595824" ID="ID_949970672" MODIFIED="1588672770865">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      Early 2000s - Custom Sharding
    </p>
    <p>
      - no cross-shard transactions, operational nightmares
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588672775914" ID="ID_1353461629" MODIFIED="1588672836910">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      2004+&#160;- NoSQL
    </p>
    <p>
      
    </p>
    <p>
      - Scale and availability above all
    </p>
    <p>
      - Sacrificed: no relational model, no transactions and indexes, manual joins
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588672895226" ID="ID_1778220203" MODIFIED="1588673093958">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      2010s - New SQL/&#160;Distributed SQL
    </p>
    <p>
      - Provide full SQL semantics
    </p>
    <p>
      - Attempt to combine the best of both worlds
    </p>
    <p>
      &#160;1. Fully-featured SQL
    </p>
    <p>
      &#160;2. Horizontally scalable and highly available
    </p>
  </body>
</html>
</richcontent>
<node CREATED="1588673350769" ID="ID_1217749648" MODIFIED="1588673354587" TEXT="Data distribution">
<node CREATED="1588673358978" ID="ID_270502318" MODIFIED="1588673481747">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      SQL
    </p>
    <p>
      
    </p>
    <p>
      option1: All data in one server
    </p>
    <p>
      &#160;- secondary replicas/backups
    </p>
    <p>
      option2: Manually shard data across separate database instances
    </p>
    <p>
      &#160;- all data for each shard is still just on one machine
    </p>
    <p>
      &#160;- all data distribution choices are being made by DB admin
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588673505225" ID="ID_120739765" MODIFIED="1588673616587">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      NoSQL/NewSQL
    </p>
    <p>
      
    </p>
    <p>
      - Core assumption: entire dataset doesn't fit on one machine
    </p>
    <p>
      
    </p>
    <p>
      - How do you divide data up?
    </p>
    <p>
      - How do you locate any particular piece of data?
    </p>
    <p>
      
    </p>
    <p>
      - 2 primary approaches: hashing and order-preserving.&#160;
    </p>
  </body>
</html>
</richcontent>
<node CREATED="1588673623559" ID="ID_1921427987" MODIFIED="1588673803895">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      Hashing
    </p>
    <p>
      - divide up data by hashing each key
    </p>
    <p>
      - deterministically map hashed keys to servers
    </p>
    <p>
      - Pros: easy to locate data by key
    </p>
    <p>
      - Cons: inefficient range scans - hash functions doesn't maintain ordering of the data
    </p>
    <p>
      
    </p>
    <p>
      Used in Amazon Dynamo DB, Cassandra by default
    </p>
  </body>
</html>
</richcontent>
</node>
<node CREATED="1588673805514" ID="ID_1211691814" MODIFIED="1588701540230">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      Order-Preserving
    </p>
    <p>
      https://www.dropbox.com/s/sbvey0mz8luehim/Screenshot%202020-05-05%2015.53.44.png?dl=0
    </p>
    <p>
      - NoSQL systems Bigtable and Hbase and NewSQL systems use order-preserving approach
    </p>
    <p>
      
    </p>
    <p>
      - Divide sorted key space up into ranges of approximately equal size chunks
    </p>
    <p>
      - Distribute these chunks across the machines
    </p>
    <p>
      - Pros: efficient scans, easy splitting
    </p>
    <p>
      - Cons: requires additional indexing on the range. Which range is in the cluster and where it is located.
    </p>
    <p>
      
    </p>
    <p>
      When do you split up the data?
    </p>
    <p>
      - Generally when the ranges get to 60-100 MB then the data is split.
    </p>
    <p>
      
    </p>
    <p>
      Placement
    </p>
    <p>
      - Each range is replicated to three or more nodes.
    </p>
    <p>
      
    </p>
    <p>
      Rebalancing
    </p>
    <p>
      - Adding a new (empty) node
    </p>
    <p>
      - A new replica is allocated, data is copied
    </p>
    <p>
      - The new replica is made live, replacing another.
    </p>
    <p>
      - The old inactive replica is deleted
    </p>
    <p>
      - Process continues until nodes are balanced.
    </p>
    <p>
      Before rebalancing: https://www.dropbox.com/s/jen2zjyuko8lnil/Screenshot%202020-05-05%2023.28.17.png?dl=0
    </p>
    <p>
      After rebalancing: https://www.dropbox.com/s/jtlav66fviu2rj6/Screenshot%202020-05-05%2023.27.49.png?dl=0
    </p>
    <p>
      
    </p>
    <p>
      Recovery
    </p>
    <p>
      - Incase of a node failure
    </p>
  </body>
</html>
</richcontent>
</node>
</node>
</node>
<node CREATED="1588701615624" ID="ID_1454249529" MODIFIED="1588701642847">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      Data replication
    </p>
    <p>
      Keeping remote copies in sync
    </p>
  </body>
</html>
</richcontent>
<node CREATED="1588701623095" ID="ID_762108770" MODIFIED="1588701627379" TEXT="SQL"/>
</node>
</node>
</node>
<node CREATED="1588665846336" ID="ID_1300573704" LINK="https://www.seas.upenn.edu/~zives/03f/cis550/codd.pdf" MODIFIED="1588665875837">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      Relational DB model paper by E.F. Codd
    </p>
    <p>
      <a charset="utf-8" href="https://www.seas.upenn.edu/~zives/03f/cis550/codd.pdf">https://www.seas.upenn.edu/~zives/03f/cis550/codd.pdf</a>
    </p>
  </body>
</html></richcontent>
</node>
</node>
</node>
</map>

class SRE implements DevOps
- Define Availability
- Level of Availability
- Plan in case of failure

SLI - Service Level Indicator
  - Request Latency
  - Batch throughput
  - Failures per request

  Product
    SRE

  95th percentile latency of homepage requests over past 5 mins < 300ms

SLO - Service Level Objectives
  - Binding target for a collection of SLIs

  Product, SRE

  95th percentile homepage SLI will succeed 99.9% over trailing year.

SLA - Service Level Agreement
  - Business agreement between a customer and service provider typically
    based on SLOs

  - Sales, Customer

  Service credits if 95th percentile homepage SLI succeeds less than
  99.5% over trailing year.


SLIs drive SLOs which inform SLAs

Day-1
terraform refresh - reconciles what terraform thinks the world looks like with the real world.

terraform plan - what it needs to do, desired configuration

terraform apply - executes the plan

when tf applies, it figures out the order of execution.

tf is intelligent enough where to perform parrallel operations vs sequential operations.

Day-2
We evolve our infrastructure. Day2 is similar to Day1.

terraform destroy - to destroy

### TF architecture
2 things

1. TF core
Responsible for two things:
    a. Takes tf configuration and takes tf view/state of the world.
    b. Dependency graph
2. Providers


https://www.youtube.com/watch?v=SLB_c_ayRMo

Define provider
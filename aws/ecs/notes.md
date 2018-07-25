### AWS ECS - ec2 container service

What makes up the ECS?
* clusters
  grouping of container instances
* container instances
  ec2 instances running the ecs agent and registered in a cluster
* task definitions
  description of an application with one or more container definitions
* scheduling
  how we get our tasks on the container instances
* services
  an ecs service allows us to run or maintain a # instances of a task
  defintion
* tasks
  an instance of a task definition
* containers
  a linux container created as part of a task

Setting up
* create a user (ecs-admin)
* create a group (ecsadmins) with ecs related info

linuxacademy elastic beanstalk course

### what is a virtual machine (VM)?
  * emulation of computer system type.
  * virtualization software allows you to set up one OS within
    another. Host and guest (vm) OSs share the same physical hardware. The
    virtual machine is isolated from Host hardware and has to
    communicate with it through Hypervisor.

### what is a container?
  * entirely isolated set of packages, libraries and/or applications
    that are completely independent from its surroundings.
  * to use more system resources efficiently.
  * share host OS

### container architecture (docker)
  * linux containers (LXC) are there from a long. Docker makes it easy to use
    linux containers.
  * dcoker is a client-server arch where both the daemon and client can
    be run on the same system or we can connect to Docker client with a
    remote Docker daemon
  * containers rest on top of single linux instance.
  * docker engine or LXC process is not an equivalent to hypervisor. It
    simply encapsulates process on the underlying system.

  other container projects:
  * FreeBSD - Jails
  * Sun - Zones
  * Google - lmctfy (let me contain that for you)
  * OpenVZ - from mainframe space

### Introduction to Docker
  * it is a tool that packages up an application and all its
    depnedencies in a virtual container so that it can be run on any
    linux system or distribution.
  * docker offers the ability to isolate applications, standardize the
    build and deployment process and to create standard, repeatable
    processes in software and infrastructure.

### Introduction to Elastic Beanstalk
  * allows quick development and easy management of applications in AWS
    without getting bogged down in the infrastructure details.
  * AWS EB allows us to upload an application while it automatically
    handles capacity provisioning, load balancing, monitoring and
    scaling for us.
  * no additional charges for EB service, charged for AWS resources
    (EC2, RDS or S3) used as part of EB application.

    steps:
    - create an application
    - upload version
    - launch environment
    - manage environment




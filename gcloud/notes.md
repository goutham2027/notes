## Google cloud networking
https://cloud.google.com/compute/docs/networking#before-you-begin

* A network is constrained to a single project; it cannot span projects.
  But a project can have multiple networks.

* Each instance has a an internal IP address drawn from the IP range of
  the subnetwork or network it is created in. Each instance can also
  have external IP address (ephemeral or static).

* Any communication between instances in different networks, even within
  the same project, must be through external IP address.

### Network Types

#### Legacy (non-subnet) network
* A legacy compute engine network has a single network IPv4 prefix range
  and a single gateway IP address for the whole network. The network is
  global in scope and spans all cloud regions

* In a legacy network, instance IP addresses are not grouped by region
  or zone. One IP address can appear in one region, and the following IP
  address can be in a different region.

* Instances in a region can have IP addresses that are not grouped in
  any way.

#### Subnet network
* A subnet network divides the global network in to regional subnets, each
with its own IPV4 prefix.

* Subnetwork networks are regional in scope, unlike legacy networks
  which are global in scope.

* Subnets in a network do not have to be contiguous.
Eg: one subnet in a network can have a range of 10.240.0.0/16 and
another can have a range of 192.168.0.0/16.

* As such, there is no overall network IP range or gateway address in a
  subnet network. Instead, each subnetwork has an IP range and gateway
  address.

* Subnetworks belonging to a single network must have non-overlapping
  ranges. A subnetwork belongs to only one network and each instance can
  only belong to one subnetwork.

* A subnet network can be of auto mode or custom mode. An auto mode n/w
  has one subnet per region, each with a predetermined IP range and
  gateway. These subnets are created automatically when we create the
  auto mode network, and each submnet has the same name as the overall
  network.

* A custom mode network has no subnets at creation. In order to create
  an instancein a custom mode network, we must first create a subnetwork
  in that region and specify its IP range. A custom mode 

* Traffic from the internet passes through a global switching function
  in the network, then down to individual instances.

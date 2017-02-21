```
gcloud auth list
gcloud config list project
gcloud config set project <PROJECT_ID>
```

jump host - an instance with access to all required tools and other
VMs.
```
gcloud compute instances create jumphost --scopes cloud-platform --zone us-central1-f --metadata startup-script-url=gs://nw101/startupscript.sh
```

To get the list of compute instances
```
gcloud compute instances list
```
The above script creates 6 vms.
us-vm1 (us-central1-f), us-vm2 (us-central1-f), us-vmc (us-central1-c), eu-vm (europe-west1-d), asia-vm (asia-east1-c), jumphost (us-central1-f)

To connect to VM using gcloud
```
gcloud config set project <project-name>
gcloud compute ssh --zone <vm-zone> [vm-name]
```

SSH to each of the vm instances and run the following commands except us-vm2(centos).
```
$ sudo apt-get -y update
$ sudo apt-get -y install traceroute mtr tcpdump iperf whois host dnsutils
```

for centOS vms use
```
$ sudo yum check-update
$ sudo yum -y install epel-release traceroute mtr tcpdump whois bind-utils
$ sudo yum -y install iperf
```

### Checking the latency between VMs
Tool: ping
Ping measures packet loss and latency.

packet loss: network does not reach it's target or there is an error on
the network path.

latency: the round trip time network packets take to get from one
host to the other and back. Lower latency improves user experience.
(anything over 100ms application delay is noticeable)

ping uses ICMP messages to test connectivity.
```
$ ping eu-vm
```

how latency is calculated?
Ideal speed of light in fiber - 202562km/s or 125866m/s
eg: us-vm1 to eu-vm
distance as the crow files (straight line distance) - 7197.57km
ideal latency ((7192.58 km/202562 km/s) * 1000 ms/s)) * 2 = 71.07 mx
Observed latency: 100.88 ms

the difference is due to a non-ideal path

```
$ ping -i0.2 us-vm2 #(sends a ping every 200ms)
$ sudo ping -i0.05 us-vm2 -c 1000 #(sends a ping every 50ms, 1000 times)
$ sudo ping -f -i0.05 us-vm2 #(flood ping, adds a dot for every sent packet, and removes one for every received packet)
$ # careful with flood ping without interval, it will send packets as fast as possible, which within the same zone is very fast
$ sudo ping -i0.05 us-vm2 -c 100 -s 1400 #(send larger packets, does it get slower?)
```

Start from here
https://codelabs.developers.google.com/codelabs/cloud-networking-101/index.html?index=..%2F..%2Findex#7

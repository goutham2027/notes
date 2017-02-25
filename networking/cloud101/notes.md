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

### Traceroute
Tool to trace the path between two hosts.
Traceroute shows all the layer3 (routing layer) hops between the hosts.
This is achieved by sending packets to the remote destination with
incresing TTL (Time to Live) value starting at 1.
The TTL field in the IP packet gets decreased by 1 at every router and
once the the TTL hits zero, the packet gets discarded and a "TTL
exceeded" ICMP message is passed to the sender.

Decreasing of TTL value is to avoid routing loops, as packets cannot
lopp continuously because the TTL field will eventually decrement at 0.
By default OS sets the TTL value to a high value (64, 128, 255).

Traceroute sends packets first with TTL value of 1, then TTL value of
2 and soon, causing these packets to expire at the first/second/etc
router in the path. It then takes the source IP/host of the ICMP TTL
exceeded message returned to show the name/IP of the intermediate hop.
Once the TTL is high enough the packet reaches the destination, and the
destination responds.

Linux sends UDP packets to a high, unused port.

`traceroute www.icann.org`

Understanding traceroute output.
* last hop on traceroute is not destination: true for all external
  examples. The reasons for this is that traceroute performs a reverse
  DNS lookup for every host in the path. The reverse lookup for the last
  host might not be implemented or might be different than the name
  given for the forward DNS.

* traceroute shows only stars at the end: this is means there is
  probably a firewall in-between blocking either the incoming UDP/ICMP
  packets or outgoing ICMP or both.

  Note: Traceroute originally named Matt's traceroute(MTR) combines the
  functions of the traceroute and ping programs.
  `mtr wikipedia.org`

* Multiple paths showing: traceroute always sends 3 packets with the
  same TTL and those might be routed over different paths.

* Traceroute shows stars in the middle: because a host in the middle
  might not respond correctly with TTL exceeded messages or those might
  be filtered somewhere on the way.

Traceroute to bad.horse is an easter egg and can be build with bunch of
public IPs. Ref: http://beaglenetworks.net/post/42707829171/star-wars-traceroute

Caveats when working with traceroute/mtr:
* traceroute only shows the route from source to destination hosts. To
  get a full picture, we will need to provide traceroute from the source
  to the destination as well as from destination to the source.

* High latency or even loss on intermediate hops do not necessarily
  indicate a problem. Many hardware routers treat packets destined
  for/orginating from the router in software, so they are slow, while
  packets passing through are forwareded in hardware.

* the number of hops is largely irrelevant and a high number of hops
  does not indicate a problem.

start from https://codelabs.developers.google.com/codelabs/cloud-networking-101/index.html?index=..%2F..%2Findex#9


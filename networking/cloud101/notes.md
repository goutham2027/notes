Note: Most of the notes is from Google Cloud Networking 101

### Ping
### Traceroute
### iperf
### TCPDump
### Firewall and load balancing in google cloud

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


### Performance testing with iperf
iPerf is to measure the maximum achievable bandwidth on IP networks.

iPerf supports tuning of various parameters related to timing, buffers,
and protocols.

Using iperf we can test the performance between two hosts.

To test make one node as iperf server to accept connections.
```
# server mode
$ iperf -s

# to run in client mode and connect to server
$ iperf -c <iperf server>
```

We can increase bandwidth between hosts by using UDP.
```
# on server
$ iperf -s -u

# on client
$ iperf -c <iperf server> -u -b 2G
```

Higher speeds can be achieved by running a bunch of TCP iperfs in
parallel. The combined bandwidth should be really close to the maximum
achievable bandwidth.
```
$ iperf -s

# on client
$ iperf -c <iperf server> -P 20
```

To reach the maximum bandwidth, it's not sufficient by running a single
TCP stream (eg: file copy), we need to have several TCP sessions in
parallel.
Explore TCP parameters: window size and slow start.

With bbcp (https://github.com/eeertekin/bbcp) we can copy files as fast as possible by parallelizing
transfers and window size is also configureable.


### TCPDump
TCPDump captures network traffic(packets). Useful in debugging network
issues.

```
# collects 1000 packets.
$ sudo tcpdump -c 1000 -i eth0 not tcp port 22
```

Good practice to pass -c parameter

```
# packet capture for http request.
sudo tcpdump -i eth0 -s 1460 -w webserver.pacp tcp port 80

# analyzing the packet capture file
# shows basic protocol, source and destination.
# wireshark and cloudshard can be used to get more information
sudo tcpdump -nr webserver.pacp
```

### Firewall and load balancing in google cloud

```
# to get gcloud compute instances
$ gcloud compute instances list
```

Lab exercise: Install nginx on  us-vmc compute instance and run it on
port 81.
To run on port 81 use: `echo "server { listen 81; root /usr/share/nginx/html; }" > /etc/nginx/sites-enabled/default`

But from external hosts us-vmc ip is not reachable. Update the firewall
rule.

Solution:
To open the GCE firewall, we need to give the following information:
* source ip range or tags (will will open for the internet 0.0.0.0/0)
* destination protocol and port (tcp:81)
* destination tags (we will have to create if we don't want to open port
  on all vms)
* network (codelab)

```
# to create the tag nginx-81
$ gcloud compute instances add-tags us-vmc --tags nginx-81 --zone us-central1-c
# to open the port
$ gcloud compute firewall-rules create nginx-81 --allow tcp:81 --network codelab --source
-ranges 0.0.0.0/0 --target-tags nginx-81
```

#### Network Load Balancer (lb)
The Network lb is a Layer3 lb which provides IP packets for a certain
port to a target pool of multiple instances. This is done by forwarding
rules. Forwarding rules can also forward traffic to a single instance
(which gives it additional external IPs without NAT).

Network lb is regional, it can span multiple zones but not regions.

The traffic is forwarded as is and responses are sent from the instance
directly to the source of the connection. For this the load
balancer/forwarding rule IP gets added to the instances which are the
target for the load balancer.

Exercise
```
# installing a web server (on jumphost)
 $ for a in us-vm1 us-vmc eu-vm asia-vm ; do gcloud compute ssh --command "sudo apt-get -y install apache2" --zone `gcloud compute instances list $a --format text | grep zone: | awk '{print $2}'` $a; done

 # for centos (on jumphost)
 $ gcloud compute ssh --ssh-flag="-t" --command "sudo yum -y install httpd; sudo service httpd start" --zone us-central1-f us-vm2

 # to verify apache2 is installed
 $ for a in us-vm1 us-vm2 us-vmc eu-vm asia-vm ; do curl $a; done

 # before creating the US loadbalancer we need to open the firewall to
 all instances.
  $ gcloud compute firewall-rules create http --allow tcp:80 --network codelab
```

A network lb consists of a
 * target pool of machines in one region. In
our example we will be creating a lb for all US instances.
 * a http health check to check if those machines are healthy (default
   check on port 80 and path /)
 * forwarding rule that gives us external ip pointing at this pool of
   instances.

```
$ gcloud compute http-health-checks create basic-check
$ gcloud compute target-pools create apaches --region us-central1 --health-check basic-check
$ gcloud compute target-pools add-instances apaches --instances us-vm1,us-vm2 --zone us-central1-f
$ gcloud compute target-pools add-instances apaches --instances us-vmc --zone us-central1-c
$ gcloud compute forwarding-rules create interwebs --region us-central1 --port-range 80 --target-pool apaches
```

The last command returns the ip address which is the load balancer ip
address.

```
# to modify index.html to display IP address
 for a in us-vm1 us-vmc eu-vm asia-vm ; do gcloud compute ssh --command "sudo hostname | sudo tee /var/www/html/index.html > /dev/null" --zone `gcloud compute instances list $a --format text | grep zone: | awk '{print $2}'` $a; done
```

exercise create the n/w load balancer for both eu-vm and asia-vm
```
# to directly add a forwarding rule to the single instance
$ gcloud compute target-instances create eu-target --instance eu-vm --zone europe-west1-d
$ gcloud compute forwarding-rules create interwebs-eu --region europe-west1 --port-range 80 --target-instance eu-target --target-instance-zone europe-west1-d
```

#### Notes on HTTP load balancer:
The HTTP load balancer is a layer 7 lb. the tcp session gets terminated
by the lb and proxies the traffic to the target instances. responses also get proxied through the lb.
layer 7 lb can be used as a global lb where packets are always forwarded
to the closes region

http load balancer currently does not support session affinity. what is
session affinity?


Setting up global load balancer across all VMs
https://www.dropbox.com/s/w09jldcm4mb6t7l/Screenshot%202017-03-07%2009.13.47.png?dl=0

To achieve this we need to build the system from the target to the
beginning.

step-1: setup an instance group for every zone, as the HTTP load
balancer balances between instance groups. we are using unmanaged
(non-autoscaled_ instance group.
`gcloud compute instance-groups unmanaged create us-f --zone us-central1-f`

step-2: add instances to the instance group
`gcloud compute instance-groups unmanaged add-instances us-f --instances us-vm1,us-vm2 --zone us-central1-f`

step-3: add named port (http) to each newly created instance group.
`gcloud compute instance-groups unmanaged set-named-ports us-f --named-ports http:80 --zone us-central1-f`

step-4: create a backend service containing all those instance groups
and add an URL map pointing all URLs to this backend service.
```
gcloud compute backend-services create global-bs --protocol http --http-health-check basic-check
gcloud compute backend-services add-backend global-bs --instance-group us-f --zone us-central1-f
gcloud compute url-maps create global-map --default-service global-bs
```

step-5: finally we create a target proxy and point a forwarding rule
with a global ip to the target proxy
```
$ gcloud compute target-http-proxies create global-proxy --url-map global-map
$ gcloud compute forwarding-rules create global-lb --global --target-http-proxy global-proxy --ports 80
```

The last command returns the IP. This IP will take user to the closest
instance.

By default the logs shows the lb ip. to surface the ips in the log run
the following commands.
```
# for debian based
$ sed -e 's/%h/%{X-Forwarded-For}i/' /etc/apache2/apache2.conf | sudo tee /etc/apache2/apache2.conf > /dev/null
$ sudo service apache2 reload

# for centos based
$ sudo sed -e 's/%h/%{X-Forwarded-For}i/' /etc/httpd/conf/httpd.conf | sudo tee /etc/httpd/conf/httpd.conf > /dev/null
$ sudo service httpd reload
```

### https load balancer
http lb can be converted to https using a certificate.
steps to create a self signed certificate
```
$ mkdir ssl
$ cd ssl
$ openssl genrsa -out my.key 2048
$ openssl req -new -key my.key -out my.csr #Enter data at each prompt except PW
$ openssl x509 -req -days 365 -in my.csr -signkey my.key -out my.crt

$ gcloud compute ssl-certificates create ssl --certificate my.crt --private-key my.key

# create target https proxy and a forwarding rule
$ gcloud compute target-https-proxies create ssl-proxy --url-map global-map --ssl-certificate ssl
$ gcloud compute forwarding-rules create ssl-lb --global --target-https-proxy ssl-proxy --ports 443
```



### GCE specific settings to instances
TCP keepalive settings: default: 10 mins
To change this use the following command.
`$ sudo /sbin/sysctl -w net.ipv4.tcp_keepalive_time=60 net.ipv4.tcp_keepalive_intvl=60 net.ipv4.tcp_keepalive_probes=5`


cleanup script source <(curl -s https://storage.googleapis.com/nw101/cleanup.sh)

google cloud networking: https://cloud.google.com/compute/docs/networking?hl=en

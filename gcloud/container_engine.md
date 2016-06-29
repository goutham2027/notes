### container cluster
A container engine cluster is a group of compute engine instances
running Kubernetes. It consits of one or more node instances and a
managed kuberenetes master endpoint.

A container cluster is the foundation of a container engine application
- pods, services and replication controllers all run on top of a
  cluster.

Kubernetes master
 Every container cluster has a single master endpoint, which is managed
 by container engine. THe master provides a unified view into the
 cluster.

 The managed master also runs the Kubernetes API server which services
 REST requests, schedules pod creation and deletion on worker nodes and
 synchronizes pod information.

 Container Engine also upgrades the master to future kubernetes
 versions.

Nodes: A container cluster can have one ore more node instances. These
are managed from the master and run the services necessary to support
docker containers.

Each node runs the Docker runtime and hosts a Kubelet agent, which
manages docker containers scheduled on the host.

Each node also runs a simple network proxy.

### creating container clusters
 we can create from cloud console or from gcloud command line

```
# switching between clusters
gcloud config set container/cluster NAME
gcloud container clusters list
gcloud container clusters get-credentials NAME
gcloud container clusters describe NAME --zone ZONE
gcloud container clusters delete NAME --zone ZONE
```

Once the cluster has been created, use kubectl commands to create and
manage resources on the cluster. Get credentials before running kubectl
commands.

kubectl commands
```
kubectl get deployments
kubectl logs <pod-name>
kubectl cluster-info
```

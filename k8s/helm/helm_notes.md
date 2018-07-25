## Helm

Helm - Package Manager (like HomeBrew/apt/yum) for K8s.

### Helm components
Client - Command line.
`brew install kubernetes-helm`
Tiller - Stays in the cluster. Manages, releases, history and introspection.
  `helm init` will install tiller to the cluster kube-system namespace.
  `helm init` will install to the current kube context. We can specify
  context using `--kube-context`.
  `kubectl get pods --namespace kube-system` - we can see Tiller
  running.

`helm version` - will tell both the client and server version.

### 3 Big Concepts
* Chart
  - Packaged K8s resources.
  - A chart is a Helm package.
  - Contains all the resources definitions necessary to run an
    application, tool or service inside of a Kubernetes cluster.
  - equivalent to Homebrew formula, Apt dpkg
* Chart Repository
  - Registry of consumable charts. (like Dockerhub)
  - place where charts can be collected and shared.
* Release
  - An instance of a chart running in a Kubernetes cluster.
  - Once chart can be installed many times into the same cluster.
    Each time it is installed, a new release is created.
  - A deployed instance of a chart. Install the same chart it's
  a new release.
  - eg: A MySQL chart. If we want 2 databases running in the cluster, we
    can install that chart twice. Each one will have its own release,
    which will inturn have its own release name.

```
helm search
helm search mysql
helm instpect stable/mariadb

helm install stable/mariadb
helm status <release name>
helm ls # to list releases
```

### Charts
Helm uses a packaging format called charts. A chart is a collection of
files that describe set of Kubernetes resources.

* A chart is organized as a collection of files inside of a directory.
  The directory name is the name of the chart.

eg: notesapp
```
notesapp/
  Chart.yaml
  requirements.yaml # to list dependencies for the Chart.
  values.yaml # default configuration values for this chart
  charts/ # a directory containing any charts upon which this chart depends.
  templates/ # a directory of templates that, when combined with values,
             # will generate valid Kubernetes manifest files.

```

After specifying dependencies use `helm dependency update notesapp`,
this command will download all the specified charts into `charts/`
directory.

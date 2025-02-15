# Dope

Deployment oriented programming. `Dope` is a highly Opinionated framework for developing golang microservices on kubernetes environment. The `Dope` framework dictates a strictly standardized stricture for golang microservices, and in return promises to provide a fast shipping of code to production, by handling all the tedious work required between a code that is written locally to a service deployed on kubernetes. 

### Key features:
* *Build* - `Dope` would be responsible to build a valid helm chart to each service, given only a basic configuration.
* *Environments consistency* - `Dope` integrates your IaC with our code by default, allows managing your local environment(s), development environments and production environment without additional set-up.
* *Debug* - Debugging containers on remote environments is not always a pleasure. with `Dope` it's as easy as switch a parameter to `true` for a service in the environment variable. 
* *Tests and docs* - As `Dope` standardizes the code base, it can easily create unit tests, e2e tests and open-api specs for documentation.
* *CI/CD* - though not enabled by default, `Dope` would gladly handle the CI/CD workflows for you. This feature requires an integration with `ArgoCD`.


## Installation

Install the cli (linux):

```bash
$ go install github.com/shaharby7/Dope/cmd/dopecli@latest
$ mv $GOPATH/bin/dopecli /usr/local/bin/dope
```

## Quick start

### Preretirement:

* up and running kubernetes cluster for testing ([minikube](https://minikube.sigs.k8s.io/docs/) would do the job!)
* [kubectl](https://kubernetes.io/docs/reference/kubectl/) installed and  configured to the same cluster
* [helm cli](https://helm.sh/docs/intro/install/) installed

### Steps

1. To create a new project run a `create` command with `project` variable, and potentially add custom path. You would be asked to give a name to the project: 
```shell
$ dope create project [path]
```

Output:
```
Name: test
Description: test
```

2. Create your first application! 

```shell
$ dope create app
```

Go to the newly created file and add values to your first app, for example:

```yaml
api: Dope
type: App
name: "myapp"
description: "example app"
values:
  version: 6
  controllers:
  - name: "server1"
    description: "http sever controlled by app1"
    type: HTTPServer
    actions:
    - name: "/api/greet"
      description: "Great anyone who wants"
      package: pkg/greeter
      ref: Greet
      controllerBinding:
        method: POST
  requirements:
    env:
    - "UGLY_NAMES"
```

3. Add your first environment 

```shell
$ dope create env
```

Go to the newly created file and add values to your first environment, for example:

```yaml
api: Dope
type: Env
name: local
description: "local deployment example"
values:
  providers:
    git:
      url: https://github.com/shaharby7/Dope.git
      path: example
      ref: shahar
    kubernetes:
      type: minikube
    storage:
      type: minio
      managed: true
    cd:
      type: argo-cd
      managed: true
      values:
        crds:
          install: false
        redis-ha:
          enabled: false
        controller:
          replicas: 1
        server:
          replicas: 1
        repoServer:
          replicas: 1
        applicationSet:
          replicaCount: 1
    workflows:
      type: argo-workflows
      managed: false
```

5. Define how the application should be deployed on this environment:
```shell
$ dope create appenv
Name: app-local
✔ Description: test deployment█
✔ Env: local█
App: myapp
```

Go to the newly created file and add values to your first app environment, for example:

```yaml
api: Dope
type: AppEnv
name: "local-myapp"
binding:
  Env: local
  App: myapp
values:
  registry: "shaharby7/app1-local"
  values:
    serviceAccount:
      create: true
      automount: true
      annotations: {}
  controllers:
  - name: "server1"
    env:
    - name: UGLY_NAMES # overrides the value from the defaults
      value: "shahar,danny"
    - name: FAMOUS_NAMES # creates new env var for this controller only
      value: "donald"
    replicas: 3
    resources:
      requests:
        cpu: 1 # overrides the default value only for requests.cpu
  controllersDefaults:
    env:
    - name: FORBIDDEN_NAMES
      value: "nice"
    - name: UGLY_NAMES
      value: "shahar"
    resources:
      requests:
        cpu: 2
        memory: "2Gi"
      limits:
        cpu: 2
        memory: "4Gi"

```

5. Build! 
```
$ dope build 
```

6. Add `dope` helm repo:
```shell
$ helm repo add dope https://shaharby7.github.io/Dope/helm/charts
```

7. Deploy your application!:
```shell
dope install
```

8. Review your deployment and get your greetings! 
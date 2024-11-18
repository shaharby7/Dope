# Dope

Deployment oriented programming. `Dope` is a highly Opinionated framework for developing golang microservices on kubernetes environment. The `Dope` framework dictates a strictly standardized stricture for golang microservices, and in return promises to provide a fast shipping of code to production, by handling all the tedious work required between a code that is written locally to a service deployed on kubernetes. 

### Key features:
* *Build* - `Dope` would be responsible to build a valid helm chart to each service, given only a basic configuration.
* *Environments consistency* - `Dope` integrates your IaC with our code by default, allows managing your local environment(s), development environments and production environment without additional set-up.
* *Debug* - Debugging containers on remote environments is not always a pleasure. with `Dope` it's as easy as switch a parameter to `true` for a service in the environment variable. 
* *Tests and docs* - As `Dope` standardizes the code base, it can easily create unit tests, e2e tests and open-api specs for documentation.
* *CI/CD* - though not enabled by default, `Dope` would gladly handle the CI/CD workflows for you. This feature requires an integration with `ArgoCD`.


## Quick start

### Preretirement:

* up and running kubernetes cluster for testing ([minikube](https://minikube.sigs.k8s.io/docs/) would do the job!)
* [kubectl](https://kubernetes.io/docs/reference/kubectl/) installed and  configured to the same cluster
* [helm cli](https://helm.sh/docs/intro/install/) installed

### Steps

1. Download the `dope-cli`:
```shell
$ go install github.com/shaharby7/Dope@latest
```

2. Generate a new project:
```shell
$ dope create /path/to/project
Please type the name of your project? myproj
Please type the domain for your project (github.com by default)? 
```

3. cd to the project and open it with your favorite IDE.
4. Create your first application! Open the `project.dope.yaml` file and add to the `apps` node your first app:

```yaml
apps:
- name: "greater"
  description: "app example"
  controllers:
  - name: "server1"
    description: "http sever controlled by greater"
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

5. add to the `environments` your first environment:
```yaml
environments:
- name: local
  apps:
  - name: "greater"
    repository: docker.io/shaharby7/greater-local
    env:
    - name: UGLY_NAMES
      value: "shahar,danny"
    resources: 
      requests:
        cpu: 1
        memory: 2GB
      limits:
        cpu: 2
        memory: 4GB
```

6. Build! 
```
$ dope build 
```
Note - `dope-cli` would search for `./project.dope.yaml` file by default, and would locate all the outputs in a `./build` directory by default, both can be set manually, see the `dope-cli` docs (#todo) 


7. Review the output helm charts and values in the `./build` directory

8. Add `dope` helm repo:
```shell
$ helm repo add dope https://shaharby7.github.io/Dope/helm/charts
```

8. Deploy your application!:
```shell
helm install dope dope/dope -n dope -f ./example/build/helm/local/dope/values.yaml  --create-namespace
```

9. Review your deployment and get your greetings! (#todo)
```
$ 
```
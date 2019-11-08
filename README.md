# kubist-example-ci-integratation

This repository is an example on how to create regular releases with vamp kubist and the readme explains how to deploy a http frontend service.

# Prerequisites

For the simplicity of this example;
For project name, we will use project1
For cluster name, we will use cluster1
For user name, we will use user1

* A user who authorised as admin in the project1 on a vamp kubist cluster or service
* A cluster user has access to and authorised kubectl installed locally.
* Latest version of vamp kubist client is locally installed.
* helm is locally installed.
* curl for http requests

To install a new vamp kubist cluster for yourself check this documentation:
https://github.com/magneticio/vampkubistdocs/blob/master/INSTALLATION.md

# Setting up cluster and namespace
Login as user1
```shell
$ vamp login --user user1
```

Validate if you have access to the project by listing projects
```shell
$ vamp list projects
- project1
```

Set context to project1
```shell
$ vamp config set -p project1
```

Bootstrap your cluster as cluster1 in project1
```shell
$ vamp bootstrap cluster cluster1
```
Bootstrapping creates required credentials for vamp kubist to access to your cluster and creates a cluster with these credentials. It is a very short process.

Set context to cluster1
```shell
$ vamp config set -c cluster1
```

Create a kubernetes namespace for this example:
```shell
$ kubectl create ns kubist-example
```
Vamp kubist does not create or delete namespaces, it will detect when it is added to Vamp kubist.
To add a namespace to vamp kubist, create a virtual cluster with the same name as your namespace.
```shell
$ vamp create virtualcluster kubist-example --init
```

Now deploy first version of our application. This application has a helm template and we will use helm template and kubectl approach to deploy this application:
Note: It is important that resource names doesn't include dots ("."). In the example helm chart "." are replaced with "-".
```shell
$ helm template achart --set image.tag=v0.0.1 | kubectl apply -n kubist-example -f -
```


Set context to kubist-example
```shell
$ vamp config set -r kubist-example
```

We are using the image tag to define version.
This will create a deployment in kubist-example namespace with version v0.0.1

Create a destination for your service
```shell
$ vamp create destination kubist-example-destination -f ./vamp/destination.yaml
```

```shell
$ vamp create gateway kubist-example-gateway -f ./vamp/gateway.yaml
```

Create a vampservice for your service
```shell
$ vamp create vampservice kubist-example-service -f ./vamp/vamp-service.yaml
```

You can get the ip adress of the gateway and check with curl if it is working.
You should see OK in the output.
```shell
$ GATEWAY_IP=$(vamp get gateway kubist-example-gateway -o=json --jsonpath '$.status.ip' --wait  --number-of-tries 20 | tr -d '"')
$ curl $GATEWAY_IP
OK
```
You can also copy paste the address in $GATEWAY_IP to a browser.

If you don't see OK. Check if gateway is still initialising.
There are two statuses, deploymentStatus and loadBalancerStatus. Both them should be ready.
```shell
$ vamp get gateway kubist-example-gateway -o=json --jsonpath '$.status.deploymentStatus'
"Ready"
$ vamp get gateway kubist-example-gateway -o=json --jsonpath '$.status.loadBalancerStatus'
"Ready"
```

# Circle Ci configuration
Before triggering a release you need to configure circleci
Get your user credentials by running:
```shell
$ vamp config get
```

It is recommended to create a service account to use with ci tools.
A documentation will be added on how to create service accounts.

You need these variables to set up in the circle ci as env variables:
* url should be stored as VAMP_URL
* token should be stored as VAMP_TOKEN
* cert should base64 encoded and stored as VAMP_CERT_BASE64

You will also need docker hub credentials to push your docker image.
This project assumes it as a public repo, if it is a private repo you need to set up your pull credentials while deploying your application.
These env variables are:
* DOCKER_USER
* DOCKER_PASS


# Trigger a release
Now you are ready to do a health based canary releasing.
We trigger a canary release when repo has a new git tag as version tag.
Version tags should follow vd.d.d where d is a decimal number.

We added a release.sh script to trigger a new release.
On master branch run:
```shell
./release.sh v0.0.7
```
If the version already exists it will fail and you need to increment the version to the version you want to deploy.

# Happy Releasing
Everything is automated after this point.
You will only need to trigger the last step for new a release.

# minimesos-golang

This will be a Golang client for the new minimesos API. See https://github.com/ContainerSolutions/minimesos/issues/452

# Status

Currently it is just some boilerplate code that prints some statements. It does not interact with the API yet.

# TODO

* Launch a `containersol/minimesos-api` container
* Call the /start endpoint
* Add tests
* Implement more methods

# Installation

```
go get gopkg.in/fsouza/go-dockerclient.v0
```

# Run

```
$ go run main.go
Created Mesos cluster. API running at port 9999
Started Mesos cluster.
Destroyed Mesos cluster.

```

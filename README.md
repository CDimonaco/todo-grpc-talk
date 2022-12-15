# TODO-GRPC-TALK - Intro to Protobuf & GRPC talk - Source code

[![Go Report Card](https://goreportcard.com/badge/github.com/CDimonaco/todo-grpc-talk)](https://goreportcard.com/report/github.com/CDimonaco/todo-grpc-talk)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)

![Main branch](https://github.com/CDimonaco/todo-grpc-talk/actions/workflows/github-actions-on-push.yaml/badge.svg)


## Installation

Checkout the latest release on github, you can find both server and client.

### Development

Requirements

- golang 
- task 
- protoc 
- protoc-gen-go 
- protolint
- protoc-gen-go-grpc 
- mockery 
- golangci-lint

You can install them using [asdf-vm]("https://asdf-vm.com/")

Install `asdf-vm` plugins

```
asdf plugin-add golangci-lint
asdf plugin-add mockery
asdf plugin-add protoc-gen-go-grpc
asdf plugin-add protolint
asdf plugin-add protoc-gen-go
asdf plugin-add protoc
asdf plugin-add task
asdf plugin-add golang
```

Install all the dependencies

```
asdf install
asdf reshim
```
- docker
- docker-compose
Task managent through `taskfile.dev`

```
task: Available tasks for this project:
* build-proto-stubs:                 Build protobuf and grpc stubs from proto definition files
* build-todo-grpc-talk-client:       Build todo-grpc-talk-client binary for release
* build-todo-grpc-talk-server:       Build todo-grpc-talk-server binary for release
* lint:                              lint the project
* start-grpcui:                      Start GRPCUI
```

### GRPC-UI

Use the task `start-grpcui` to start [grpcui]("https://github.com/fullstorydev/grpcui"), a webui to interact with grpc services, like postman or insomnia, this uses the `grpc introspection` which in this project is enabled by default.

You need `docker` and `docker-compose`
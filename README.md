# goloco

## Introduction
`goloco` is a sample application based on microservices architecture. Its in very intial stages and I use it as my playground application to experiment with microservices. 

## Microservices
There are 3 microservices running written in different languages communicating over GRPC.
- frontend (written in Go-lang)
- locationservice (written in Python)
- recommendationservice (written in Kotlin)

## Database
There are currently two databses that I added in the application
- Postgres (in an external cluster)
- Redis (dockerized and deployed in kubernetes cluster)

## Deployment
Skaffold: Application is deployed to Kubernetes with a single command using Skaffold.

## Installation
1. Install tools to run a Kubernetes cluster locally:

   - Local Kubernetes cluster deployment tool:
        - [Minikube](https://kubernetes.io/docs/setup/minikube/).
        - Docker for Desktop: It provides Kubernetes support as [noted
     here](https://docs.docker.com/docker-for-mac/kubernetes/).
   - [skaffold]( https://skaffold.dev/docs/install/)
   
   1. Run `kubectl get nodes` to verify you're connected to “Kubernetes on Docker”.

1. Run `skaffold run`
   This will build and deploy the application.

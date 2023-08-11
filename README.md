# tiktok

**tiktok** is a distributed **simple-tiktok** backend based on RPC and HTTP protocols using Kitex + Hertz + Etcd + MySQL + Jaeger + Docker + Thrift

# Basic operation

We use docker to create project's dev environment. So before everything, we need to start docker in project root:

```bash
docker-compose up -d # set environment
docker-compose down  # unset
```

Then, we should launch all of the different miceoservices.

We step in every microservice's folder: **./cmd**，in microservice's workspace, we start it by :

```bash
# FOR API
make     # build and serve
make new # create new hertz project
make gen # update existing hertz project

# FOR OTHERS
make        # build and serve
make gen    # create/update kitex projects
make test   # test your code
```
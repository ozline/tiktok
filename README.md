# tiktok

**tiktok** is a distributed **simple-tiktok** backend based on RPC and HTTP protocols using Kitex + Hertz + etcd + MySQL + Jaeger + Docker + Thrift + Prometheus + Grafana

# Feature

- Extremely easy to use and deploy.
- Relatively mature CI/CD.
- Relatively high code quality
- Safety Considerations
- Performance Optimization for Interfaces

# Quick start

We will introduce how to quickly start this project using Docker. If you need to build and run it locally, please refer to: [start-by-local](./docs/start-by-local.md)

Due to the script I have written, the process has been greatly simplified. You just need to use the following command to quickly start the environment and run the program in a containerized manner.

```bash
    make env-up      # launch environment, env-down for remove
    make docker      # build docker-image
    sh docker-run.sh # launch all services(carrying the service name allows you to start a specified service)
```

then you can send HTTP request on `localhost:10001` for test or others things

# Quick deploy

We use a fully automated process to streamline the workload, so you can always use our Docker image packaged with the latest code.

You need to **install and start Docker**, and at the same time, you need to **configure the corresponding server environment**. If it is just for testing deployment, you can directly use `docker-compose.yml` to start. This way, you don't need to modify the config file either.

```bash
    cd deploy             # or you can only move this dir to your server rather than git clone all codes
    mkdir -p config
    cd config
    touch config.yaml     # you can simply copy the examples in the config directory and make modifications according to the instructions inside.
    touch prometheus.yaml # same as above
    cd ..
    sh restart-service-all.sh # start all services
```

The script will automatically pull the latest image from Aliyun ACR, find and delete the running containers, and re-run them with the latest image.

Then you can send HTTP request on `localhost:10001` for test or others things

# Architecture

## Overall
```bash
.
├── Dockerfile
├── LICENSE
├── Makefile              # some useful commands
├── README.md
├── cmd                   # microservices
├── config                # for run-directly config and config-example
├── deploy                # for deploy
├── docker-compose.ci.yml # for ci env
├── docker-compose.yml
├── docker-run.sh         # for local docker-run
├── docs
├── go.mod
├── go.sum
├── idl                   # interface definition
├── kitex_gen
├── pkg
│   ├── constants         # store any consts
│   ├── errno             # custom error
│   ├── middleware        # common middleware
│   ├── tracer            # for jaeger
│   └── utils             # useful funcs
└── test
```

## Gateway/api service

```bash
.
├── Makefile
├── biz
│   ├── handler     # solve request/send response
│   ├── middleware
│   ├── model
│   ├── pack        # pack response
│   ├── router      # for route
│   └── rpc         # send rpc request to microservices
├── build.sh
├── main.go
├── output          # build binary
├── router.go
├── router_gen.go
└── script
```

## Microservices
```bash
.
├── Makefile        # useful commands
├── build.sh        # build binary
├── dal
│   ├── cache       # redis
│   ├── db          # MySQL
│   └── mq          # RabbitMQ
├── handler.go
├── kitex_info.yaml
├── main.go
├── output          # build binary
├── pack            # pack response
├── rpc             # send request to other services
├── script
├── coverage        # coverage test(some service not exist)
└── service
```


# Test interfaces

you can drop `.postman/tiktok.openapi.json` to **postman** then start this project and test

# Contribute & Question

you can send Pull Request(PR) for contribute, I will quickly response

If you have any questions, you can create an issue.
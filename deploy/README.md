# Deploy

All files in this directory are only used for CI/CD purposes. But you can also manually use these scripts or configurations in the right place.

# Introduction

## docker-entrypoint.sh

only use for docker build, **DO NOT EDIT OR MOVE.**

## restart-service.sh

this shell file is used for start/restart the specific service (**list: api user follow interaction video chat**)

```bash
sh restart-service.sh api # or others
```

- the container use **host** network
- you need to **move config.yaml** to the same dir as this file
- **Before starting the service**, you need to configure the corresponding environment (refer to the configuration in config.yaml).
- The script will **automatically** detect and delete the containers, no manual deletion is required.

## Others

- restart-service-all.sh: no further explanation is needed.
- common.sh: set any constants
- remove-all-containers: common scripts
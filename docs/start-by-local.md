# Qucik start by local

## Basic environment and config

We can use the docker-compose we wrote to quickly set up the runtime environment.

```bash
docker-compose up -d
```

After that, we need to update the config file. You just need to do the following two things:
1. Open `./config/config_example.yaml` and **complete** the configuration **according to the comments** in the file (This is not much, most of it we have already set up for you.).
2. Rename the above file as `config.yaml`.

## Build and run

We have greatly simplified the commands using Makefile, and all you need to do is:
1. We have 6 services: `api`, `user`, `follow`, `chat`, `video`, and `interaction`. To run a specific service, go to the root directory and execute it by **make**

```bash
make api # or others
```
2. Start each of the 6 services one by one (this may require opening 6 or more terminals).
3. Makefile will automatically help you build the binary program and the necessary script configuration, and move these artifacts to the `output` folder in the root directory. Finally, it will automatically run them for you.

## Test & send request by yourself

To facilitate your testing (even though it does not comply with go test specifications), we have placed all the test files in the `test` directory under the root directory. You only need to run the following program to let go test automatically execute the tests for you.

```bash
go test ./test/...
```

you can also send request by yourself, just step in `.postman` folder and import json file into postman!
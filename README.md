# Golang Echo API

## Config

Create the `.env` file.

```bash
aws s3 cp "s3://dph-developer-tools/dph/scripts/local-env-vars.sh" "$HOME/dph-local-env-vars.sh"
```

```bash
bash "$HOME/dph-local-env-vars.sh" "eu-west-1" "/dph/config/global;/dph/config/local" $(pwd)
```

## Use case: maintenace

### Running API

To start the API, first you need to install dependencies by executing following commands:

```bash
#  $ go mod download
#  $ go get -u github.com/swaggo/swag/cmd/swag
 $ go install github.com/swaggo/swag/cmd/swag
 $ $(go env GOPATH)/bin/swag init -g oph.api.comingsoon.app/app/app.go
```

Now you can either run the VSCode debugger to execute the project or execute the command:

```bash
$ export $(echo $(cat .env | sed 's/#.*//g'| xargs) | envsubst) && go run ./oph.api.comingsoon.app/main.go
```

In the logs, the API will output something like:

```bash
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [host]:[port]
```

This indicates the API has started.

You can open Swagger with this URL: `http://127.0.0.1:7890/swagger/`

To run the docker container, execute command:

```bash
$ export $(echo $(cat .env | sed 's/#.*//g'| xargs) | envsubst) &&  docker-compose -f 'docker-compose.yml' up --build
```

# if you get an auth error for `docker login`, enter the following command:

```bash
$ export $(echo $(cat .env | sed 's/#.*//g'| xargs) | envsubst) &&  echo $(aws ecr get-login-password | docker login --username AWS --password-stdin $IMAGE_REGISTRY_BASE_URL)
```

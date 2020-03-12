# Go-Kit Example
Microservices example using [go-kit](https://github.com/go-kit/kit).

### Prerequisite
Before start this project please install:
 - [Docker](https://www.docker.com)
 - [Docker compose](https://docs.docker.com/compose/install/)

### Get Started

```shell
$ go get github.com/ranggakusuma/go-kit-example
$ docker-compose up
```

### API

```shell
$ curl --request GET \
  --url 'http://localhost:8080/search?searchword=batman&pagination=2'
```

 

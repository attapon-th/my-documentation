# My Documentation


## Contents
- [My Documentation](#My-Documentation) 
     - [Contents](#contents)
     - [Dockerfile](#dockerfile)
     - [Docker-Compose](#docker-compose)
     - [GOLANG](#golang)

## SSH Tunnel
```shell
       local:server-remote-to      <- ssh ->
ssh -L 9000:8.8.8.8:8080 -N  mophid
```


## Dockerfile 
- Python
    - [Python - basic image](./dockerfile/Dockerfile-python-basic)

- Golang
    - [Golang - scratch imange](./dockerfile/Dockerfile-golang-scratch)

- Vue.js
    - [Vuejs Nuxt - nuxt generate image](./dockerfile/Dockerfile-vuejs-nuxt)

## docker-compose
- Traefik
    - [docker@service - for service other join `traefic Network`](./docker-compose/docker-compose-for-tarfick.yml)

## SSH Key
- [SSH remote from keyfile](./ssh/README.MD)

## Makefile
- [Golang - for go project](./makefile/golang-makefile)

## GOLANG
- [GO Summary](./golang/summary.md)
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE0Njc5Nzk0NjZdfQ==
-->

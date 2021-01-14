# Golang Summary Project

##  Makefile Example
```makefile
.PHONY: all

# Set binary output
BINARY=<Project name>


VERSION=1.0.0

ifndef BUILD
BUILD=`git rev-parse HEAD`
endif

CHECKSUM=`shasum ${BINARY}`
APP_DEV=false
# Set Version and Build HEAD for 
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

dev:
	APP_DEV=true \
	go run ${LDFLAGS} "cmd/${BINARY}/main.go"

build:
	go build ${LDFLAGS} -o ${BINARY} "cmd/${BINARY}/main.go"

docker-build:
	docker build --build-arg BUILDDOCKER=$(BUILD) -t registry.gitlab.com/indev-moph/covid-dashboard/api-summary:latest .

docker-addtags:
	docker tag registry.gitlab.com/indev-moph/covid-dashboard/api-summary:latest registry.gitlab.com/indev-moph/covid-dashboard/api-summary:${VERSION}

docker-push: docker-addtags
	docker push -a registry.gitlab.com/indev-moph/covid-dashboard/api-summary

build-in-docker:
	CGO_ENABLED=0 GOOS=linux go build \
	-a -installsuffix cgo \
	${LDFLAGS} \
	-o AppMain "cmd/${BINARY}/main.go"

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f "${BINARY}.sum" ] ; then rm "${BINARY}.sum" ; fi
```

## Dockerfile 

```docke


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbNDU5MTc0NzA1XX0=
-->
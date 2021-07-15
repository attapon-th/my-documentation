# Docker file for Golang

```dockerfile
############################
# - Docker file Example for Golang into scratch image - #
############################
# Argument require
ARG BUILDDOCKER # build version in git commint state

############################
# STEP 1 build executable binary
############################
FROM golang:1.16-alpine  as builder
# install package for build
RUN apk -U --no-cache add \
    build-base git gcc bash tzdata git make ca-certificates \
    && update-ca-certificates
    
# Set TimeZone (require)
ENV TZ=Asia/Bangkok

# 
WORKDIR /go/src/AppBuild

# Download Modules
ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download

# add project 
# ควรมีไฟล์ .dockerignore เพื่อไม่เอาไฟล์ที่ไม่จำเป็น
ADD . .


# Create directory for binary
RUN mkdir -p /app 



# build go file
RUN BUILD=${BUILDDOCKER} \
	BINARY=${BINARY} \
	make build-in-docker

# copy file to path: /app
RUN make move-in-docker

############################
# STEP 2 build a small image
############################
FROM scratch

WORKDIR /app # fix path

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /app /app

# Set TimeZone
ENV TZ=Asia/Bangkok

ENTRYPOINT /app/${BINARY}
```

## Makefile For Build Docker Image
```makefile
.PHONY: all
ifndef BINARY
BINARY=AppMain
endif
ifndef VERSION
VERSION=0.0.1
endif
ifndef BUILD
BUILD=`git rev-parse HEAD`
endif
# go main file
GOMAINFILE=main.go
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
GIT_REGISTRY_URL="registry.gitlab.com/"




docker-build:
	docker build \
	--build-arg BUILDDOCKER=$(BUILD) \
	-t ${GIT_REGISTRY_URL}:latest .

docker-addtags:
	docker tag ${GIT_REGISTRY_URL}:latest ${GIT_REGISTRY_URL}:${VERSION} 

docker-push:
	docker push ${GIT_REGISTRY_URL}:latest 

build-in-docker:
	CGO_ENABLED=0 GOOS=linux go build \
	-a -installsuffix cgo ${LDFLAGS} \
	-o ${BINARY} ${GOMAINFILE}
	
move-in-docker:
	mv ${BINARY} /app/${BINARY}  
```


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbNzEwOTM0NjQsNDc3MjQ4Njk5LDk0OTExND
Q4MSwtMTg5NDQ5NTk5NywxNzQ4NDk2NTEzXX0=
-->
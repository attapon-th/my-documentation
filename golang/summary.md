# Golang Summary Project

##  Makefile Example
```makefile
.PHONY: all

# Set binary output
BINARY=covid_summary_api


VERSION=1.1.0

ifndef BUILD
BUILD=`git rev-parse HEAD`
endif

CHECKSUM=`shasum ${BINARY}`
APP_DEV=false
# Set Version and Build HEAD for 
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

GIT_REPO_URL=registry.gitlab.com/indev-moph/covid-project/api-summary

dev:
	APP_DEV=true URL_PREFIX=/summary-api \
	go run ${LDFLAGS} cmd/${BINARY}/main.go -s localhost:8080 -c .env

build:
	go build ${LDFLAGS} -o ${BINARY} "cmd/${BINARY}/main.go"

docker-build:
	docker build --build-arg BUILDDOCKER=$(BUILD) -t ${GIT_REPO_URL}:latest .

docker-addtags:
	docker tag ${GIT_REPO_URL}:latest ${GIT_REPO_URL}:${VERSION}

docker-push: docker-addtags
	docker push ${GIT_REPO_URL}:latest
	docker push ${GIT_REPO_URL}:${VERSION}


build-in-docker:
	CGO_ENABLED=0 GOOS=linux go build \
	-a -installsuffix cgo \
	${LDFLAGS} \
	-o AppMain "cmd/${BINARY}/main.go"

rebuild: clean ${BINARY}
	echo $(shasum ${BINARY}) " ${VERSION} ${BUILD}" > "${BINARY}.sum"

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f "${BINARY}.sum" ] ; then rm "${BINARY}.sum" ; fi




send-static:
	scp ./public/* indev:/home/attapon/covid-summary-api/public
	scp ./docker-compose-indev.yml indev:/home/attapon/covid-summary-api/docker-compose.yml
docker-compose-up: send-static	
	ssh indev "cd /home/attapon/covid-summary-api; \
	docker-compose pull; \
	docker-compose  up -d"

```

## Dockerfile 

```dockerfile
############################
# STEP 1 build executable binary
############################
FROM golang:1.15-alpine  as builder


# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
# tzdata
RUN apk -U --no-cache add \
    build-base \
    git \
    gcc \
    bash \
    tzdata \
    git \
    make \
    ca-certificates \
    && update-ca-certificates


# Set TimeZone
ENV TZ=Asia/Bangkok
# Create appuser
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /go/src/AppBuild

# Download Modules
ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download

# Add project file.
ADD . .

# Create directory for binary
RUN mkdir -p /app 


# Build the binary
ARG BUILDDOCKER
RUN BUILD=${BUILDDOCKER} make build-in-docker
RUN mv AppMain /app/AppMain \
    && mv .env /app/.env
# RUN CGO_ENABLED=1  go build -o /app/covid-report-api ./cmd/api/api.go


############################
# STEP 2 build a small image
############################
FROM scratch

WORKDIR /app

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /app /app

# Use an unprivileged user.
USER appuser:appuser

# Set TimeZone
ENV TZ=Asia/Bangkok
EXPOSE 80

ENTRYPOINT ["/app/AppMain"]

# default run port 80
CMD ["-s", ":80"]
```


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTk4ODE3ODk4LDE4Mjg1NTM2NywtODA5Nz
Q3MDA4XX0=
-->
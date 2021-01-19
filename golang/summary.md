# Golang Summary Project

##  Makefile Example
```makefile

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
eyJoaXN0b3J5IjpbMTgyODU1MzY3LC04MDk3NDcwMDhdfQ==
-->
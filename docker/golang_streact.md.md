# Docker file for Golang

```dockerfile
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


# Build the binary
ARG BUILDDOCKER # build version in git commint state
ARG BINARY # output file binary name 

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
WORKDIR /app
# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /app /app

COPY docs /app/docs
# Use an unprivileged user.
# USER appuser:appuser


# Set TimeZone
ENV TZ=Asia/Bangkok
EXPOSE 80


ENTRYPOINT ["/app/AppMain"]
```

> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTQ2MDY0ODQxMl19
-->
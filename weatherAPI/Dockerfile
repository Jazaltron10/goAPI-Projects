
FROM --platform=${BUILDPLATFORM} golang:1.15-alpine3.12 AS base

# fetching the dependencies
RUN apk update && apk add --no-cache git

# Set the current working directory
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. 
# Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 
# Add the source from the current directory
ADD . .
# Access source folder to build main.go
WORKDIR /app/src
ENV CGO_ENABLED=0

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache/go-build \
GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main .

# add unittests
#FROM base AS unit-test
#RUN --mount=type=cache,target=/root/.cache/go-build \
#go test -v .

# Start a new stage from scratch
FROM alpine:latest AS executable
RUN apk --no-cache add ca-certificates

#WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
# This should also an .env file
COPY --from=build /app .
COPY --from=build /app/src .
COPY --from=build /app/init.sh /init.sh

RUN chmod +x /init.sh

ENTRYPOINT ["/init.sh"]
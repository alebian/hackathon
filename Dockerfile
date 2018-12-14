# Create the builder image
FROM golang:1.11.2-alpine as builder
RUN apk update && \
  apk add git

# Download the dependencies manually to use Docker cache
RUN go get github.com/gin-gonic/gin
RUN go get github.com/contribsys/faktory_worker_go

COPY . $GOPATH/src/hackathon/
WORKDIR $GOPATH/src/hackathon/
# In case we missed some dependencies
RUN go get -d -v
RUN go build -o /go/bin/app *.go

# Create a scratch image with just what we need
FROM alpine:3.8
WORKDIR /app
COPY --from=builder /go/bin/app /app/app
CMD ["./app"]
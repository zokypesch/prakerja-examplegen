FROM golang:1.13.3

# Set Arguments
ARG APP_NAME=platform

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

# build directories
ENV SRC_DIR="/go/src/app"
RUN mkdir /app
RUN mkdir $SRC_DIR

# Copy current directory
COPY ./Gopkg.lock ./Gopkg.toml $SRC_DIR/
WORKDIR $SRC_DIR

# Go dep!
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -vendor-only

# Build my app
# RUN go build -o /app/main .
COPY . $SRC_DIR/
RUN GOOS=linux go build -a -o /app/main/$APP_NAME .
CMD ["/app/main/$APP_NAME"]

# ARG APP_NAME=platform

# FROM golang:1.13-stretch as builder

# ARG APP_NAME

# ENV SRC_DIR="/go/src/github.com/zokypesch/prakerja-examplegen/$APP_NAME"

# # Gopkg.toml and Gopkg.lock lists source dependencies.
# # These layers will only be re-built when Gopkg files are updated.
# COPY ./app/Gopkg.lock ./app/Gopkg.toml $SRC_DIR/

# WORKDIR $SRC_DIR

# # Install library dependencies.
# RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -vendor-only

# # Copy all source and build it.
# # This layer will be rebuilt whenever a file has changed in the source directory.
# COPY ./app .
# RUN GOOS=linux go build -a -o /bin/$APP_NAME .
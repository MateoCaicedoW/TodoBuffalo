FROM golang:1.18-alpine as builder

# Installing nodejs and other dependecies
RUN apk add --update nodejs-current npm curl bash build-base

# Installing Yarn
RUN curl -o- -L https://yarnpkg.com/install.sh | bash -s -- --version 1.22.10
ENV PATH="$PATH:/root/.yarn/bin:/root/.config/yarn/global/node_modules"

WORKDIR /todo
ADD go.mod .
ADD go.sum .
RUN go mod download -x

# Installing ox
RUN go install github.com/wawandco/oxpecker/cmd/ox@master
ADD . .

# Building the application binary in bin/app
RUN go build -o ./bin/cli -ldflags '-linkmode external -extldflags "-static"' ./cmd/ox 
RUN ox build --static -o ./bin/app

FROM alpine

# Binaries
WORKDIR /bin/
COPY --from=builder /todo/bin/* /bin/
COPY --from=builder /todo/migrations /bin/migrations
ENV ADDR=0.0.0.0
EXPOSE 3000


CMD /bin/cli db migrate; /bin/cli db task create:users:admin; /bin/app
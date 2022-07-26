FROM golang:1.18-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org/

# Installing nodejs and other dependecies
RUN apk add --update nodejs-current curl bash build-base

# Installing Yarn
RUN curl -o- -L https://yarnpkg.com/install.sh | bash -s -- --version 1.22.10
ENV PATH="$PATH:/root/.yarn/bin:/root/.config/yarn/global/node_modules"

WORKDIR /lit_gorge_57839
ADD go.mod .
ADD go.sum .
RUN go mod download -x

# Installing ox
RUN go install github.com/wawandco/oxpecker/cmd/ox@master
ADD . .

# Building the application binary in bin/app
RUN ox build --static -o bin/app

RUN go build -o ./bin/cli -ldflags '-linkmode external -extldflags "-static"' ./cmd/ox 

FROM alpine



# Binaries
WORKDIR /bin/

COPY --from=builder /bin/* /bin/

ENV ADDR=0.0.0.0
EXPOSE 3000

# For migrations use
CMD /bin/cli db migrate; /bin/app
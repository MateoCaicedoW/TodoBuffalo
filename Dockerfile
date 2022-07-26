FROM golang:1.18-alpine as builder

# Installing nodejs
RUN apk add --update nodejs-current curl bash build-base

# Installing Yarn
RUN curl -o- -L https://yarnpkg.com/install.sh | bash
ENV PATH="$PATH:/root/.yarn/bin:/root/.config/yarn/global/node_modules"

# Installing Ox
RUN go install github.com/wawandco/ox/cmd/ox@latest
WORKDIR /lit_gorge_57839
ADD . .

# Building the application binary in bin/app 
RUN ox build --static -o bin/app

FROM alpine
# Binaries

WORKDIR /bin/

COPY --from=builder  /bin/* /bin/

# For migrations use 
# CMD ox db migrate; app 
CMD app
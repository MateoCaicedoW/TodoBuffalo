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

# Binaries
COPY --from=builder /lit_gorge_57839/bin/app /bin/app
COPY --from=builder /lit_gorge_57839/bin/cli /bin/cli

ENV ADDR=0.0.0.0
EXPOSE 3000

# For migrations use
CMD /bin/cli db migrate; /bin/app
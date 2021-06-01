# Build Step
FROM golang:1.16-alpine as BUILD_STEP
LABEL maintainer="Jintae, Kim <6199@outlook.kr>"

COPY . /app
ENV HOME=/app

# Build
WORKDIR ${HOME}
RUN apk --no-cache add tzdata \
&& go build -o bin/main main.go

# Deploy Step
FROM alpine:latest

# Build Argument Set
ARG BROKER

# Env Set
ENV GIN_MODE=release
ENV PORT=8010
ENV BROKER=${BROKER}

# Timezone Set
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul

RUN mkdir /app && apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=BUILD_STEP /app/bin/main .

EXPOSE $PORT
ENTRYPOINT ["./main"]

FROM golang:1.16-alpine
LABEL maintainer="Jintae, Kim <6199@outlook.kr>"

COPY . /app
ENV HOME=/app

# Build Argument Set
ARG BROKER

# Env Set
ENV GIN_MODE=release
ENV PORT=8000
ENV BROKER=${BROKER}

# Timezone Set
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul

# Build
WORKDIR ${HOME}
RUN apk --no-cache add tzdata && go build main.go

EXPOSE $PORT
ENTRYPOINT ["./main"]

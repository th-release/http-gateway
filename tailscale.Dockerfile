FROM golang:alpine
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go mod download
RUN go build -o http-gateway
EXPOSE ${PORT:-8080}

# 의존성 설치 및 정리
RUN apk update && \
    apk add --no-cache curl iproute2 tailscale && \
    rm -rf /var/cache/apk/*

# Tailscale 소켓 및 캐시 디렉토리 생성
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale /var/log/tailscale && \
    chmod -R 755 /var/run/tailscale /var/cache/tailscale /var/lib/tailscale /var/log/tailscale

# TUN 디바이스 생성
RUN mkdir -p /dev/net && \
    mknod /dev/net/tun c 10 200 && \
    chmod 0666 /dev/net/tun

# Java 애플리케이션 JAR 복사
COPY ./http-gateway /app/http-gateway

# 작업 디렉토리 설정
WORKDIR /app

CMD ["sh", "-c", "if [ -z \"$AUTH_KEY\" ]; then echo 'Error: AUTH_KEY missing, skipping Tailscale'; else tailscaled --state=/var/cache/tailscale/tailscaled.state --socket=/var/run/tailscale/tailscaled.sock --verbose=1 & sleep 10 && tailscale up --auth-key=\"$AUTH_KEY\" --accept-routes --accept-dns --hostname=\"$TAILSCALE_HOSTNAME\" & echo 'Tailscale configured'; fi && exec ./http-gateway --PORT=${PORT} --SECRET=${SECRET} --LOGGING=${LOGGING}"]
FROM golang:1.19 as builder
WORKDIR /go/src/rbot-helper
COPY . .
RUN make rbot-helper

FROM alpine:latest
WORKDIR /root
RUN apk update \
	&& apk add --no-cache tzdata docker-cli \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV DOCKER_HOST=unix:///var/run/docker.sock
COPY --from=builder /go/src/rbot-helper/build/bin/rbot-helper /root/rbot-helper
CMD ["/root/rbot-helper"]
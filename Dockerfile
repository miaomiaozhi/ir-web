# 多阶段构建 构建 app
FROM golang AS builder

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

WORKDIR /app
COPY . /app

RUN go mod tidy && \
    GOSUMDB=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags "-w -s" -o main ./cmd/main.go

FROM ubuntu AS runner
LABEL maintainer="mozezhao <mozezhao@moresec.cn>"

ENV LANG en_US.utf8

WORKDIR /workspace

COPY --from=builder /app/main /workspace/
COPY --from=builder /app/conf/config.json /workspace/
COPY --from=builder /app/pages/ /workspace/pages/
RUN chmod -R 755 /workspace

EXPOSE 8080
ENTRYPOINT ["./main", "--config=./config.json"]
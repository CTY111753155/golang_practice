# 构建阶段
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 先复制依赖文件，利用缓存
COPY go.mod go.sum ./
RUN go mod download

# 复制源码并编译
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o booking ./cmd/web

# 运行阶段：使用最小镜像
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Taipei

WORKDIR /app

# 从构建阶段复制二进制和静态资源
COPY --from=builder /app/booking .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/email-templates ./email-templates

# 可选：使用非 root 用户运行
RUN adduser -D -g '' appuser
USER appuser

EXPOSE 3000

# 用 CMD 預設執行；compose 用完整 command 覆蓋，避免參數傳遞問題
CMD ["./booking"]

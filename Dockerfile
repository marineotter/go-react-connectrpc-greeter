# マルチステージビルド: Frontend Build stage
FROM node:24.2.0-alpine AS frontend-builder

# フロントエンドのビルド
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run build

# Backend Build stage
FROM golang:1.24.4-alpine AS backend-builder

# バックエンドのビルド
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend-app .

# Production stage
FROM nginx:alpine

# nginxの設定をコピー
COPY docker/nginx.conf /etc/nginx/nginx.conf

# ビルドされたバックエンドバイナリをコピー
COPY --from=backend-builder /app/backend/backend-app /app/backend/backend-app
RUN chmod +x /app/backend/backend-app

# ビルドされたフロントエンドの静的ファイルをコピー
COPY --from=frontend-builder /app/frontend/build/client /app/frontend/

# 起動スクリプトをコピー
COPY docker/start.sh /app/start.sh
RUN chmod +x /app/start.sh

# 作業ディレクトリを設定
WORKDIR /app

# nginxポート8080を公開
EXPOSE 8080

# 起動スクリプトを実行
CMD ["/app/start.sh"]

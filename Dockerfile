FROM golang:1.27.0 as stage1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gestorvideo ./cmd


############
FROM scratch
COPY --from=stage1 /app/gestorvideo /

ENTRYPOINT [ "/gestorvideo" ]
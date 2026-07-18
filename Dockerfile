FROM docker.io/library/golang:1.26-alpine AS builder
RUN apk add --no-cache make
WORKDIR /app
COPY go.mod Makefile ./
COPY . . 
RUN make build


FROM docker.io/library/golang:1.26-alpine AS runner 
WORKDIR /app
ENV GO_ENV=production

COPY --from=builder /app/bin ./bin

EXPOSE 1729

CMD ["./bin/goxe"]

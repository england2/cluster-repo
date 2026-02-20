FROM golang:1.25.4-bookworm AS build

WORKDIR /src/service2

COPY service2/go.mod ./
RUN go mod download

COPY service2/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/service2 .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /out/service2 /app/service2

ENTRYPOINT ["/app/service2"]

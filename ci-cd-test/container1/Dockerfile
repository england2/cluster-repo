FROM golang:1.25.4-bookworm AS build

WORKDIR /src/service1

COPY service1/go.mod ./
RUN go mod download

COPY service1/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/service1 .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /out/service1 /app/service1

EXPOSE 4872

ENTRYPOINT ["/app/service1"]

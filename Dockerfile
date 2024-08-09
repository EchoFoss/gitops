FROM golang:1.22.6-alpine AS build

WORKDIR /app
COPY . .

RUN go build -o server

FROM scratch

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 8080

ENTRYPOINT ["./server"]
FROM golang:alpine AS build

# GOPROXY resolves dependencies treefrom cache or repository
ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/api
COPY . .
# Set OS as linux
RUN GOOS=linux go build -o /go/bin/api cmd/main.go

EXPOSE 8080

FROM alpine
COPY --from=build /go/bin/api /go/bin/api
COPY --from=build /go/src/api/.env /go/bin/.env
ENTRYPOINT ["go/bin/api"]   
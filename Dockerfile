# Builder image
FROM golang:1.17-alpine
RUN apk --update add upx
WORKDIR /
COPY . .
RUN go mod download
RUN go mod verify
# Strip debug info + compress binary with upx
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o statuscode && upx statuscode

# Small final image
FROM alpine:latest
LABEL org.opencontainers.image.source https://github.com/jakewmeyer/status-codes
LABEL maintainer="jakewmeyer@gmail.com"
ENV GO_ENV=production
EXPOSE 7000
USER guest
ENTRYPOINT ["/statuscode"]
WORKDIR /
COPY --from=0 /statuscode .

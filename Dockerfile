FROM golang:1.22-alpine as builder
RUN apk add --no-cache git make tzdata upx
COPY . /src
WORKDIR /src
RUN make build-static
RUN upx --best --lzma dirsizer

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /src/dirsizer /usr/bin/dirsizer
ENTRYPOINT ["/usr/bin/dirsizer"]

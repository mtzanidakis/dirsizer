FROM golang:1.18-alpine as builder
RUN apk add --no-cache git upx
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 go build  -ldflags='-w -s -extldflags "-static"' -a
RUN upx dirsizer
RUN egrep '^root:' /etc/passwd > /etc/passwd.scratch && \
	egrep '^root:' /etc/group > /etc/group.scratch

FROM scratch as base
COPY --from=builder /etc/passwd.scratch /etc/passwd
COPY --from=builder /etc/group.scratch /etc/group
COPY --from=builder /src/dirsizer /bin/

CMD ["/bin/dirsizer"]

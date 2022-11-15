FROM jeyrce/golang:1.19.3-alpine as builder
ARG module
ARG goProxy
ENV GOPATH=/go
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GOPROXY=${goProxy}
WORKDIR /go/src/${module}
COPY . .
#RUN apk add make git
RUN make binary

FROM jeyrce/alpine:3.16.2-cn as runner
ARG commitId
ARG module
ARG app
LABEL poweredBy=${module} \
      commitId=${commitId}
WORKDIR /usr/local/bin
COPY config.yml /etc/${app}/config.yml
COPY media /var/lib/${app}/media/
COPY --from=builder /go/src/${module}/_out/* .
EXPOSE 80
VOLUME ["/etc/${app}/", "/var/lib/${app}/"]
VOLUME ["/etc/timezone:/etc/timezone", "/etc/locatime:/etc/localtime"]
CMD ["/usr/local/bin/${app}", "-c", "/etc/${app}/config.yml"]

#FROM alpine
#FROM golang
FROM ubuntu

COPY bin/go2/httpd /app/httpd
ENV TZ=Asia/Tokyo
RUN apt-get update && apt-get install -y tzdata

ENTRYPOINT ["/app/httpd"]
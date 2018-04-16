FROM alpine:3.7
MAINTAINER Neven Miculinic <neven.miculinic@gmail.com>
RUN apk update && apk add curl tcpdump bash bind-tools
ADD ./hello_world /usr/bin/hello_world
ENTRYPOINT ["hello_world"]

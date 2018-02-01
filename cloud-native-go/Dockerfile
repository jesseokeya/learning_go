FROM golang:1.7.3-alpine
MAINTAINER Jesse .C. Okeya

ENV SOURCES /go/src/github.com/jesseokeya/learning_go/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT cloud-native-go

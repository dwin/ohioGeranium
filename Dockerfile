# STEP 1 build executable binary
FROM golang:alpine as builder

LABEL maintainer "Darwin Smith II <dwin@dlsmi.com>"
LABEL app_version="0.1.0" architecture="amd64"

COPY /app $GOPATH/src/github.com/dwin/ohioGeranium/app
WORKDIR $GOPATH/src/github.com/dwin/ohioGeranium/app

#get dependancies
RUN apk add --update ca-certificates
#you can also use dep
RUN go get -d -v

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/ohioGeranium

# STEP 2 build a small image
# start from scratch
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/ohioGeranium /go/bin/ohioGeranium
EXPOSE 1313
ENTRYPOINT ["/go/bin/ohioGeranium"]

# docker build . -t dwin/ohioGeranium
# docker push dwin/go-ohioGeranium

# docker run -d -p 1313:1313 --name testApp dwin/ohioGeranium # use docker-compose up

# docker run -d --name testApp dwin/ohioGeranium
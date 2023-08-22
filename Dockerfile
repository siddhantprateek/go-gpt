# stage 1
FROM golang:alpine as builder 
RUN apk add --no-cache git 
WORKDIR /go/src/go-gpt  
COPY . . 
RUN go get -d -v ./.. 
RUN go build -o /go/bin/go-gpt -v ./.. 


# stage 2
FROM alpine:latest 
RUN apk --no-cache add ca-certificates 
COPY --from=builder /go/bin/go-gpt /go-gpt 
ENTRYPOINT /go-gpt 
LABEL Name="gogpt" Version=0.0.1 
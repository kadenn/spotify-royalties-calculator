FROM golang:latest 

ENV GIN_MODE=release
ENV PORT=8080

RUN mkdir /server 
ADD . /server/ 
WORKDIR /server 
RUN go mod download
RUN go build -o main .

EXPOSE $PORT
CMD ["./main"]
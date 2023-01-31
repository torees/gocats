FROM golang:1.19-alpine

WORKDIR /app

COPY go* ./
COPY main.go ./
COPY catlist.txt ./

RUN go mod download
RUN go build -o /gocats
ENV APP_PORT="8000" 
ENV	APP_MSG="Howdy wrangler" 
ENV	LOG_LEVEL="MUTE"

EXPOSE 8000

ENTRYPOINT ["/gocats"]
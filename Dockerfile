FROM golang:1.18.2

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go mod download \ 
&& go get github.com/dimeko/assets-proxy/api \
&& go get github.com/dimeko/assets-proxy/db 

CMD [ "go", "run", ".", "server" ]
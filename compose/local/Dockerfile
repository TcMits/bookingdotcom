FROM golang:1.20-bullseye

COPY ./compose/local/start /start
RUN sed -i 's/\r$//g' /start
RUN chmod +x /start

WORKDIR /bookingdotcom

COPY go.mod go.sum /bookingdotcom/
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

CMD ["/start"]

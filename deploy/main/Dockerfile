# Use the official Golang image as the base image
FROM golang:1.21.4
LABEL authors="rfanazhari"

ENV APP_USER apps
ENV APP_SERVICE api
ENV APP_HOME /app

WORKDIR $APP_HOME
USER $USER

COPY --chown=$APP_USER . .

ENV GO111MODULE on

RUN go mod tidy
RUN go mod vendor

RUN go build $APP_HOME/cmd/main.go

EXPOSE 80 443

CMD ["./main"]

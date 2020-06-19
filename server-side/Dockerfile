FROM golang:1.14-alpine
WORKDIR /app

RUN apk add --no-cache alpine-sdk git && go get -u github.com/oxequa/realize

EXPOSE 8080
CMD ["realize", "start"]

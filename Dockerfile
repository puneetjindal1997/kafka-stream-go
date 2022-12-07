FROM golang:alpine

WORKDIR /disastermanagement

COPY . .

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash make

EXPOSE 8080 4040

# RUN make all

CMD [ "make", "all" ]

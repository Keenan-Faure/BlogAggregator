FROM debian:stable-slim

FROM golang:1.20
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN mkdir web-app && cd web-app
WORKDIR /web-app/

# Migrations
ADD /sql/ /web-app/sql/
ADD migrations.sh /web-app/sql/schema/
COPY .env /web-app/sql/schema/

RUN ["chmod", "+x", "/web-app/sql/schema/migrations.sh"]

COPY out /web-app/
ENTRYPOINT [ "/web-app/out" ]
CMD ["/web-app/sql/schema/migrations.sh"]

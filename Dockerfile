FROM alpine:latest

RUN mkdir -p /opt/goapp

WORKDIR /opt/goapp

COPY ./configs/db/migrate ./configs/db/migrate
COPY ./build/gosolid .
COPY ./.env .

RUN chmod +x ./gosolid

EXPOSE 8080

ENTRYPOINT [ "/opt/goapp/gosolid" ]
CMD [ "serve" ]
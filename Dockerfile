FROM alpine:latest

RUN mkdir -p /opt/goapp

WORKDIR /opt/goapp

COPY ./configs/db/migrate ./configs/db/migrate
COPY ./build/gosolid .

RUN chmod +x ./gosolid

EXPOSE 9000

ENTRYPOINT [ "/opt/goapp/gosolid" ]
CMD [ "serve" ]
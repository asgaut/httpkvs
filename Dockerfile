FROM alpine:latest
EXPOSE 8080
WORKDIR /app
COPY httpkvs /app
ENTRYPOINT ["/app/httpkvs"]

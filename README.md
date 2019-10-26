# httpkvs

HTTP key-value server using PUT and GET verbs

## Building image and running in a Docker container
Disabling CGO is required to get a go binary without pthreads and libc dependency (which are not in alpine image).

```shell
CGO_ENABLED=0 go build httpkvs.go
docker build -t httpkvs .
```

`docker run --rm -d -p 8080:8080 httpkvs`

## Testing

Test inserting/updating:
```
curl -X PUT -d "dette er en test" http://localhost:8080/key1
curl -X PUT -H "Content-Type: application/json" -d "{asd: 'asdasd'}" http://localhost:8080/key2
```

Test retrieval:
```
curl -X GET http://localhost:8080/key1
```

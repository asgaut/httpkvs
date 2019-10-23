# httpkvs

HTTP key-value server using PUT and GET verbs

Test inserting/updating:
```
curl -X PUT -d "dette er en test" http://localhost:8080/key1
curl -X PUT -H "Content-Type: application/json" -d "{asd: 'asdasd'}" http://localhost:8080/key2
```

Test retrieval:
```
curl -X GET http://localhost:8080/key1
```

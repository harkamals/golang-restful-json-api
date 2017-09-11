### Restful JSON API

##### v2: Perform Tests
```sh
# Run Tests
go test ./v2 -v
```
##### Running Examples

```sh
curl http://localhost:8080/orders
```

##### Start web server
```sh
# Run example 2
go run ./v2/*.go
```

##### v1: Start web server
```sh
# Run example 1 (in-memory processing)
go run ./v1/*.go
```
##### Running Examples

```sh
# List
curl http://localhost:8080/todos
 
# Filter
curl http://localhost:8080/todos/1
 
# Create New
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
```

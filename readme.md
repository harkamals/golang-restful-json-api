# ToDo app: Restful JSON API

## Start web server
```sh
# Run example 1
go run ./v1/*.go
```

## Running Examples

```sh
# List
curl http://localhost:8080/todos
```

```sh
# Filter
curl http://localhost:8080/todos/1
```

```sh
# Create New
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
```


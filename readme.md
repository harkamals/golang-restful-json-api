### Restful JSON API

##### Change Log

```sh
# Version 3
  + Middleware for Logging
  + Using GORM - in-progress
 
# Version 2
  + Orders model
  + Use postgresql as database backend
  + Testing
  ~ Refactoring
    Viper for configuration mgmt
 
# Version 1
  + Initial release
  + In-memory 'todo' app model
```

##### Perform Tests
```sh
# Run Tests
go test ./v3 -v
```

##### Start web server
```sh
# Run example 
go run ./v3/*.go
```
##### REST Endpoints

```sh
# Database backend
curl http://localhost:8080/orders
curl http://localhost:8080/order/1
curl -H "Content-Type: application/json" -d '{"name":"New Order"}'  http://localhost:8080/order
 
# In-memory
curl http://localhost:8080/todos
curl http://localhost:8080/todos/1
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}'  http://localhost:8080/todos
```

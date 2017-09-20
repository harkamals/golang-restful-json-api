### Restful JSON API

##### Change Log

```sh
# Version 3
  + TLS, http redirect
  + Middleware for Logging
  + Middleware for API Auth - WIP
  + Using GORM for database models
 
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
go test -v
```

##### Generate Self-signed certs
```sh
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

##### Start web server
```sh
# Run example 
go run main.go
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


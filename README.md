# go-crud-app
A CRUD REST API with mysql<br>
Uses:
* gorilla/mux
* gorm

## Implemented

* [✔] `GET /book` returns list of books as JSON
* [✔] `GET /book/{id}` returns details of a specific book as JSON
* [✔] `POST /book` accepts a new book to be added
* [✔] `POST /book` returns status 415 if content is not `application/json`

### Data types
A book object should look like this:

```json
{
    "id": 3,
    "title": "Karnali Blues",
    "author": "Buddhi Sagar",
    "publication": "XYZ"
}
```

### Persistence
MYSQL database is used. 
Configure mysql database credentials by editing .../pkg/config/app.go

```go
// Database credentials
var (
	username = "root"
	password = "password"
	db       = "crud_app"
)
```

### <u>Setup</u>
```shell
go get "github.com/gorilla/mux"

go get 	"github.com/jinzhu/gorm"

go get "github.com/jinzhu/gorm/dialects/mysql"

go run cmd/main/main.go
```

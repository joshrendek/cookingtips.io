PostgreSQL client for Golang [![Build Status](https://travis-ci.org/go-pg/pg.svg)](https://travis-ci.org/go-pg/pg)
===

Supports:

- Basic types: integers, floats, string, bool, time.Time, and pointers to these types.
- sql.NullBool, sql.NullString, sql.NullInt64 and sql.Float64.
- [sql.Scanner](http://golang.org/pkg/database/sql/#Scanner) and [sql/driver.Valuer](http://golang.org/pkg/database/sql/driver/#Valuer) interfaces.
- Arrays.
- Partially hstore.
- Transactions.
- Prepared statements.
- Notifications: `LISTEN`/`NOTIFY`.
- `COPY FROM` and `COPY TO`.
- Timeouts. Client sends `CancelRequest` message on timeout.
- Connection pool.
- Queries are retried when possible.
- PostgreSQL to Go struct mapping.

API docs: http://godoc.org/gopkg.in/pg.v2. Make sure to check examples: http://godoc.org/gopkg.in/pg.v2#pkg-examples.

Installation
------------

Install:

    go get gopkg.in/pg.v2

Changelog
---------

### 0.2

* Support for named placeholders:

```go
a := &Article{Id: 1, Name: "Hello world"}
_, err := db.ExecOne(`UPDATE articles SET name = ?name WHERE id = ?id`, a)
```

* CopyFrom/CopyTo support:

```go
r := strings.NewReader("hello\t5\nworld\t5\nfoo\t3\nbar\t3\n")
res, err := t.db.CopyFrom(r, "COPY test FROM STDIN")
```

* Simplify collections:

```go
type Articles []*Article

func (articles *Articles) New() interface{} {
    a := &Article{}
    *articles = append(*articles, a)
    return a
}
```

### 0.1

Initial release.

Example
-------
```go
package pg_test

import (
	"fmt"

	"gopkg.in/pg.v2"
)

type User struct {
	Name   string
	Emails []string
}

type Users []*User

func (users *Users) New() interface{} {
	u := &User{}
	*users = append(*users, u)
	return u
}

func CreateUser(db *pg.DB, user *User) error {
	_, err := db.ExecOne(`INSERT INTO users VALUES (?name, ?emails)`, user)
	return err
}

func GetUsers(db *pg.DB) ([]*User, error) {
	var users Users
	_, err := db.Query(&users, `SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func ExampleDB_Query() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	_, err := db.Exec(`CREATE TEMP TABLE users (name text, emails text[])`)
	if err != nil {
		panic(err)
	}

	err = CreateUser(db, &User{"admin", []string{"admin1@admin", "admin2@admin"}})
	if err != nil {
		panic(err)
	}

	err = CreateUser(db, &User{"root", []string{"root1@root", "root2@root"}})
	if err != nil {
		panic(err)
	}

	users, err := GetUsers(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(users[0], users[1])
	// Output: &{admin [admin1@admin admin2@admin]} &{root [root1@root root2@root]}
}
```

Clinet=> (Request) = middleware =(Response)=> Server

FirstClass function =>
//sum signature ตรงกับ Math

literal function => short-cut of call function => anonymous + parameter (variable = function)

go routine

higher order fuction => (middleware) function receive/return function
-next = home
-wrappred := func("world")

SQL AWS elephantSQL
|command| desc|
|---|---|
| go mod init https://github.com/GunPkate | create module file |
| go mod tidy | Downlaod libraries |

```
createTb := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name TEXT,age INT)`
db.Exec(createTb)
db.QueryRow("INSERT INTO users (name,age) VALUES($1,$2) RETURNING id", "GunP", 23)
```

db.QueryRow => return value
db.exec => run only

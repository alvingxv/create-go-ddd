# Create Go DDD Project

> This npm package allows you to quickly create a Go project using the Domain-Driven Design (DDD) architecture. It automates the process of cloning a starter template, modifying the Go module path, and installing the project dependencies.

# Prerequisites
Before using this package, ensure that you have the following prerequisites installed on your machine:

- [npx](https://www.npmjs.com/package/npx)
- [Git](https://git-scm.com)
- [Go](https://golang.org)

# Installation

To create a new Go DDD project, run the following command using npx:

```bash
npx create-go-ddd
```

You will be prompted to enter the project name in the format "username/projectname". Provide the desired name and press Enter.

# Running the App

1. After creating the project, navigate to the project directory:

```bash
cd username/projectname
```
2. Open the db.go file under database folder.
3. Locate the following code block and replace the variables with your database credentials:
```go
var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
	dialect  = "postgres"
)

```
4. Save the file and run `go run main.go` and you're good to go!🚀🚀

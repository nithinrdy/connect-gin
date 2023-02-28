# Connect-gin

The backend web server for Connect. Built using Gin, Golang-JWT (and lib/pq as the database driver to connect to Postgres (Supabase)).

## Installing

1. Clone the repo

2. Install all the dependencies using `go get .`.

3. Create a .env file in the root directory and add the following environment variables: `DB_PASSWORD`, `DB_HOST_URL`, `DB_PORT`. See the models folder for the database schemas when setting up your own database.

4. Run the server using `go run main.go`.

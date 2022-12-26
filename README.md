## Setup development environment

- go 1.19+
- PostgreSQL 14.3

## Before Running

- Create database `db_destination` and `db_source` first
- Execute files `destination_product.sql` and `source_product.sql` in the `database` folder to create the table
- Execute files `destination_product_seed.sql` and `source_product_seed.sql` to insert 500 almost identical dummy data

## How to run
- Run `go mod tidy` 
- The services are in the `services` folder
- In the `cmd` folder, change file `.env.example` in each services into `.env` and change the credential connection for your postgresql
- Run `go run main.go` in each folder `cmd` services
- Execute `{{host}}/destination-product/update-all` method `GET` to update all the destination product table
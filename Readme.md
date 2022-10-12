# SMS-Service
To Run the application locally:
Run Postgres.
`docker run --name sms-service -d -e POSTGRES_USER=user -e POSTGRES_DB=sms -e POSTGRES_PASSWORD=root123 -p 5424:5432 postgres:13.4`
Load Schema
Note : It is optional as it will be loaded using migrations on app start

`docker exec -i sms-service psql -U root sms-service < db/setup.sql`
Run Application
`go run main.go`
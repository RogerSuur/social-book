# Social Network

## Useful links

- Social network [audit questions](https://github.com/01-edu/public/tree/master/subjects/social-network)
- [Intra link](https://01.kood.tech/intra/johvi/div-01/social-network?event=28) for the project description
- [Gitea Link](https://01.kood.tech/git/Jollyroger/social-network) for the project repo
- [Database schema migration](https://engineering.qubecinema.com/2019/09/20/sqlite-database-schema-migration-using-golang.html) example
- [Database access organising](https://www.alexedwards.net/blog/organising-database-access) example
- [Designing a notification system](https://tannguyenit95.medium.com/designing-a-notification-system-1da83ca971bc) example
- [Service pattern (project structure)](https://www.alexedwards.net/blog/the-fat-service-pattern)

## Running frontend server

- Option for development: `cd frontend` -> `npm start`
- Option for production: `npm run build` -> `node server.js` from frontend repo

## Running backend server

- `go run ./api/.`

If you wish to seed the database, run the command:
- `go run ./api/. seed`

## Database

- Creating new database migration: <br/>

        migrate create -ext sql -dir api/pkg/db/migrations/sqlite -seq schema_name

  "schema_name" is the name of the migration

- Updating the migration.go file code:

        go-bindata -o api/pkg/db/sqlite/migration.go   -prefix api/pkg/db/migrations/sqlite/  -pkg database api/pkg/db/migrations/sqlite

# social-network

# Database
* Creating new database migration: <br/>

        migrate create -ext sql -dir api/pkg/db/migrations/sqlite -seq schema_name

    "schema_name" is the name of the migration
* Updating the migration.go file code:

        go-bindata -o api/pkg/db/sqlite/migration.go   -prefix api/pkg/db migrations/sqlite/  -pkg database api/pkg/db/migrations/sqlite

# Useful links

* Social network [audit questions](https://github.com/01-edu/public/tree/master/subjects/social-network)
* [Intra link](https://01.kood.tech/intra/johvi/div-01/social-network?event=28) for the project description
* [Gitea Link](https://01.kood.tech/git/Jollyroger/social-network) for the project repo
* [Database access organising](https://www.alexedwards.net/blog/organising-database-access) example
* [Designing a notification system](https://tannguyenit95.medium.com/designing-a-notification-system-1da83ca971bc) example

module github.com/kirimi/rb2backend/user_service

go 1.23.1

require (
	github.com/kirimi/rb2backend/libs/db v0.0.0
	github.com/labstack/gommon v0.4.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.20.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/kirimi/rb2backend/libs/db v0.0.0 => ../../libs/db

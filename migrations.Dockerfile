FROM migrate/migrate

COPY ./repository/postgres/migrations migrations

ENTRYPOINT migrate -path migrations -database postgres://postgres:123@db:5432/test?sslmode=disable up
# GymControl-
Gestor De Alunos de Academia/Treinos e Dieta

## Para iniciar a aplicação suba o container
migrate -path=infra/database/migrations -database "postgres://pguser:pgpass@localhost:5433/usermanagementdb?sslmode=disable" -verbose up
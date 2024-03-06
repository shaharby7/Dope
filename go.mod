module github.com/shaharby7/Dope

go 1.21.5

require Dope/deployable v0.0.0-00010101000000-000000000000

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)

replace Dope/deployable => ./pkg/deployable

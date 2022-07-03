module github.com/alpheres/medialab/services/digitaltv-recorder

go 1.18

require (
	github.com/alpheres/medialab/pkg/utils v0.0.0
	github.com/rs/zerolog v1.27.0
)

replace github.com/alpheres/medialab/pkg/utils v0.0.0 => ../../pkg/utils

require (
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
)

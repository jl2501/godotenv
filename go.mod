module github.com/jl2501/godotenv/v3

go 1.23.0

toolchain go1.23.2

require github.com/spf13/afero v1.14.0

require golang.org/x/text v0.23.0 // indirect

retract (
	v3.0.2 // only contains retractions
	v3.0.1 // works, but should be under v1 as it is backwards compatible
	v3.0.0 // broken due to invalid module path in go.mod
)

module github.com/jl2501/godotenv

go 1.23.0

toolchain go1.23.2

require github.com/spf13/afero v1.14.0

require golang.org/x/text v0.23.0 // indirect

retract (
	v1.0.1 // has breaking API changes re-released as v2.0.1
	v1.0.0 // broken/ missing pieces -  was more of an internal PoC than a release
)

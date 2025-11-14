// Package godotenv is a go port of the ruby dotenv library (https://github.com/bkeepers/dotenv)
//
// Examples/readme can be found on the GitHub page at https://github.com/joho/godotenv
//
// The TL;DR is that you make a .env file that looks something like
//
//	SOME_ENV_VAR=somevalue
//
// and then in your go code you can call
//
//	Load()
//
// and all the env vars declared in .env will be available through os.Getenv("SOME_ENV_VAR")
package godotenv

import (
	"sync"
)

var getOsEnv = sync.OnceValue(func() *GoDotEnv {
	defaultEnv := NewGoDotEnv(NewOsFs())
	defaultEnv.Load()

	return defaultEnv
})

// Load will read your env file(s) and load them into ENV for this process.
//
// Call this function as close as possible to the start of your program (ideally in main).
//
// If you call Load without any args it will default to loading .env in the current path.
//
// You can otherwise tell it which files to load (there can be more than one) like:
//
//	Load("fileone", "filetwo")
//
// It's important to note that it WILL NOT OVERRIDE an env variable that already exists - consider the .env file to set dev vars or sensible defaults.
func Load(filenames ...string) (err error) {
	return getOsEnv().Load(filenames...)
}

// Overload will read your env file(s) and load them into ENV for this process.
//
// Call this function as close as possible to the start of your program (ideally in main).
//
// If you call Overload without any args it will default to loading .env in the current path.
//
// You can otherwise tell it which files to load (there can be more than one) like:
//
//	Overload("fileone", "filetwo")
//
// It's important to note this WILL OVERRIDE an env variable that already exists - consider the .env file to forcefully set all vars.
func Overload(filenames ...string) (err error) {
	return getOsEnv().Overload(filenames...)
}

// Read all env (with same file loading semantics as Load) but return values as
// a map rather than automatically writing values into env

func Read(filenames ...string) (map[string]string, error) {
	return getOsEnv().Read(filenames...)
}

// Exec loads env vars from the specified filenames (empty map falls back to default)
// then executes the cmd specified.
//
// Simply hooks up os.Stdin/err/out to the command and calls Run().
//
// If you want more fine grained control over your command it's recommended
// that you use `Load()`, `Overload()` or `Read()` and the `os/exec` package yourself.
func Exec(filenames []string, cmd string, cmdArgs []string, overload bool) error {
	return getOsEnv().Exec(filenames, cmd, cmdArgs, overload)
}

// Write serializes the given environment and writes it to a file.
func Write(envMap map[string]string, filename string) error {
	return getOsEnv().Write(envMap, filename)
}

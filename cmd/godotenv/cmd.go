package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/jl2501/godotenv"
)

func main() {
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "show help")
	var rawEnvFilenames string
	flag.StringVar(&rawEnvFilenames, "f", "", "comma separated paths to .env files")
	var overload bool
	flag.BoolVar(&overload, "o", false, "override existing .env variables")

	flag.Parse()

	usage := `
Run a process with an env setup from a .env file

godotenv [-o] [-f ENV_FILE_PATHS] COMMAND_ARGS

ENV_FILE_PATHS: comma separated paths to .env files
COMMAND_ARGS: command and args you want to run

example
  godotenv -f /path/to/something/.env,/another/path/.env fortune
`
	// if no args or -h flag
	// print usage and return
	args := flag.Args()
	if showHelp || len(args) == 0 {
		fmt.Println(usage)
		return
	}

	// load env
	var envFilenames []string
	if rawEnvFilenames != "" {
		envFilenames = strings.Split(rawEnvFilenames, ",")
	}

	// take rest of args and "exec" them
	cmd := args[0]
	cmdArgs := args[1:]

	err := godotenv.Exec(envFilenames, cmd, cmdArgs, overload)
	if err != nil {
		log.Fatal(err)
	}
}

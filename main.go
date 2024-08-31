package main

import "cmdprac/cmd"

var version string

func main() {
	cmd.Version = version
	cmd.Execute()
}

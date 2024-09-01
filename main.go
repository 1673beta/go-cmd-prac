package main

import "mkctl/cmd"

var version string

func main() {
	cmd.Version = version
	cmd.Execute()
}

package main

import "github.com/levibostian/dotenv/cmd"

var (
	version string
)

func main() {
	cmd.Execute(version)
}

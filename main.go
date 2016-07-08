package main

import "os"

const Version string = "v0.1.0" // Version バージョン

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

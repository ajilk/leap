package main

import (
	"leap/cmd"
	"leap/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}

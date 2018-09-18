
package main

import _ "github.com/benjaminbedard/grifts"
import "os"
import "log"
import "github.com/markbates/grift/grift"
import "path/filepath"

func main() {
	grift.CommandName = "buffalo task"
	if err := os.Chdir(filepath.Dir("C:/Users/Ruska/go/src/github.com/benjaminbedard/grifts")); err != nil {
	  log.Fatal(err)
	}
	err := grift.Exec(os.Args[1:], false)
	if err != nil {
		log.Fatal(err)
	}
}
package main

import "github.com/codegangsta/negroni"

func main() {
	n := negroni.Classic()

	n.UseHandler(NewRouter())
	n.Run(":3010")
}

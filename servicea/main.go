package main

import (
	"google.golang.org/appengine"

	_ "github.com/shihanng/gaegoasample/svc/servicea/cmd/servicea_svc"
)

func main() {
	appengine.Main()
}

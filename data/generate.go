// +build ignore

package main

import (
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	assets := http.Dir("assets")
	err := vfsgen.Generate(assets, vfsgen.Options{
		PackageName:  "data",
		VariableName: "assets",
	})
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"github.com/guotianqi/gotour/example"
	"github.com/guotianqi/gotour/exercise"
)

func main() {
	exampleRun()

	fmt.Println()

	exerciseRun()

	fmt.Println()
	httpServerRun()
}

func exampleRun() {
	fmt.Println("========= Example =========")

	example.FunctionsMain()
	example.VariablesMain()
	example.HelloMain()
	example.MutatingMapsMain()
	example.SelectMain()
}

func exerciseRun() {
	fmt.Println("========= Exercise =========")

	exercise.StringerMain()
	exercise.ErrorsMain()
	// exercise.ImagesMain()
	exercise.EquivalentBinaryTreesMain()
	exercise.WebCarwlerMain()
}

func httpServerRun() {
	fmt.Println("========= Http Server Run =========")
	// example.WebServersMain()
	exercise.HttpHandlersMain()
}

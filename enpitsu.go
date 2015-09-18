package main

//go:generate genqrc assets

import (
	"fmt"
	"github.com/mdouchement/enpitsu-go/app"
	"gopkg.in/qml.v1"
	"os"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	app.NewController(engine.Context())

	component, err := engine.LoadFile("qrc:///assets/enpitsu.qml")
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)
	win.Show()
	win.Wait()

	return nil
}

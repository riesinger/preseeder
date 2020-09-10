package main

import (
	"fmt"
	"os"

	"github.com/riesinger/preseeder/api"
	"github.com/riesinger/preseeder/preseed"
)

const (
	// Host determines to which host the HTTP server will bind
	Host = "0.0.0.0"
	// Port determines to which port the HTTP server will bind
	Port = 3000
)

var (
	// BaseDir is the directory in which the preseed files are contained
	BaseDir = "./preseeds"
)

func init() {
	if baseDir, ok := os.LookupEnv("BASE_DIR"); ok {
		BaseDir = baseDir
	}
}

func main() {
	preseedRenderer := &preseed.Renderer{BaseDirectory: BaseDir}
	if !preseedRenderer.DefaultPreseedExists() {
		fmt.Fprintf(os.Stderr, "Default preseed %s not found, exiting\n", preseedRenderer.DefaultPreseedPath())
		os.Exit(1)
	}
	server := api.NewHTTPServer(fmt.Sprintf("%s:%d", Host, Port), preseedRenderer)
	fmt.Printf("Listening for HTTP calls on %s:%d\n", Host, Port)
	server.Start()
}

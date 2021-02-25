package main

import (
	"os"

	"github.com/cosmicpunk/supermassive/app"
	"github.com/cosmicpunk/supermassive/cmd/supermassived/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"github.com/spf13/pflag"
	"math/rand"
	"mygo/cmd/server/app"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewServerCommand()
	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
	//pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

// WordSepNormalizeFunc changes all flags that contain "_" separators
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}

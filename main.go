package main

import (
	"github.com/kintuda/go-html-to-pdf/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	root := cmd.NewRootCmd()
	if err := root.Execute(); err != nil {
		log.Error().Err(err).Msg("command resulted in error")
	}
}

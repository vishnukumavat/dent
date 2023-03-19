package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagIsActive = "is-active"
)

func FlagSetQueryPolls() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.Bool(FlagIsActive, false, "returns only active polls if set true | default false")

	return fs
}

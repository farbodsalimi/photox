package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	photox "photox/pkg"
	"photox/pkg/config"
)

type runCmd struct {
	from string
	to   string
}

func (*runCmd) Name() string     { return "run" }
func (*runCmd) Synopsis() string { return "Print args to stdout." }
func (*runCmd) Usage() string {
	return `run [-from] <path to your photos>:
  Organize your photos by datetime.
`
}

func (r *runCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&r.from, "from", "", "path to your photos")
	f.StringVar(&r.to, "to", config.BasePath, "path to the place you want your organize photos to be")
}

func (r *runCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf("From: %s", r.from)
	fmt.Printf("To: %s", r.to)
	photox.Run(r.from, r.to)

	return subcommands.ExitSuccess
}

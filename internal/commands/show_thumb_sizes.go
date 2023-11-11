package commands

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/thumb"
	"github.com/photoprism/photoprism/pkg/report"
)

// ShowThumbSizesCommand configures the command name, flags, and action.
var ShowThumbSizesCommand = cli.Command{
	Name:   "thumb-sizes",
	Usage:  "Displays supported standard thumbnail sizes",
	Flags:  report.CliFlags,
	Action: showThumbSizesAction,
}

// showThumbSizesAction displays supported standard thumbnail sizes.
func showThumbSizesAction(ctx *cli.Context) error {
	rows, cols := thumb.Report(thumb.Sizes.All(), false)

	result, err := report.RenderFormat(rows, cols, report.CliFormat(ctx))

	fmt.Println(result)

	return err
}

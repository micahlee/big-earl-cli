package grow

import (
	"github.com/micahlee/big-earl-cli/command"
	"github.com/micahlee/big-earl-cli/util/api"
	"github.com/urfave/cli"
)

var openInBrowser, copyToClipboard bool

var flags = []cli.Flag{
	cli.BoolFlag{
		Name:        "clipboard,c",
		Usage:       "copies the expanded URL to your clipboard",
		Destination: &copyToClipboard,
	},
	cli.BoolFlag{
		Name:        "open,o",
		Usage:       "opens the expanded URL in your browser",
		Destination: &openInBrowser,
	},
}

func growUrl(inputUrl string, c *cli.Context) (string, error) {
	resp, err := api.ExpandUrl(inputUrl)
	if err != nil {
		return "", err
	}

	return resp.RedirectTo, nil
}

var Command = cli.Command{
	Name:    "grow",
	Aliases: []string{"g"},
	Usage:   "expands a short URL to its original URL",
	Flags:   flags,
	Action:  command.NewUrlTransformAction(growUrl, &openInBrowser, &copyToClipboard),
}

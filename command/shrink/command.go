package shrink

import (
	"github.com/micahlee/big-earl-cli/command"
	"github.com/micahlee/big-earl-cli/util/api"
	"github.com/urfave/cli"
)

var copyToClipboard, openInBrowser bool

var flags = []cli.Flag{
	cli.BoolFlag{
		Name:        "clipboard,c",
		Usage:       "copies the short URL to your clipboard",
		Destination: &copyToClipboard,
	},
	cli.BoolFlag{
		Name:        "open,o",
		Usage:       "opens the short URL in your browser",
		Destination: &openInBrowser,
	},
}

func shrinkUrl(inputUrl string, c *cli.Context) (string, error) {
	apiUrl := c.GlobalString("big_earl_api")

	resp, err := api.ShrinkUrl(apiUrl, inputUrl)
	if err != nil {
		return "", err
	}

	return resp.ShortUrl, nil
}

var Command = cli.Command{
	Name:    "shrink",
	Aliases: []string{"s"},
	Usage:   "gets a new short URL for the given input URL",
	Flags:   flags,
	Action:  command.NewUrlTransformAction(shrinkUrl, &openInBrowser, &copyToClipboard),
}

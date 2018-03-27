package preview

import (
	"github.com/micahlee/big-earl-cli/command"
	"github.com/micahlee/big-earl-cli/util/api"
	"github.com/urfave/cli"
)

var openInBrowser, copyToClipboard bool

var flags = []cli.Flag{
	cli.BoolFlag{
		Name:        "clipboard,c",
		Usage:       "copies the preview URL to your clipboard",
		Destination: &copyToClipboard,
	},
	cli.BoolFlag{
		Name:        "open,o",
		Usage:       "opens the preview URL in your browser",
		Destination: &openInBrowser,
	},
}

func previewUrl(inputUrl string, c *cli.Context) (string, error) {
	resp, err := api.ExpandUrl(inputUrl)
	if err != nil {
		return "", err
	}

	return resp.PreviewUrl, nil
}

var Command = cli.Command{
	Name:    "preview",
	Aliases: []string{"p"},
	Usage:   "provides the preview URL for a given short URL",
	Flags:   flags,
	Action:  command.NewUrlTransformAction(previewUrl, &openInBrowser, &copyToClipboard),
}

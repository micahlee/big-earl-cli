package command

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/micahlee/big-earl-cli/util/browser"
	"github.com/micahlee/big-earl-cli/util/clipboard"
	"github.com/urfave/cli"
)

type UrlConversion func(string, *cli.Context) (string, error)

// NewUrlTransformAction creates a new CLI command
// action which expects one argument containing a
// url in a string.
//
// This action will then transform that url with the
// provided converter and write it to std out, and
// optionally the clipboard and browser.
func NewUrlTransformAction(converter UrlConversion, openInBrowser *bool, copyToClipboard *bool) func(*cli.Context) error {
	return func(c *cli.Context) error {

		var inputUrl string

		// Ensure the URL input argument is present
		if c.NArg() < 1 {

			// Check to see if the input is being piped in
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				var err error
				reader := bufio.NewReader(os.Stdin)
				text, err := reader.ReadString('\n')
				if err != nil {
					return err
				}

				inputUrl = strings.TrimSpace(text)
			} else {
				// Otherwise raise an error that the URL is missing
				cli.ShowCommandHelp(c, "shrink")
				return cli.NewExitError("ERROR: Missing argument for short URL", 1)
			}
		} else {
			inputUrl = c.Args().First()
		}

		// Validate that the input is a URL
		_, err := url.ParseRequestURI(inputUrl)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("ERROR: Invalid URL: %v", inputUrl), 1)
		}

		// Delegate to the URL transformation
		outputUrl, err := converter(inputUrl, c)
		if err != nil {
			return err
		}

		// Print the transformed url to the stdout stream
		fmt.Println(outputUrl)

		// Open a browser to the output URL, if requested
		if *openInBrowser {
			fmt.Fprintf(os.Stderr, "Opening '%v' in browser...\n", outputUrl)
			browser.OpenUrl(outputUrl)
		}

		// Copy the output URL to the clipboard, if requested
		if *copyToClipboard {
			fmt.Fprintf(os.Stderr, "Copying '%v' to clipboard...\n", outputUrl)
			clipboard.Copy(outputUrl)
		}

		return nil
	}
}

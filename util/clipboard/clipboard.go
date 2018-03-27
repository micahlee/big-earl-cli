package clipboard

import (
	"github.com/atotto/clipboard"
)

func Copy(url string) error {
	return clipboard.WriteAll(url)
}

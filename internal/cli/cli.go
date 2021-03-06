package cli

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/atotto/clipboard"
)

// CopyToClipboard outputs markdown options and copies to the clipboard.
func CopyToClipboard(id string, options []string) error {
	var buf bytes.Buffer

	for _, o := range options {
		fmt.Fprintln(&buf, Link(id, o))
	}

	if err := clipboard.WriteAll(buf.String()); err == nil {
		return err
	}

	return nil
}

// Link returns a poll option link with image.
func Link(id, option string) string {
	return fmt.Sprintf(`[%s](https://api.gh-polls.com/poll/%s/%s/vote)`, Image(id, option), id, url.PathEscape(option))
}

// Image returns a poll option image.
func Image(id, option string) string {
	return fmt.Sprintf(`![](https://api.gh-polls.com/poll/%s/%s)`, id, url.PathEscape(option))
}

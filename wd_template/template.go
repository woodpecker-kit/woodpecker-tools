package wd_template

import (
	"fmt"
	"github.com/aymerick/raymond"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	RenderStatusShow = "success"
	RenderStatusHide = "failure"
)

// Render parses and executes a template, returning the results in string
// format. Trailing or leading spaces or new-lines are not getting truncated. It
// is able to read templates from remote paths, local files or directly from the
// string.
// please use wd_template.RegisterSettings(DefaultFunctions) once to use.
func Render(template string, payload interface{}) (s string, errOut error) {
	u, err := url.Parse(template)

	if err == nil {
		switch u.Scheme {
		case "http", "https":
			res, errGet := http.Get(template)

			if errGet != nil {
				errOut = fmt.Errorf("failed to fetch: %w", errGet)
				return s, errOut
			}

			defer func(Body io.ReadCloser) {
				errBodyRead := Body.Close()
				if errBodyRead != nil {
					errOut = fmt.Errorf("failed to close body: %w", errBodyRead)
				}
			}(res.Body)

			out, errRead := io.ReadAll(res.Body)

			if errRead != nil {
				errOut = fmt.Errorf("failed to read: %w", errRead)
				return s, errOut
			}

			template = string(out)
		case "file":
			out, errReadFile := os.ReadFile(u.Path)

			if errReadFile != nil {
				errOut = fmt.Errorf("failed to read: %w", errReadFile)
				return s, errOut
			}

			template = string(out)
		}
	}

	return raymond.Render(template, payload)
}

// RenderTrim parses and executes a template, returning the results in string
// format. The result is trimmed to remove left and right padding and newlines
// that may be added unintentially in the template markup.
// please use wd_template.RegisterSettings(DefaultFunctions) once to use.
func RenderTrim(template string, payload interface{}) (string, error) {
	out, err := Render(template, payload)
	return strings.Trim(out, " \n"), err
}

package multierror

import (
	"fmt"
	"strings"
)

// ErrorFormatFunc is a function callback that is called by Error to
// turn the list of errors into a string.
type ErrorFormatFunc func([]error) string

// ListFormatFunc is a basic formatter that outputs the number of errors
// that occurred along with a bullet point list of the errors.
func ListFormatFunc(es []error) string {
	if len(es) == 1 {
		return es[0].Error()
	}

	points := make([]string, len(es))
	for i, err := range es {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf(
		"%d errors occurred:\n\t%s",
		len(es), strings.Join(points, "\n\t"))
}

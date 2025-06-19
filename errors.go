// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.

package http

import (
	"fmt"

	api "github.com/gdt-dev/gdt/api"
)

var (
	// ErrAliasOrURL is returned when the test author failed to provide either
	// a URL and Method or specify one of the aliases like GET, POST, or DELETE
	ErrAliasOrURL = fmt.Errorf(
		"%w: either specify a URL and Method or specify one "+
			"of GET, POST, PUT, PATCH or DELETE",
		api.ErrParse,
	)
	// ErrExpectedLocationHeader indicates that the user specified the special
	// `$LOCATION` string in the `url` or `GET` fields of the HTTP test spec
	// but there have been no previous HTTP responses to find a Location HTTP
	// Header within.
	ErrExpectedLocationHeader = fmt.Errorf(
		"%w: expected Location HTTP Header in previous response",
		api.RuntimeError,
	)
)

// HTTPStatusNotEqual returns an ErrNotEqual when an expected thing doesn't equal an
// observed thing.
func HTTPStatusNotEqual(exp, got interface{}) error {
	return fmt.Errorf(
		"%w: expected HTTP status %v but got %v",
		api.ErrNotEqual, exp, got,
	)
}

// HTTPHeaderNotIn returns an ErrNotIn when an expected header doesn't appear
// in a response's headers.
func HTTPHeaderNotIn(element, container interface{}) error {
	return fmt.Errorf(
		"%w: expected HTTP headers %v to contain %v",
		api.ErrNotIn, container, element,
	)
}

// HTTPNotInBody returns an ErrNotIn when an expected thing doesn't appear in a
// a response's Body.
func HTTPNotInBody(element string) error {
	return fmt.Errorf(
		"%w: expected HTTP body to contain %v",
		api.ErrNotIn, element,
	)
}

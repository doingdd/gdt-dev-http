// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.

package http

import (
	"context"

	"github.com/gdt-dev/gdt/api"
	gdtcontext "github.com/gdt-dev/gdt/context"
	"gopkg.in/yaml.v3"
)

type httpDefaults struct {
	// BaseURL is used as the base of the URLs called by the gdt-http plugin's
	// Specs. If empty, fixtures are asked if they contain a "http.base_url"
	// state key and if so, that is used as the URL base.
	//
	// See the `httpServerFixture` for an example of how this works.
	BaseURL string `yaml:"base_url,omitempty"`
}

// Defaults is the known HTTP plugin defaults collection
type Defaults struct {
	httpDefaults
}

func (d *Defaults) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return api.ExpectedMapAt(node)
	}
	// maps/structs are stored in a top-level Node.Content field which is a
	// concatenated slice of Node pointers in pairs of key/values.
	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		if keyNode.Kind != yaml.ScalarNode {
			return api.ExpectedScalarAt(keyNode)
		}
		key := keyNode.Value
		valNode := node.Content[i+1]
		switch key {
		case "http":
			if valNode.Kind != yaml.MappingNode {
				return api.ExpectedMapAt(valNode)
			}
			hd := httpDefaults{}
			if err := valNode.Decode(&hd); err != nil {
				return err
			}
			d.httpDefaults = hd
		default:
			continue
		}
	}
	return nil
}

// BaseURLFromContext returns the base URL to use when constructing HTTP
// requests. If the Defaults is non-nil and has a BaseURL value, use that.
// Otherwise we look up a base URL from the context's fixtures.
func (d *Defaults) BaseURLFromContext(ctx context.Context) string {
	// If the httpFile has been manually configured and the configuration
	// contains a base URL, use that. Otherwise, check to see if there is a
	// fixture in the registry that has an "http.base_url" state key and use
	// that if found.
	if d != nil && d.BaseURL != "" {
		return d.BaseURL
	}
	// query the fixture registry to determine if any of them contain an
	// http.base_url state attribute.
	fixtures := gdtcontext.Fixtures(ctx)
	for _, f := range fixtures {
		if f.HasState(StateKeyBaseURL) {
			return f.State(StateKeyBaseURL).(string)
		}
	}
	return ""
}

// fromBaseDefaults returns an gdt-http plugin-specific Defaults from a Spec
func fromBaseDefaults(base *api.Defaults) *Defaults {
	if base == nil {
		return nil
	}
	d := base.For(pluginName)
	if d == nil {
		return nil
	}
	return d.(*Defaults)
}

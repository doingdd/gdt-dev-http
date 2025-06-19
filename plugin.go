// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.

package http

import (
	"gopkg.in/yaml.v3"

	"github.com/gdt-dev/gdt"
	api "github.com/gdt-dev/gdt/api"
)

func init() {
	gdt.RegisterPlugin(Plugin())
}

const (
	pluginName = "http"
)

type plugin struct{}

func (p *plugin) Info() api.PluginInfo {
	return api.PluginInfo{
		Name: pluginName,
	}
}

func (p *plugin) Defaults() yaml.Unmarshaler {
	return &Defaults{}
}

func (p *plugin) Specs() []api.Evaluable {
	return []api.Evaluable{&Spec{}}
}

// Plugin returns the HTTP gdt plugin
func Plugin() api.Plugin {
	return &plugin{}
}

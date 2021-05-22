package template

import (
	clouddns "github.com/aputs/libdns-clouddns"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *clouddns.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.clouddns",
		New: func() caddy.Module { return &Provider{new(clouddns.Provider)} },
	}
}

// TODO: This is just an example. Useful to allow env variable placeholders; update accordingly.
// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.JsonKeyFile = caddy.NewReplacer().ReplaceAll(p.Provider.JsonKeyFile, "")

	// Initialize the CloudDNS client session
	return p.NewSession(ctx)
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// clouddns {env.GOOGLE_APPLICATION_CREDENTIALS}
//
// clouddns {
//     json_key_file {env.GOOGLE_APPLICATION_CREDENTIALS}
//     project {env.PROJECT_ID}
// }
//
// **THIS IS JUST AN EXAMPLE AND NEEDS TO BE CUSTOMIZED.**
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.JsonKeyFile = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "json_key_file":
				if p.Provider.JsonKeyFile != "" {
					return d.Err("Json Key File already set")
				}
				if d.NextArg() {
					p.Provider.JsonKeyFile = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "project":
				if d.NextArg() {
					p.Provider.Project = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)

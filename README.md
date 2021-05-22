# clouddns module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [google clouddns](https://cloud.google.com/dns).

## Caddy module name

```
dns.provider.clouddns
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "clouddns"
        "json_key_file": "path to json key file",
        "project": "project name"
      }
    }
  }
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns clouddns {
    json_key_file {env.GOOGLE_APPLICATION_CREDENTIALS}
    project {env.PROJECT_ID}
  }
}
```

```
# one site
tls {
	dns clouddns {
    json_key_file {env.GOOGLE_APPLICATION_CREDENTIALS}
    project {env.PROJECT_ID}
  }
}
```

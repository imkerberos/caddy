package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"

	//dns providers
	_ "github.com/caddyserver/dnsproviders/auroradns"
	_ "github.com/caddyserver/dnsproviders/azure"
	_ "github.com/caddyserver/dnsproviders/cloudflare"
	_ "github.com/caddyserver/dnsproviders/cloudxns"
	_ "github.com/caddyserver/dnsproviders/conoha"
	_ "github.com/caddyserver/dnsproviders/digitalocean"
	_ "github.com/caddyserver/dnsproviders/dnsimple"
	_ "github.com/caddyserver/dnsproviders/dnsmadeeasy"
	_ "github.com/caddyserver/dnsproviders/dnspod"
	_ "github.com/caddyserver/dnsproviders/duckdns"
	_ "github.com/caddyserver/dnsproviders/dyn"
	_ "github.com/caddyserver/dnsproviders/exoscale"
	_ "github.com/caddyserver/dnsproviders/fastdns"
	_ "github.com/caddyserver/dnsproviders/gandi"
	_ "github.com/caddyserver/dnsproviders/gandiv5"
	_ "github.com/caddyserver/dnsproviders/generic"
	_ "github.com/caddyserver/dnsproviders/glesys"
	_ "github.com/caddyserver/dnsproviders/godaddy"
	_ "github.com/caddyserver/dnsproviders/googlecloud"
	_ "github.com/caddyserver/dnsproviders/lightsail"
	_ "github.com/caddyserver/dnsproviders/linode"
	_ "github.com/caddyserver/dnsproviders/namecheap"
	_ "github.com/caddyserver/dnsproviders/namedotcom"
	_ "github.com/caddyserver/dnsproviders/nifcloud"
	_ "github.com/caddyserver/dnsproviders/ns1"
	_ "github.com/caddyserver/dnsproviders/otc"
	_ "github.com/caddyserver/dnsproviders/ovh"
	_ "github.com/caddyserver/dnsproviders/pdns"
	_ "github.com/caddyserver/dnsproviders/rackspace"
	_ "github.com/caddyserver/dnsproviders/rfc2136"
	_ "github.com/caddyserver/dnsproviders/route53"
	_ "github.com/caddyserver/dnsproviders/selectel"
	_ "github.com/caddyserver/dnsproviders/stackpath"
	_ "github.com/caddyserver/dnsproviders/vscale"
	_ "github.com/caddyserver/dnsproviders/vultr"

	//plugins
	_ "blitznote.com/src/http.upload"
	_ "github.com/BTBurke/caddy-jwt"
	_ "github.com/SchumacherFM/mailout"
	_ "github.com/Xumeiquer/nobots"
	_ "github.com/aablinov/caddy-geoip"
	_ "github.com/abiosoft/caddy-git"
	_ "github.com/caddyserver/forwardproxy"
	_ "github.com/captncraig/caddy-realip"
	_ "github.com/captncraig/cors/caddy"
	_ "github.com/casbin/caddy-authz"
	_ "github.com/coopernurse/caddy-awslambda"
	_ "github.com/dhaavi/caddy-permission"
	_ "github.com/echocat/caddy-filter"
	_ "github.com/epicagency/caddy-expires"
	_ "github.com/freman/caddy-reauth"
	_ "github.com/hacdias/caddy-minify"
	_ "github.com/hacdias/caddy-service"
	_ "github.com/hacdias/caddy-webdav"
	_ "github.com/jung-kurt/caddy-cgi"
	_ "github.com/jung-kurt/caddy-pubsub"
	_ "github.com/linkonoid/caddy-dyndns"
	_ "github.com/lucaslorentz/caddy-docker-proxy/plugin"
	_ "github.com/mastercactapus/caddy-proxyprotocol"
	_ "github.com/miekg/caddy-prometheus"
	_ "github.com/miquella/caddy-awses"
	_ "github.com/nicolasazrak/caddy-cache"
	_ "github.com/payintech/caddy-datadog"
	_ "github.com/pieterlouw/caddy-grpc"
	_ "github.com/pyed/ipfilter"
	_ "github.com/restic/caddy"
	_ "github.com/simia-tech/caddy-locale"
	_ "github.com/tarent/loginsrv/caddy"
	_ "github.com/techknowlogick/caddy-s3browser"
	_ "github.com/xuqingfeng/caddy-rate-limit"
	_ "github.com/zikes/gopkg"
	_ "go.okkur.org/torproxy"

	//server types
	_ "github.com/lucaslorentz/caddy-supervisor/httpplugin"
	_ "github.com/lucaslorentz/caddy-supervisor/servertype"
	_ "github.com/pieterlouw/caddy-net/caddynet"

	//tls clustering
	_ "github.com/pteich/caddy-tlsconsul"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}

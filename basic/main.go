package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"

	//Plugins

	//DNS Providers
	_ "github.com/caddyserver/dnsproviders/acmedns"
	_ "github.com/caddyserver/dnsproviders/alidns"
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
	_ "github.com/caddyserver/dnsproviders/httpreq"
	_ "github.com/caddyserver/dnsproviders/inwx"
	_ "github.com/caddyserver/dnsproviders/lightsail"
	_ "github.com/caddyserver/dnsproviders/linode"
	_ "github.com/caddyserver/dnsproviders/linodev4"
	_ "github.com/caddyserver/dnsproviders/namecheap"
	_ "github.com/caddyserver/dnsproviders/namedotcom"
	_ "github.com/caddyserver/dnsproviders/namesilo"
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
	_ "github.com/caddyserver/dnsproviders/transip"
	_ "github.com/caddyserver/dnsproviders/vscale"
	_ "github.com/caddyserver/dnsproviders/vultr"

	//Directives/Middleware
	_ "github.com/BTBurke/caddy-jwt"                        //http.jwt
	_ "github.com/SchumacherFM/mailout"                     //http.mailout
	_ "github.com/Xumeiquer/nobots"                         //http.nobots
	_ "github.com/aablinov/caddy-geoip"                     //http.geoip
	_ "github.com/abiosoft/caddy-git"                       //http.git
	_ "github.com/caddyserver/forwardproxy"                 //http.forwardproxy
	_ "github.com/captncraig/caddy-realip"                  //http.realip
	_ "github.com/captncraig/cors/caddy"                    //http.cors
	_ "github.com/casbin/caddy-authz"                       //http.authz
	_ "github.com/echocat/caddy-filter"                     //http.filter
	_ "github.com/epicagency/caddy-expires"                 //http.expires
	_ "github.com/hacdias/caddy-minify"                     //http.minify
	_ "github.com/jung-kurt/caddy-cgi"                      //http.cgi
	_ "github.com/nicolasazrak/caddy-cache"                 //http.cache
	_ "github.com/pieterlouw/caddy-grpc"                    //http.grpc
	_ "github.com/pyed/ipfilter"                            //http.ipfilter
	_ "github.com/shuxs/caddy/gopkg"                        //http.gopkg
	_ "github.com/simia-tech/caddy-locale"                  //http.locale
	_ "github.com/xuqingfeng/caddy-rate-limit"              //http.ratelimit
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}

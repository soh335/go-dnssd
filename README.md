# go-dnssd

Golang interface to the dns-sd library.

* It is alpha quality.
* Requires [dns_sd.h](http://www.opensource.apple.com/source/mDNSResponder/mDNSResponder-320.5/mDNSShared/dns_sd.h).

Use OSX's `dns-sd` utility to debug:
```bash
# Browse http services on this network
dns-sd -B _http._tcp. .

# Get details aboout the "My Test" service
dns-sd -L "My Test" _http._tcp. .

# Register a "My Test" service
dns-sd -R "My Test" _http._tcp . 3000 path=/path-to-page.html
```


## Test DNS Server Based on geodns
We want to create a DNS server that can help create test & staging versions of an application. Historically this was done by modifying the hostname.  www.app.com would become www-staging.app.com. This clutters up the code and becomes cumbersome if the app is made up of several servers.  A better solution is to override the DNS for those hosts and point them to the test/staging servers.

We need a DNS server that:

1. **has an API** so that it orcestation tools like terraform can push the IP of the various servers to it.
1. **can be authoritative for the domains associated with the app but proxy everything else**.  This is different than normal as typically DNS servers are authoritative for an entire domain.  This is not what we want when we are testing a single server or application.  We only want to overide the hostnames for the servers we are testing.

We are going to base our work on the [geodns](https://github.com/miekg/dns) server.  It's open source, written in golang and has been shown to be performant in our internal testing.

### Adding an API
- the API will mimic the [Rage4 API](https://gbshouse.uservoice.com/knowledgebase/articles/109834-rage4-dns-developers-ap) as it's pretty simple, has a go package to access it and soon will have a terraform plugin.   
- geoDNS already has a `configWatcher` that can reload the zones when the config file changes.  We can use the code from that to generate the methods for the API.
- authentication for the API will be stored in the config file
- add the RESTful API inteface by using the built-in go http code (no framework needed).  This is because we don't need auth tokens or sessions as Rage4's API includes the credentials in each request

  - calls readZoneFile
    - NewZone
    - setupZoneData
      - label := Zone.AddLabel(dk)
      - label.Records[dnsType] = make(Records, len(records[rType]))
      - record := new(Record)
      - label.Records[dnsType][i] = *record
      - setupSOA(Zone)
  - addHandler( for that zone)
	-  dns.HandleFunc(name, setupServerFunc(config))
		- serve(w, r, Zone)

### Proxying Non-Authoritative Traffic
- will get upstream dns servers to use from the config file
- geoDNS uses the miekg/dns package.  This actually includes the DNS server and it also has the ability to proxy traffic.  So the work should mainly be configuring the proxy features of the miekg/dns package.
  - add a `dns.HandleFunc( *, func)` to catch all non-auth traffic
  - func will make a query to upstream and send reply back to caller
     - use [miekg/dnsrouter](https://github.com/miekg/dnsrouter/) as sample on how to receive a query a forward it

### Misc
- one issue with the geoDNS server is it uses the older, deprecated MaxMind [db file format](http://dev.maxmind.com/geoip/legacy/geolite/). We need to geo features to support for the LB for IPSec. We should upgrade the code to use v2 of the db



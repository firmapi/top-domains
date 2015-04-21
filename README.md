Top Domains
===========

A simple API to know if a domain is one of Alexa's top 1 millions websites.

## Examples

```
https://top-domains.herokuapp.com/rank?domain=google.com
=> {"status":"success","found":true,"domain":"google.com","position":1}
```

```
https://top-domains.herokuapp.com/rank?domain=wrong.site
=> {"status":"success","found":false,"domain":"wrong.site","position":0}
```

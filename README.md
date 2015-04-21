Top Domains
===========

A simple API to know if a domain is one of Alexa's top 1 millions websites.

## Examples

```
/rank?domain=google.com
=> {"status":"success","found":true,"domain":"google.com","position":1}
```

```
/rank?domain=wrong.site
=> {"status":"success","found":false,"domain":"wrong.site","position":0}
```

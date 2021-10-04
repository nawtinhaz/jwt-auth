# jwt-auth
Clean and simple authentication microservice written in Go.

## Internals

auth-service:
- signin
- signup
- blacklist

jwt package must provide:
- create
- validate

databases:
sqlite
redis
mongo
in-memory

We have to save this data somehow:
users -> username | password
tokens -> token | expiry_date | refresh_token???


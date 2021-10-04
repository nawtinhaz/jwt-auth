# jwt-auth
Clean and simple authentication microservice written in Go.

### To run project
```bash
make run
```

### To setup git hooks run

```bash
make install-hooks
```

### To run all checks manually

```bash
make check
```

### To skip pre-commit checks

```bash
make skipcheck-precommit
```

### To skip pre-push checks

```bash
make skipcheck-prepush
```

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


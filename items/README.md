# Test app

This is a tiny webservice used for testing.

## Build

```bash
docker build -t items:latest .
```

## Push new version

* Bump version in `VERSION.txt`
* Commit and push the change to `main`
* Check `push-items` workflow in GitHub Actions

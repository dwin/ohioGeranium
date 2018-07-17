# ohioGeranium
[![Go Report Card](https://goreportcard.com/badge/github.com/dwin/ohioGeranium)](https://goreportcard.com/report/github.com/dwin/ohioGeranium)


Application retrieves data from Postgres database if available or else retrieves from [NPPES API Endpoint](https://npiregistry.cms.hhs.gov/registry/help-api) and saves for future requests.

[Click here for API Documentation](https://documenter.getpostman.com/view/3362843/RWMENogB)

# Testing

*Step One*
Download using browser or ```curl``` to desired folder:
```bash
curl -O https://raw.githubusercontent.com/dwin/ohioGeranium/master/docker-compose.yml
```

*Step Two*
```bash
docker-compose up db & sleep 5 && docker-compose up
```

*Step Three*
```bash
curl --request GET \
  --url http://localhost:1313/status
```
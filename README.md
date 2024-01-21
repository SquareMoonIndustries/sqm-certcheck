![](https://github.com/SquareMoonIndustries/sqm-certcheck/actions/workflows/go.yml/badge.svg)
# sqm-certcheck

Simple service to check the expiration of a certificate for a domain.

Call with
```
curl -L "https://certcheck.sudde.eu/" \
-H "Content-Type: application/json" \
-d "{
    \"urls\":[
        {
            \"url\": \"https://github.com/\"
        },
        {
            \"url\": \"https://www.google.com/\"
        },
        {
            \"url\": \"https://www.google23.com/\"
        }
    ]
}"
```

Get responce
```
{
    "urls": [
        {
            "url": "https://github.com/",
            "expire": "2024-03-14T23:59:59Z",
            "error": ""
        },
        {
            "url": "https://www.google.com/",
            "expire": "2024-03-04T08:09:59Z",
            "error": ""
        },
        {
            "url": "https://www.google23.com/",
            "expire": "0001-01-01T00:00:00Z",
            "error": "Get \"https://www.google23.com/\": dial tcp: lookup www.google23.com on 127.0.0.53:53: no such host"
        }
    ]
}
```


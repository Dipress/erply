# erply

## Development

```sh
make dev
go run cmd/articles/main.go -db_url="postgres://articles:123456@localhost:5432/articles?sslmode=disable" -token=verysecrettoken
```

# Create Article

**URL** : `/articles`

**Method** : `POST`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content examples**

```json
{
    "id": 1234,
    "title": "Title",
    "body": "Bloggs",
    "created_at": "2019-02-26T19:33:03.219592Z",
    "updated_at": "2019-02-26T19:33:03.219592Z"
}
```

``` curl
curl -i -X POST http://localhost:8080/articles -H "Authorization: Bearer verysecrettoken" -d '{"title": "Title", "body": "Body"}'
```

# Find Article

**URL** : `/articles/{id}`

**Method** : `GET`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content examples**

```json
{
    "id": 1234,
    "title": "Title",
    "body": "Bloggs",
    "created_at": "2019-02-26T19:33:03.219592Z",
    "updated_at": "2019-02-26T19:33:03.219592Z"
}
```

``` curl
curl -i -X GET http://localhost:8080/articles/1 -H "Authorization: Bearer verysecrettoken"
```

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

**Example**:

``` curl
curl -i -X POST http://localhost:8080/articles -H "Authorization: Bearer verysecrettoken" -d '{"title": "Title", "body": "Body"}'
```

## Body

  * **title** - required, length from 1 to 50
  * **body** - required

## Responses

**Code** : `200 OK`

**Content example**

```json
{
    "id": 1234,
    "title": "Title",
    "body": "Bloggs",
    "created_at": "2019-02-26T19:33:03.219592Z",
    "updated_at": "2019-02-26T19:33:03.219592Z"
}
```

**Code** : `422 Unprocessable Entit`

**Content example**

```json
{
    "message": "you have validation errors",
    "errors": {
        "title": "cannot be blank",
        "body": "cannot be blank"
    }
}
```

**Code** : `400 Bad Request`

**Content example**

```json
{
    "message": "bad request"
}
```

**Code** : `501 Internal Server Error`

**Content example**

```json
{
    "message": "intenrnal server error",
}
```

# Find Article

**URL** : `/articles/{id}`

**Method** : `GET`

**Auth required** : YES

**Example**:

``` curl
curl -i -X GET http://localhost:8080/articles/1 -H "Authorization: Bearer verysecrettoken"
```

## Responses

**Code** : `200 OK`

**Content example**

```json
{
    "id": 1234,
    "title": "Title",
    "body": "Bloggs",
    "created_at": "2019-02-26T19:33:03.219592Z",
    "updated_at": "2019-02-26T19:33:03.219592Z"
}
```

**Code** : `404 Not Found`

**Content example**

```json
{
    "message": "not found",
}
```

**Code** : `501 Internal Server Error`

**Content example**

```json
{
    "message": "intenrnal server error",
}
```

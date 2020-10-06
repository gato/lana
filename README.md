# Lana REST API server

This Api implements the basic functionality requested on [https://github.com/lana/backend-challenge/blob/master/README.md]

## Implemented Endpoints:

## GET /api/v1/basket/

List available baskets

* input: *None*
* output: *an array of baskets ids*

```json
[
    "65e87c01-d21d-4225-a32e-b0921bb144d7",
    "c1f583b9-5157-4b5e-b5ed-977dd10ef010",
    "4794dc24-43b2-45c9-80c6-79005a8ced11",
    "c89e46f5-a616-4659-afbd-3a9cc32661ef"
]
```

## POST /api/v1/basket/

Create a new basket

* input: *None*
* output: *id of created basket*

```json
{
    "id": "c89e46f5-a616-4659-afbd-3a9cc32661ef"
}
```

## GET /api/v1/basket/:id

Get a basket by id

* input: *id (basket uuid) in url*
* output: *complete basket with items and total cost as json*

```json
{
    "amount": 12.5,
    "id": "c89e46f5-a616-4659-afbd-3a9cc32661ef",
    "items": [
        {
            "product": "MUG",
            "count": 1
        },
        {
            "product": "PEN",
            "count": 2
        }
    ]
}
```

## DELETE /api/v1/basket/:id

Delete a basket by id

* input: *id (basket uuid) in url*
* output: *None*

## POST /api/v1/basket/:id

Add an item to a basket.

* input: *id (basket uuid) in url*
* payload

```json
{
    "product" : "PEN",
    "count" : 2
}
```

* output: *current count of items of that type in basket*

```json
{
    "count" : 8
}
```

## Build process

Api was developed using current go (1.15) and go mods.
building locally is as simple as running

```bash
go build
```

or it can be built with docker running

```bash
docker build -t lana/0.1.0 .
```

This is a 2 steps build process that compiles in one image and then copies the result into a smaller image.

## How to run it

to run the api server in the command line just type

```bash
./lana
```

if default port is not usable you can pass a different port like this

```bash
./lana --port=12345
```

to run the docker image after building it just run

```bash
docker run -p 8080:8080 lana/0.1.0
```

## Test

to run the unittest asociated with the api just type

```bash
go test ./...
```

to also see coverage run

```bash
go test --cover ./...
```

this will output something like this

```text
?       github.com/gato/lana    [no test files]
ok      github.com/gato/lana/checkout   (cached)        coverage: 98.6% of statements
ok      github.com/gato/lana/merchandise        (cached)        coverage: 100.0% of statements
```

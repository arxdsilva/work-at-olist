# Olist

This project is a simple API that calculates bills from certain telephone records, it is also possible to insert records and search for calls.

## Installing

```shell
    $ go get github.com/arxdsilva/olist
```

## Testing

```shell
    $ cd ($GOPATH)/github.com/arxdsilva/olist
    $ go test -v ./...
```

## API

- [GoDoc](https://godoc.org/github.com/arxdsilva/olist)

### Routes

#### [GET]   "/"

This is a simple healthcheck route that always returns 200. 


#### [POST]  "/records"

It's the way of adding data to the API, by using the following templates:

- Call Start Record
```json
{
	"type": "start",
	"timestamp":"2018-07-28T21:57:13Z",
	"call_id": "22",
	"source": "2199999999",
	"destination": "2199999998"
}
```
- Call End Record
```json
{
	"type": "end",
	"timestamp":"2018-07-28T22:07:13Z",
	"call_id": "22"
}
```

Responses: 201, 400, 500

#### [GET]   "/bills/:subscriber"

The bills route needs a subscriber number as a param and optionaly as query params `month` and `year` (also numbers). Some examples of calls to bills are: `/bills/2199999999?month=1&year=2018` and `/bills/2199999999`.

This route responds always a JSON with an error or with the bill for the month/year requested (if not requested the last month will be used). A valid response will be in the bellow format.

```json
{
    "id": "219999999772018",
    "subscriber": "2199999997",
    "month": "7",
    "year": "2018",
    "calls": [
        {
            "bill_id": "219999999772018",
            "destination": "2199999998",
            "start_date": 28,
            "start_time": "9:57PM",
            "duration": "10m0s",
            "price": 0.54
        }
    ],
    "total": 0.54
}
```

Responses: 200, 400, 500

## Local development

In order to run this app locally, you'll need to run the following command with your postgres credentials:

```
    $ ➜  olist git:(master) ✗ DB_NAME=****** DB_HOST=localhost DB_USER=****** DB_PASS=****** go run main.go
```

## Development Environment

- Hardware: Asus g75vw
- OS: Ubuntu 18.04.1 LTS
- Text editor: VS Code
- Libraries:
    - [Echo web framework](https://github.com/labstack/echo)
    - [Check.v1](gopkg.in/check.v1)
    - [UUID](https://github.com/google/uuid)



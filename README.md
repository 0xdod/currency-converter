# A Currency Converter API

This is a currency converter API which converts only the following currencies from anyone to the other.
- Nigerian Naira (NGN)
- Ghanian Cedis (GHS)
- Kenyan Shilling (KHS)


#### Endpoint:
* GET /convert/{amount}/{from}/{to} converts the stated amount from one currency to another and upon successful conversion, returns data in the following format
```json
{
   "status": "success",
   "message": "Conversion successful",
   "data": {
       "from": {
           "currency": "ksh",
           "value": 10.00,
       },
       "to" : {
           "currency": "ngn"
           "value": 0.10,
       }
   }
}
```

#### Running locally
Clone the repo and make sure you have &gt;go1.16 installed, run `go build` and start the program.

#### Running tests
In the root directory, you can run tests with the following command 
```
go test
```


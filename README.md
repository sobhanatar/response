## Response Helper
A package to add more agility on response conventions change
## Description
This package takes data from developer and convert them into [convention format](https://conf.fidibo.com/x/kwL).

## Installation
To install the package enter below command into you terminal:  
`go get https://newgit.fidibo.com/fidiborearc/go/packages/response`
## Usage
After installing the package, add import statement to you import section of file:

`import resHelper "newgit.fidibo.com/fidiborearc/go/services/subscription/helper/response"`  

**Note:** There are two types of response this package can handle:
- success response (2xx responses)
- fail response (Any other responses except 2xx)

### How to use package for success responses
1- Create an instance of the response:  
`resHelper.NewSuccessResponse()`

2- Call any of provided method as chain of methods to achieve desired response:
- `Message(msg string)`: This method fills `message` field on response structure
- `Version(v int)`: This method fills `version` field on response structure
- `Data(d []map[string]interface{})`: This method fills `data->result` field on response structure
- `SingleData(d map[string]interface{})`: Just like `Data` method.The difference is this method only gets a single map 
instead of slice of maps.This is just a handy function for developers to be more comfortable as in most cases there is
  only a single set of data we want to pass out
- `Total(t int)`: This method fills `data->total` field on response structure (default zero value (0) will be used if
this method is not used)
- `PerPage(pp int)`: This method fills `data->per_page` field on response structure (default zero value (0) will be used if
  this method is not used)
- `Error(e string)`: This method fills `error` field on response structure. Most of the time you will not use this func
- `HttpCode(hc int)`: does not affect the response structure and only is used to increase the flexibility.We will be more
    flexible on changing convention based on http code
- `Generate()`: Creates and returns response with convention style

**Important Note:** `Data` and `SingleData` functions data are appended into `data->result` and not replaced.So you can
use them as many times as you want and all data are appended.
#### Example:
```go
package main

import (
	"fmt"
	resHelper "newgit.fidibo.com/fidiborearc/go/services/subscription/helper/response"
)

func main() {
  response := resHelper.NewSuccessResponse().
    Message("Successfully Done").
    Version("v2").
    Total(100).
    HttpCode(200).
    PerPage(25).
    SingleData(map[string]interface{}{"is_single_page": true}).
    SingleData(map[string]interface{}{"location": "Iran"}).
    Data([]map[string]interface{}{{"book": "a road to moon", "language": "english"}}).
    Generate()
  fmt.Println(response)
}

```

```json
{
  "data": {
    "per_page": 25,
    "result": [
      {
        "is_single_page": true
      },
      {
        "location": "Iran"
      },
      {
        "book": "a road to moon",
        "language": "english"
      }
    ],
    "total": 100
  },
  "error": "",
  "message": "Successfully Done",
  "represented_at": "2022-06-25 17:30:09",
  "version": "v2"
}
```

### How to use package for error responses
1- Create an instance of the error response:  
`resHelper.NewErrResponse()`

2- Call any of provided method as chain of methods to achieve desired response:
- `Message(msg string)`: This method fills `message` field on response structure
- `Error(e string)`: This method fills `error` field on response structure.
- `Version(v int)`: This method fills `version` field on response structure
- `Data(d []map[string]interface{})`: This method fills `data->result` field on response structure
- `SingleData(d map[string]interface{})`: Just like `Data` method.The difference is this method only gets a single map
  instead of slice of maps.This is just a handy function for developers to be more comfortable as in most cases there is
only a single set of data we want to pass out
- `ValidationErrors(errors map[string][]string)`: Receives a single map and converts it into the convention of validation
errors: `{"field": "FIELD_NAME", "error": "FIELD_ERROR"}`.Finally, the converted result will be appended into `data->result`
of response structure
- `HttpCode(hc int)`: does not affect the response structure and only is used to increase the flexibility.We will be more 
flexible on changing convention based on http code

- `Generate()`: Creates and returns response with convention style

**Important Note:** `Data`,`SingleData` and `ValidationErrors` functions data are appended into `data->result` and not replaced.So you can
use them as many times as you want and all data are appended.
#### Example:
```go
package main

import (
	"fmt"
	resHelper "newgit.fidibo.com/fidiborearc/go/services/subscription/helper/response"
)

func main() {
	response := resHelper.
		NewErrResponse().
		Message("Invalid data").
		Error("Data not found").
		SingleData(map[string]interface{}{"field_error": "invalid"}).
		SingleData(map[string]interface{}{"error": "just a single writer is accepted"}).
		Data([]map[string]interface{}{{"field": "custom_field", "error": "custom error message"}}).
		ValidationErrors(map[string][]string{"title": []string{"This field is required", "This field length is too high"}}).
		HttpCode(422).
		Generate()

    fmt.Println(response)
}

```

```json
{
  "data": {
    "per_page": 0,
    "result": [
      {
        "field_error": "invalid"
      },
      {
        "error": "just a single writer is accepted"
      },
      {
        "error": "custom error message",
        "field": "custom_field"
      },
      {
        "error": "This field is required",
        "field": "title"
      },
      {
        "error": "This field length is too high",
        "field": "title"
      }
    ],
    "total": 0
  },
  "error": "Data not found",
  "message": "Invalid data",
  "represented_at": "2022-06-25 17:56:29",
  "version": "v1"
}
```
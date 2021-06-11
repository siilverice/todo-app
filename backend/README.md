# Available endpoints

## Get all TODO list
`GET:    /v1/todo` 

Example Response
```
[
    {
        "ID": 1,
        "title": "Issue1",
        "Description": "Dont' use capital D here!"
    }
]
```

## Create TODO
`POST:   /v1/todo`

Example Request
```
{
    "title": "Hello world",
    "description": "this is description"
}
```

Example Response
```
{
    "ID": 12,
    "title": "Hello world",
    "Description": "this is description"
}
```

## Get specific TODO
`GET:    /v1/todo/:id`
 
 Example Response
 ```
 {
    "ID": 12,
    "title": "Hello world",
    "Description": "this is description"
}
 ```

 ## Update TODO
 `PUT:    /v1/todo/:id`

 Example Request
```
{
    "description": "this is new description"
}
```

Example Response
```
{
    "ID": 12,
    "title": "Hello world",
    "Description": "this is new description"
}
```

## Delete TODO
`DELETE: /v1/todo/:id`

No response body

<br>

# How to run it locally

`go run main.go`

Application will be available at http://localhost:8080

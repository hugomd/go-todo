# go-todo
> Simple to do list API with Gin and Gorm (with Postgres)

## Docker
Clone this repository and run:
```
docker-compose up
```

You can then hit the following endpoints:

| Method | Route      | Body                                         |
| ------ | ---------- | -------------------------------------------- |
| GET    | /tasks     |                                              |
| POST   | /tasks     | `{"title": "task title"}`                    |
| DELETE | /tasks/:id |                                              |
| PUT    | /tasks/:id | `{"title": "task title", "completed": true}` |

## Development
### Installing
I'm using [`dep`](https://github.com/golang/dep):
```bash
dep ensure
```
### Running locally
```
go run main.go
```

test

# Golang Clean Web API (Dockerize) with a full sample project (User profile creation with role)

## Used Tools

1. Gin as web framework
2. JWT for authentication and authorization
3. GORM as ORM

## How to run

### Run on local system

#### Start dependencies on docker

```bash
docker compose -f "docker/docker-compose.yml" up -d
```

#### Run migrations

```bash
cd src/scripts
./migrate.sh -up
```

#### Install swagger and run app

```bash
cd src
go install github.com/swaggo/swag/cmd/swag@latest
cd src/cmd
go run main.go
```

##### Address: [http://localhost:5005](http://localhost:5000)

#### Stop

```bash
docker compose -f "docker/docker-compose.yml" down
```

#### Examples

##### Login

```bash
curl -X 'POST' \
  'http://localhost:5000/api/v1/user/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "john.doe@gmail.com",
  "password": "password123"
}'
```

##### List all users

```bash
curl --location --request GET 'http://localhost:5000/apiv1/dashboard/listAllUsers' \
--header 'Content-Type: application/json' \
--header 'Authorization: <ENTER_YOUR_TOKEN>' \
--header 'Cookie: csrftoken=vGyWONqdzB8DyhUqN0MNsEE54cAVmq24' \
--data '{
    "page": 1,
    "pageSize": 10
}'
```

##### Sample Response

###### City filter and sort

```json
{
  "users": [
    {
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@gmail.com",
      "roleName": "admin"
    },
    {
      "firstName": "Jane",
      "lastName": "Smith",
      "email": "jane.smith@gmail.com",
      "roleName": "user"
    },
    {
      "firstName": "Alice",
      "lastName": "Johnson",
      "email": "alice.johnson@gmail.com",
      "roleName": "manager"
    }
  ],
  "page": 1,
  "pageSize": 10,
  "totalRecords": 3,
  "totalPages": 1
}
```

## Project preview

## Swagger

<p align="center"><img src='/docs/images/swagger.png' alt='Golang Web API preview' /></p>

# go-backend-docker-postgres
Go Backend API use of docker and 

## Run Go app with 
```bash
docker compose up --build -d
```

## POST Request via data get from JSON 
```bash
curl -d '{"email":"user@a.com", "name": "Name2", "password": "password"}' -H "Content-Type: application/json" -X POST http://127.0.0.1:8000/api/users
```

Output:
```json
{
  "email":"user@a.com",
  "password":"password",
  "name":"User2"
}
```


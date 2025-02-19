# go-backend-docker-postgres
Backend API use of docker and 

## Run Go app with 
```bash
docker compose up --build -d
```

## Run Postgres with 
```bash
curl -d '{"email":"karat@a.com", "id": 1, "name": "Name2", "password": "password"}' -H "Content-Type: application/json" -X POST http://127.0.0.1:8000/api/users
```
Output:
```json
{"email":"karat@a.com","password":"password","name":"Name2","id":1}
```
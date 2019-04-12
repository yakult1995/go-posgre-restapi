# go-posgre-restapi
RESTfulAPI with Golang and Postgres

## How to use
### .env
`cp docker/pgweb/.env.sample docker/pgweb/.env`  
`cp docker/go/.env.sample docker/go/.env`  
`cp docker/postgres/.env.sample docker/postgres/.env`  

### docker-compose build
`docker-compose build`

### docker-compose up
`docker-compose up`

## API Doc
```
GET    /            # Hello World!!
GET    /users       # user の一覧を表示
GET    /users/:id   # 指定した id の user を表示
POST   /users       # user を追加
PUT    /users/:id   # 指定した id の user を更新
DELETE /users/:id   # 指定した id の user を削除
```

## Example
### Create User
```
curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email": "hoge@example.com" }'
```

### Update User
```
curl -XPUT -H 'Content-Type:application/json' http://localhost:8080/users/1 -d '{"name": "koudaiii", "email": "hoge@example.com" }'
```

### Show User
```
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users/1
```

### List up Users
```
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
```

### Delete User
```
curl -XDELETE -H 'Content-Type:application/json' http://localhost:8080/users/1
```

## pgweb
```
localhost:8081
```

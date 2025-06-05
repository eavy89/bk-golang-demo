# BACKEND REST API DEMO
### For this demo I used the following technologies:
- Go
- Gin (https://github.com/gin-gonic/gin)
- SQLite with GORM
- JWT
- REST API

### To start the project you just need to execute the command:
```
    ./run_app.sh
```

### Enabled Endpoints:
#### Public:
1. [POST] /register
```aiignore
    {
        "username":"ivan2",
        "password":"1234"
    }
```

2. [POST] /login
```aiignore
    {
        "username":"ivan2",
        "password":"1234"
    }
```

#### Private:
1. [GET] /api/profile
```aiignore
    - Barear Token must needed
    - no body data
```

2. [POST] /api/purchase
```aiignore
    - Barear Token must needed
    {
      "item": "potatoes2",
      "quantity": 25,
      "price": 30
    }
```

3. [GET] /api/purchases
```aiignore
    - Barear Token must needed
    - no body data
```
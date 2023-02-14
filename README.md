# Evendo-viator

##Stack
```postgresql```
```Docker```
```Golang```

##Setup
1. Create an `.env` file
2. Add config as below
```azure
LOG_LEVEL=debug
ENABLE_DB=true
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=postgres
PORT=8001
CRONJOB_THREAD=5
NUMBER_OF_PRODUCT_BY_THREAD=50
VIATOR_ENDPOINT=https://api.sandbox.viator.com/partner/
```

##Build
Run `docker-compose up`


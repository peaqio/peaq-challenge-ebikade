# Coding Challenge Senior Backend Developer
## 3. Write an application for an analytics of cryptocurrency exchange rate. Your application must get from any market exchange rate for pairs:

# Stack
1. Golang
2. gRPC
3. MySQL
3. Protobuf

## Services

1. Exchange
2. Analytics

Build the services with `docker-composer build`
Run the services with `docker-composer up`

## Exchange Service 
* Handles the core business logic.
* Data Manipulation
* Communication with external APIs

## Analytics Service
* The Client interface between the Exchange service and the users.
* Handles API calls and exposes endpoint for users to retrieve data

## API Doc:
The API documentation can be found [HERE](127.0.0.1:50052/docs) after running the services above

## Curl Example
```
# CURL
'127.0.0.1:50052/export/analytics?from=2020-11-01T00%3A01%3A00&to=2020-12-01T23%3A59%3A00&format=json' \
```


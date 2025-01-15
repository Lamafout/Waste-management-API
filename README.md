# Waste management API

## Get started
#### Before running API on your machine you need to access next steps: 
1. **MongoDB configuration**
    1. Install MongoDB on your machine, create connection, start MongoDB.
    2. Create database with collections:
        - fkkos
        - okpds
        - technologies
        - producers
    3. Create [path to API]/config/.env file with variables:
        - MONGODB_URI - path to your mongodb connection
        - MONGODB_DATABASE - your database's name
2. **Go configuration**
    1. Install last version of Go
    2. Run "go mod tidy" on API's directory. 
    3. Add your host and port in config/server.json

#### Run API
- Run "go run app/main.go" on API's directory



## Paths:
- /technology - handles POST and GET requests. GET request includes id as param. Example: "/technology?id=1".
- /technologies - handles GET request. Returns array of technologies with filter. Example: "/tehnologies?filter=Пере&page=0" - returns first const value objects in collection with "Пере" in name field. This also can return filtered objects by fkko.name and fkko.code fields.
- /producer - handles POST request.
- /producers - handles GET request. Returns array of producers with filter. Example: "/producers?filter=Южная&page=0" - returns first const value objects in collection with "Южная" in name field. This also can return filtered objects by fkko.name and fkko.code fields. 
- /fkkos - handles GET request. Returns array of filtered fkkos. Example: "/fkkos?filter=Отход"
- /okpds - handles GET request. Returns array of filtered okpds. Example: "/okpds?filter=Продукция"
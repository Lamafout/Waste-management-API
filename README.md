# Waste management API

## Paths:
- /technology - handles POST and GET requests. GET request includes id as param. Example: "/technology?id=1".
- /technologies - handles GET request. Returns array of technologies with filter. Example: "/tehnologies?filter=Пере&page=0" - returns first const value objects in collection with "Пере" in name field
- /producer - handles POST request.
- /producers - handles GET request. Returns array of producers with filter. Example: "/producers?filter=Южная&page=0" - returns first const value objects in collection with "Южная" in name field 
- /fkkos - handles GET request. Returns array of filtered fkkos. Example: "/fkkos?filter=Отход"
- /okpds - handles GET request. Returns array of filtered okpds. Example: "/okpds?filter=Продукция"
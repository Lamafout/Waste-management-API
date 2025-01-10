# Waste management API

## Paths:
- /technology - handles POST and GET requests. GET request includes id as param. Example: "/technology?id=1".
- /technologies - handles GET request. Returns array of technologies without filter.
- /producer - handles POST request.
- /producers - handles GET request. Returns array of producers without filter.
- /fkkos - handles GET request. Returns array of filtered fkkos. Example: "/fkkos?filter=Отход"
- /okpds - handles GET request. Returns array of filtered okpds. Example: "/okpds?filter=Продукция"
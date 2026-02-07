# temp


# Useful commands
## MongoDB
```bash
# Enter Container
docker exec -it mongodb mongosh

# Enter DB
use admin
db.auth("DB_USERNAME", "DB_PASSWORD")

show dbs
use tarmac

show collections

db.collectionName.find().pretty()
db.collectionName.find({ key: value }).pretty()
```
mongoimport --db "$MONGO_DB_NAME" --collection icecream --host "$MONGO_HOST" --port "$MONGO_PORT" --drop --jsonArray --file ./data.json

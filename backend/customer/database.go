package customer

import "app/database"

var (
	CustomerCollection database.DataStore
)

func InitDB() {
	CustomerCollection = database.NewDataStore("gpserp", "customer")
}

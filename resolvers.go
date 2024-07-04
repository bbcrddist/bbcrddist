package gql

// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

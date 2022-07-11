package database

import "github.com/caiostarke/apiAndGRPC/logger_service/app/queries"

type Queries struct {
	*queries.LogQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := MongoDBConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		LogQueries: &queries.LogQueries{DB: db},
	}, nil
}

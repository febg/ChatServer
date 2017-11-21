package datastore

type Datastore interface {
}

type LocalDB struct {
}

func NewLocalDB() (*LocalDB, error) {
	return nil, nil
}

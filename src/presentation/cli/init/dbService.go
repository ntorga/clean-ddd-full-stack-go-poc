package cliInit

import "github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"

func PersistentDatabaseService() *db.PersistentDatabaseService {
	persistentDbSvc, err := db.NewPersistentDatabaseService()
	if err != nil {
		panic("PersistentDatabaseConnectionError:" + err.Error())
	}

	return persistentDbSvc
}

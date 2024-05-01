package dev

import "github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"

func GetPersistentDbSvc() *db.PersistentDatabaseService {
	persistentDbSvc, err := db.NewPersistentDatabaseService()
	if err != nil {
		panic("GetPersistentDbSvcError: " + err.Error())
	}
	return persistentDbSvc
}

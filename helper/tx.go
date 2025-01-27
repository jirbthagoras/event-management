package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		rollbackError := tx.Rollback()
		PanicIfError(rollbackError)
		panic(err)
	} else {
		commitError := tx.Commit()
		PanicIfError(commitError)
	}
}

package helper

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/logger"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if err != nil {
			logger.Error(context.Background(), "errorRollback", errorRollback)
		}
		panic(err)
	} else {
		errorCommit := tx.Commit()
		if err != nil {
			logger.Error(context.Background(), "errorRollback", errorCommit)
		}
	}
}

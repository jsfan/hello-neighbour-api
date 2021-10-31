package dal

import (
	"fmt"
	"github.com/pkg/errors"
)

func (dalInstance *DAL) BeginTransaction() error {
	if dalInstance.tx != nil {
		return errors.New("already in transaction")
	}
	var err error
	dalInstance.tx, err = dalInstance.Db.Begin()
	return err
}

func (dalInstance *DAL) CancelTransaction() error {
	if dalInstance.tx == nil {
		return errors.New("not in transaction")
	}
	err := dalInstance.tx.Rollback()
	dalInstance.tx = nil
	return err
}

func (dalInstance *DAL) CompleteTransaction() error {
	if dalInstance.tx == nil {
		return errors.New("not in transaction")
	}
	if err := dalInstance.tx.Commit(); err != nil {
		if errAbort := dalInstance.CancelTransaction(); errAbort != nil {
			return fmt.Errorf(
				"transaction could not be committed (%v) and subsequent rollback failed: %w",
				err,
				errAbort,
			)
		}
		return err
	}
	dalInstance.tx = nil
	return nil
}

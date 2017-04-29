package main

import (
	"log"

	"github.com/jboverfelt/infest/models"
	"github.com/markbates/pop"
)

func insertIntoDb(closures []models.Closure, db *pop.Connection) error {
	for _, model := range closures {
		err := db.Transaction(func(tx *pop.Connection) error {
			cur := model
			var found models.Closure
			ok, err := tx.Where("name = ?", cur.Name).Where("closureDate = ?", cur.ClosureDate).Exists(&found)

			if err != nil {
				return err
			}

			// if the id is not found then insert
			if !ok {
				err = tx.Save(&cur)
				if err != nil {
					log.Printf("Failed to save new closure %v\n", cur)
				}
			} else {
				// id is found, update the reopen date and comments
				found.ReopenDate = cur.ReopenDate
				found.ReopenComments = cur.ReopenComments
				err = tx.Save(&found)
				if err != nil {
					log.Printf("Failed to update closure %v\n", found)
				}
			}

			return nil
		})

		if err != nil {
			return err
		}
	}
	return nil
}

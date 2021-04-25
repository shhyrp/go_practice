package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func daoQuery() error {
	return errors.Wrap(sql.ErrNoRows, "dao query failed")
}

func find() error {
	return errors.WithMessage(daoQuery(), "find failed")
}

func main() {
	err := find()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}

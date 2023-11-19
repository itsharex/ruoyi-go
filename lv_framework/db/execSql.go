package db

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExecSqlFile(sqlFile string) error {
	path := filepath.Join(sqlFile)
	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		fmt.Println("XXXXXXXXXXXXXXXXXXX" + sqlFile)
	}
	sql := string(c)
	session := GetInstance().Engine().NewSession()
	session.Begin()
	_, err := session.Exec(sql)
	if err != nil {
		fmt.Println("XXXXXXXXXXXXX" + err.Error())
		session.Rollback()
	}
	session.Commit()
	return err
}
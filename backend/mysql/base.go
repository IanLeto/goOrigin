package mysql

import "goOrigin/errors"

var MySqlBackend *MySQLBackend

func init() {
	backend, err := NewMySQLBackend(nil)
	if err != nil {
		errors.CheckInitError(err)
	}
	MySqlBackend = backend
}

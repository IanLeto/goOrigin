package DB

import "goOrigin/pkg/errors"

var MySqlBackend *MySQLBackend

func init() {
	backend, err := NewMySQLBackend(nil)
	if err != nil {
		errors.CheckInitError(err)
	}
	MySqlBackend = backend
}

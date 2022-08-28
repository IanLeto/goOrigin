package DB

var MySqlBackend *MySQLBackend

func init() {
	backend, err := NewMySQLBackend(nil)
	if err != nil {
	}
	MySqlBackend = backend
}

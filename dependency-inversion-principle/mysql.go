package dependency

import "fmt"

// MySQLDB implements IDb
type MySQLDB struct {
	host string
	port float64
}

// Ping -> Returns MySQL Connection Information
func (m *MySQLDB) Ping() string {
	return fmt.Sprintf("host: %s, port: %f", m.host, m.port)
}

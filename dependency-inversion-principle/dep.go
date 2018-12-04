package dependency

// Run -> Driver for this Package
func Run() {
	db := &MySQLDB{}
	PasswordReminder(db)
}

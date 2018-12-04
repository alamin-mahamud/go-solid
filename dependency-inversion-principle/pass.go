package dependency

import "fmt"

// PasswordReminder -> reminds user about password changes throughEmail
func PasswordReminder(db IDb) {
	fmt.Printf("%T, %+v\n", db, db)
}

package postgresql

import "github.com/reyhanmichiels/bring_coffee/domain"

func Migrate() {
	DB.Migrator().DropTable(
		&domain.User{},
	)

	DB.AutoMigrate(
		&domain.User{},
	)
}

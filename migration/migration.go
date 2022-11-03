package migration

import "log"

func Migrate() {
	//Do Migration
	log.Println("run migration...")

	/*Fill here to run your migration*/

	//Example
	log.Println("example migration")
	if err := ExampleMigration(); err != nil {
		panic(err)
	} else {
		log.Println("success example migration")
	}
}

package web

import (
	"log"
	"os"
	"testing"

	"github.com/hfurubotten/autograder/database"
)

func TestMain(m *testing.M) {
	err := database.Start("/tmp/ag")
	if err != nil {
		log.Println("Unable to start database:", err)
		return
	}
	m.Run()
	err = database.Close()
	if err != nil {
		log.Println("Unable to close the database properly")
	}
	err = os.RemoveAll("/tmp/ag")
	if err != nil {
		log.Println("Unable to clean up database file from filesystem")
	}
}

package filesystem

import (
	"bear-export/db"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kennygrant/sanitize"
)

// GetApplicationDataPath returns a path of the Bear App content folder
func GetApplicationDataPath(username string) string {
	originalPath := "/Users/" + username + "/Library/Group Containers/"
	files, err := ioutil.ReadDir(originalPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() && strings.Contains(f.Name(), ".net.shinyfrog.bear") {
			return originalPath + f.Name() + "/Application Data/"
		}
	}

	panic("No Bear Notes Found")
}

// GetDatabasePath returns database file path based on application data path
func GetDatabasePath(applicationDataPath string) string {
	return applicationDataPath + "database.sqlite"
}

// WriteNotesToFolder writes all of the notes to the folder creating new subfolder with a timestamp
func WriteNotesToFolder(notes []db.Note, path string) {
	originatingFolder := path
	folder := originatingFolder + strconv.FormatInt(time.Now().Unix(), 10) + "/"
	err := os.Mkdir(folder, 0755)
	if err != nil {
		panic(err)
	}
	for _, note := range notes {
		fileName := sanitize.Path(note.Title + ".md")
		fileName = strings.ReplaceAll(fileName, "/", "-")

		f, err := os.Create(folder + fileName)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, err = f.Write([]byte(note.Body))
		if err != nil {
			panic(err)
		}

		f.Sync()
	}
}

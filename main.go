package main

import (
	"bear-export/db"
	"bear-export/filesystem"
	"log"

	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("bear-export").
		SetVersion("0.0.1").
		SetDescription("This tool exports bear notes")
	commando.
		Register(nil).
		AddArgument("username", "your user name", "").
		AddFlag("output,o", "output folder", commando.String, "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			username := args["username"].Value
			outputFolder, _ := flags["output"].GetString()

			if len(username) == 0 {
				log.Fatalln("Username must not be empty")
			}

			if len(outputFolder) == 0 {
				log.Fatalln("Output Folder (-o) must not be empty")
			}

			notes := db.ReadNotes(filesystem.GetDatabasePath(filesystem.GetApplicationDataPath(username)))
			filesystem.WriteNotesToFolder(notes, outputFolder)
		})

	commando.Parse(nil)
}

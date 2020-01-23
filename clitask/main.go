package main

import(
  "gophercises/clitask/cmd"
  "gophercises/clitask/db"
  "fmt"
  "os"
  "path/filepath"
)

func main() {
	dbPath := filepath.Join("tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
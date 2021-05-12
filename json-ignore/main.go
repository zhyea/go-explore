package main

import (
	"fmt"
	"json-ig/json"
	"json-ig/model"
)

func main() {

	w := model.Worker{
		Name:   "robin",
		Gender: "male",
		Site:   "www.zhyea.com",
		Job:    "coder",
	}

	w.SetTitle("Programmer")

	fmt.Printf("Title: %s \n", w.GetTitle())
	fmt.Printf("Job: %s \n", w.Job)

	str := json.ToJson(w)

	fmt.Println(str)
}

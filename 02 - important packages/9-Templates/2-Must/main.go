package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"Go", 40}
	t := template.Must(template.New("CouseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}}"))
	// tmp := template.New("CourseTemplate")
	// tmp, err := tmp.Parse("Course: {{.Name}} - Workload: {{.Workload}}")
	// if err != nil {
	// 	panic(err)
	// }
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

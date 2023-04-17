package pipeline

import (
	"log"
	"strings"
)

type Human struct {
	Name string
	Age  int
}

type IncAge struct {
	BaseStage[*Human]
}

// Just override Execute
func (s *IncAge) Execute(obj *Human) error {
	obj.Age++
	return nil
}

type UpperName struct {
	BaseStage[*Human]
}

// Just override Execute
func (s *UpperName) Execute(obj *Human) error {
	obj.Name = strings.ToUpper(obj.Name)
	return nil
}

func ExamplePipeline_Execute() {
	p, err := NewPipeline[*Human]([]Stage[*Human]{&IncAge{}, &UpperName{}, &IncAge{}})
	if err != nil {
		log.Fatal(err)
	}

	// create object to pass through pipeline
	m := &Human{Name: "name", Age: 10}
	if err := p.Execute(m); err != nil {
		log.Fatal(err)
	}
}

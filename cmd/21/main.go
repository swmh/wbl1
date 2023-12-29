package main

import (
	"context"
	"fmt"
)

type DependencyRealization struct{}

func (d *DependencyRealization) Do() {
	fmt.Println("do some work")
}

type DependencyAdapter struct {
	DependencyRealization
}

func (d *DependencyAdapter) Do(_ context.Context) {
	d.DependencyRealization.Do()
}

type Dependency interface {
	Do(ctx context.Context)
}

type Service struct {
	Dependency
}

// Реализовать паттерн «адаптер» на любом примере.
func main() {
	service := Service{
		Dependency: &DependencyAdapter{},
	}

	service.Do(context.Background())
}

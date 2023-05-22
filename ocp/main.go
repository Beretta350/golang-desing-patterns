package ocp

import "fmt"

//Open-Close Principle

func Main() {
	car := Product{"Car", blue, medium}
	house := Product{"House", red, large}
	tree := Product{"Tree", green, medium}
	apple := Product{"Apple", red, small}

	products := []Product{car, house, tree, apple}
	filter := Filter{}
	sizeSpec := SizeSpecification{medium}
	res := filter.Filter(products, sizeSpec)
	for _, v := range res {
		fmt.Printf(" - %s is medium\n", v.name)
	}

	colorSpec := ColorSpecification{blue}
	andSpec := AndSpecification{sizeSpec, colorSpec}
	res = filter.Filter(products, andSpec)
	for _, v := range res {
		fmt.Printf(" - %s is medium and blue\n", v.name)
	}
}

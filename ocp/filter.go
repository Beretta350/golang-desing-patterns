package ocp

type Filter struct {
}

func (f Filter) Filter(products []Product, spec Specification) (results []Product) {
	for _, prod := range products {
		if spec.IsSatisfied(&prod) {
			results = append(results, prod)
		}
	}
	return
}

package ocp

type Specification interface {
	IsSatisfied(p *Product) bool
}

type SizeSpecification struct {
	size Size
}

type ColorSpecification struct {
	color Color
}

type NameSpecification struct {
	name string
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

func (s ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == s.color
}

func (s NameSpecification) IsSatisfied(p *Product) bool {
	return p.name == s.name
}

package ocp

type Size int
type Color int

const (
	small Size = iota
	medium
	large
)

const (
	red Color = iota
	green
	blue
)

type Product struct {
	name  string
	color Color
	size  Size
}

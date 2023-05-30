package builder

import "fmt"

func Main() {
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b.String())
}

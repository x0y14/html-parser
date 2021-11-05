package html_parser

type NodeKind int

const (
	_ NodeKind = iota
	PlainText
	HtmlTag
)

type Node struct {
	name     string
	kind     NodeKind
	attrs    []Attribute
	children []Node
}

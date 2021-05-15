package xml

type XMLElem struct {
	Name        string
	Content     string
	SelfClosing bool

	Attributes []*XMLAttr
	Elements   []*XMLElem
}

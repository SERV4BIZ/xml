package xml

type XMLDoc struct {
	Version  string
	Encoding string

	Elements []*XMLElem
}

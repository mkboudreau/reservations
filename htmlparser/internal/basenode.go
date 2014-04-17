package internal

type BaseNode struct {
	FullHtml    string
	ContentHtml string
	Children    []HtmlNode
}

func (node *BaseNode) IsDone(data *ParseData) bool {
	return !data.HasNext()
}
func (node *BaseNode) AddChild(child ParsingHtmlNode) {
}

func (node *BaseNode) ProcessAllParts(data *ParseData) {
}
func (node *BaseNode) ProcessOpenTag(data *ParseData) {

}
func (node *BaseNode) ProcessCloseTag(data *ParseData) {

}
func (node *BaseNode) ProcessContent(data *ParseData) {
}

func (node *BaseNode) AddHtmlCharacter(ch string) {
	node.appendToFullHtml(ch)
}
func (node *BaseNode) AddContentCharacter(ch string) {
	node.appendToContentHtml(ch)
	node.appendToFullHtml(ch)
}
func (node *BaseNode) appendToFullHtml(ch string) {
	node.FullHtml = appendToString(node.FullHtml, ch)
}
func (node *BaseNode) appendToContentHtml(ch string) {
	node.ContentHtml = appendToString(node.ContentHtml, ch)
}

/*
	External Html Getters
*/

func (node *BaseNode) InnerHtml() string {
	return node.ContentHtml
}
func (node *BaseNode) InnerText() string {
	return node.ContentHtml
}
func (node *BaseNode) Html() string {
	return node.FullHtml
}
func (node *BaseNode) GetChildren() HtmlNodeList {
	if node.Children == nil {
		logger.Trace("  <initializing children for node>")
		node.Children = make([]HtmlNode, 0, 10)
	}
	return node.Children
}
func (node *BaseNode) GetType() HtmlNodeType {
	return HtmlUnknownNode
}

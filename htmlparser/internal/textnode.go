package internal

import (
	"fmt"
	"reflect"
)

type TextNode struct {
	BaseNode
}

////// entity methods methods

func (node *TextNode) GetType() HtmlNodeType {
	return HtmlTextNode
}

func (node *TextNode) AddChild(child ParsingHtmlNode) {
	logger.Debug("Adding New Child:", child, " to:", node)
	logger.Debug("Type of child:", reflect.TypeOf(child))
	logger.Trace("Count of Children (Before): ", len(node.GetChildren()))
	node.Children = append(node.GetChildren(), child)
	logger.Trace("Count of Children (After): ", len(node.GetChildren()))
}

////// processing methods

func (node *TextNode) ProcessAllParts(data *ParseData) {
	node.ProcessOpenTag(data)
	node.ProcessContent(data)
	node.ProcessCloseTag(data)
}
func (node *TextNode) ProcessOpenTag(data *ParseData) {
	// do nothing
}
func (node *TextNode) ProcessCloseTag(data *ParseData) {
	// do nothing
}
func (node *TextNode) ProcessContent(data *ParseData) {

	logger.Trace(" -- processing text content -- first char [", data.PeekCurrent(), "]", " at index:", data.Index)
	for data.HasNext() && !data.IsNext(StartOpenTagToken) {

		logger.Trace("    > text peek: ", data.PeekCurrent(), " with index ", data.Index)
		node.AddContentCharacter(data.Next())
	}
}
func (node *TextNode) IsDone(data *ParseData) bool {
	logger.Trace("IsDone Text:", data.PeekCurrent())
	return !data.HasNext() || data.IsNext(StartOpenTagToken)
}

/*
	External Html Getters
*/

func (node *TextNode) InnerHtml() string {
	return node.ContentHtml
}
func (node *TextNode) InnerText() string {
	return node.ContentHtml
}
func (node *TextNode) Html() string {
	return node.FullHtml
}

/*
	Stringer
*/

func (node *TextNode) String() string {
	return fmt.Sprintf("Type: %v; with Text: %v", node.GetType(), node.FullHtml)
}

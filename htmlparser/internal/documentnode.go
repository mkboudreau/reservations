package internal

import (
	"fmt"
	"reflect"
	"strings"
)

type DocumentNode struct {
	BaseNode
}

////// entity methods methods

func (node *DocumentNode) GetType() HtmlNodeType {
	return HtmlDocumentNode
}

func (node *DocumentNode) AddChild(child ParsingHtmlNode) {
	logger.Trace(" - document - adding child of type:", reflect.TypeOf(child))
	node.Children = append(node.GetChildren(), child)
}

////// processing methods

func (node *DocumentNode) ProcessAllParts(data *ParseData) {
	node.ProcessOpenTag(data)
	node.ProcessContent(data)
	node.ProcessCloseTag(data)
}
func (node *DocumentNode) ProcessOpenTag(data *ParseData) {
	// do nothing
}
func (node *DocumentNode) ProcessCloseTag(data *ParseData) {
	// do nothing
}
func (node *DocumentNode) ProcessContent(data *ParseData) {

	logger.Trace(" -- processing document content -- first char [", data.PeekCurrent(), "]", " at index:", data.Index)
	for !node.IsDone(data) {
		//ch := data.Next()
		child := getNode(data)
		child.ProcessAllParts(data)
		node.AddChild(child)
	}

}

/*
	External Html Getters
*/

func (node *DocumentNode) InnerHtml() string {
	children := node.GetChildren()
	strSlice := make([]string, len(children))
	for i := 0; i < len(children); i++ {
		strSlice[i] = children[i].InnerHtml()
	}
	return strings.Join(strSlice, "")
	//return node.ContentHtml
}
func (node *DocumentNode) InnerText() string {
	children := node.GetChildren()
	strSlice := make([]string, len(children))
	for i := 0; i < len(children); i++ {
		strSlice[i] = children[i].InnerText()
	}
	return strings.Join(strSlice, "")
	//return node.ContentHtml
}
func (node *DocumentNode) Html() string {
	return node.InnerHtml()
}

/*
	Stringer
*/

func (node *DocumentNode) String() string {
	return fmt.Sprintf("Type: %v; %v children: %v ", node.GetType(), len(node.GetChildren()), node.GetChildren())
}

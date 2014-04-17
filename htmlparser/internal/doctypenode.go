package internal

import (
	"fmt"
)

type DocTypeNode struct {
	BaseNode
}

////// entity methods methods

func (node *DocTypeNode) GetType() HtmlNodeType {
	return HtmlDoctypeNode
}

////// processing methods

func (node *DocTypeNode) ProcessAllParts(data *ParseData) {
	node.ProcessOpenTag(data)
	node.ProcessContent(data)
	node.ProcessCloseTag(data)
}
func (node *DocTypeNode) ProcessOpenTag(data *ParseData) {
	// looking to add <!DOCTYPE
	if data.IsNextCaseInsensitive(DoctypeOpenToken) {
		node.AddHtmlCharacter(data.SeekTo(len(DoctypeOpenToken)))
	}
}
func (node *DocTypeNode) ProcessCloseTag(data *ParseData) {
	// looking to add ">"
	node.addEndingCharsLeft(data)
}
func (node *DocTypeNode) ProcessContent(data *ParseData) {
	// looking to add everything betweeen <!DOCTYPE and >
	for data.HasNext() && !node.onlyEndingCharsLeft(data) {
		node.AddHtmlCharacter(data.Next())
	}
}
func (node *DocTypeNode) onlyEndingCharsLeft(data *ParseData) bool {
	return data.IsNext(EndTagToken)
}
func (node *DocTypeNode) addEndingCharsLeft(data *ParseData) {
	if node.onlyEndingCharsLeft(data) {
		node.AddHtmlCharacter(data.SeekTo(len(EndTagToken)))
	}
}

/*
	Stringer
*/

func (node *DocTypeNode) String() string {
	return fmt.Sprintf("Type: %v; with Text: %v", node.GetType(), node.FullHtml)
}

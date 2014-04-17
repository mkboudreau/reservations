package internal

import (
	"fmt"
)

type CommentNode struct {
	BaseNode
}

////// entity methods methods

func (node *CommentNode) String() string {
	return fmt.Sprintf("Type: %v; with Text: %v", node.GetType(), node.FullHtml)
}
func (node *CommentNode) GetType() HtmlNodeType {
	return HtmlCommentNode
}

////// processing methods

func (node *CommentNode) ProcessAllParts(data *ParseData) {
	node.ProcessOpenTag(data)
	node.ProcessContent(data)
	node.ProcessCloseTag(data)
}
func (node *CommentNode) ProcessOpenTag(data *ParseData) {
	// looking to add <!--
	if data.IsNext(CommentOpenToken) {
		node.AddHtmlCharacter(data.SeekTo(len(CommentOpenToken)))
	}
}
func (node *CommentNode) ProcessCloseTag(data *ParseData) {
	// looking to add "-->"
	node.addEndingCharsLeft(data)
}

func (node *CommentNode) ProcessContent(data *ParseData) {
	// looking to add everything betweeen <!DOCTYPE and >
	for data.HasNext() && !node.onlyEndingCharsLeft(data) {
		node.AddHtmlCharacter(data.Next())
	}
}

func (node *CommentNode) onlyEndingCharsLeft(data *ParseData) bool {
	return data.IsNext(CommentCloseToken)
}
func (node *CommentNode) addEndingCharsLeft(data *ParseData) {
	if node.onlyEndingCharsLeft(data) {
		node.AddHtmlCharacter(data.SeekTo(len(CommentCloseToken)))
	}
}

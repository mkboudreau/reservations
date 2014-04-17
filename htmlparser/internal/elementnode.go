package internal

import (
	"fmt"
	"reflect"
	"strings"
)

type ElementNode struct {
	BaseNode
	TagName           string
	RawTagText        string
	HasContent        bool
	Attributes        map[string]string
	processedOpenTag  bool
	processedContent  bool
	processedCloseTag bool
}

////// entity methods methods

func (node *ElementNode) GetType() HtmlNodeType {
	return HtmlElementNode
}
func (node *ElementNode) AddChild(child ParsingHtmlNode) {
	logger.Trace(" - element - adding child of type:", reflect.TypeOf(child))
	node.Children = append(node.GetChildren(), child)
}

////// processing methods

func (node *ElementNode) ProcessAllParts(data *ParseData) {
	node.ProcessOpenTag(data)
	node.ProcessContent(data)
	node.ProcessCloseTag(data)
}
func (node *ElementNode) ProcessOpenTag(data *ParseData) {
	for data.HasNext() && !node.isOpenTagDone(data) {
		logger.Trace("    > element (open tag) peek: ", data.PeekCurrent(), " with index ", data.Index)
		node.AddHtmlCharacter(data.Next())
	}

	if node.doesStartTagCloseWithNoContent(data) {
		tmpIndex := data.Index
		finalChars := data.SeekTo(len(EndNoTagToken))
		logger.Trace("    > element (open tag) peek: ", finalChars, " with index ", tmpIndex)
		node.HasContent = false
		node.AddHtmlCharacter(finalChars)
	} else {
		tmpIndex := data.Index
		finalChars := data.SeekTo(len(EndTagToken))
		logger.Trace("    > element (open tag) peek: ", finalChars, " with index ", tmpIndex)
		node.HasContent = true
		node.AddHtmlCharacter(finalChars)
	}

	node.processedOpenTag = true
}

func (node *ElementNode) isOpenTagDone(data *ParseData) bool {
	if node.processedOpenTag {
		return true
	} else if data.IsNext(EndNoTagToken) {
		return true
	} else if data.IsNext(EndTagToken) {
		return true
	} else {
		return false
	}
}
func (node *ElementNode) doesStartTagCloseWithNoContent(data *ParseData) bool {
	return data.IsNext(EndNoTagToken)
}

func (node *ElementNode) ProcessCloseTag(data *ParseData) {

	for data.HasNext() && !node.isCloseTagDone(data) {
		logger.Trace("    > element (close tag) peek: ", data.PeekCurrent(), " with index ", data.Index)
		node.AddHtmlCharacter(data.Next())
	}
	if data.IsNext(EndTagToken) {
		logger.Trace("    > element (close tag) peek: ", data.PeekCurrent(), " with index ", data.Index)
		node.AddHtmlCharacter(data.SeekTo(len(EndTagToken)))
	}

	node.processedCloseTag = true
}
func (node *ElementNode) isCloseTagDone(data *ParseData) bool {
	if node.processedCloseTag {
		return true
	} else if data.IsNext(EndTagToken) {
		return true
	} else {
		return false
	}
}
func (node *ElementNode) ProcessContent(data *ParseData) {

	logger.Trace(" -- processing element content -- first char [", data.PeekCurrent(), "]", " at index:", data.Index)
	for data.HasNext() && node.HasContent && !node.IsDone(data) {
		child := getNode(data)
		child.ProcessAllParts(data)
		node.AddChild(child)
	}
	node.processedContent = true
}
func (node *ElementNode) IsDone(data *ParseData) bool {
	// gets called between children
	return !data.HasNext() || node.isContentDone(data)
}
func (node *ElementNode) isContentDone(data *ParseData) bool {
	if node.processedContent || !node.HasContent {
		return true
	} else if node.processedOpenTag && data.IsNext(StartCloseTagToken) {
		return true
	} else {
		return false
	}
}

/*
	External Html Getters
*/

func (node *ElementNode) InnerHtml() string {
	children := node.GetChildren()
	strSlice := make([]string, len(children))
	for i := 0; i < len(children); i++ {
		strSlice[i] = children[i].InnerHtml()
	}
	return strings.Join(strSlice, "")
	//return node.ContentHtml
}
func (node *ElementNode) InnerText() string {
	children := node.GetChildren()
	strSlice := make([]string, len(children))
	for i := 0; i < len(children); i++ {
		strSlice[i] = children[i].InnerText()
	}
	return strings.Join(strSlice, "")
	//return node.ContentHtml
}
func (node *ElementNode) Html() string {
	children := node.GetChildren()
	strSlice := make([]string, len(children))
	for i := 0; i < len(children); i++ {
		strSlice[i] = children[i].Html()
	}
	return strings.Join(strSlice, "")
}

/*
	Stringer
*/

func (node *ElementNode) String() string {
	var msgs string
	nodelist := node.GetChildren()
	for i := 0; i < len(nodelist); i++ {
		node := nodelist[i]
		tmp := fmt.Sprintf("{ Type: %T, Address: %p, Children: [%v] %v }", node, node, len(node.GetChildren()), node.GetChildren())
		if i == 0 {
			msgs = fmt.Sprintf("[ %v", tmp)
		} else {
			msgs = fmt.Sprintf("%v, %v", msgs, tmp)
		}
	}
	if len(nodelist) > 0 {
		msgs = fmt.Sprintf("%v ]", msgs)
	}
	return msgs
}

package internal

import "fmt"

type HtmlNodeType int

const (
	HtmlCommentNode HtmlNodeType = iota
	HtmlDocumentNode
	HtmlTextNode
	HtmlElementNode
	HtmlDoctypeNode
	HtmlAttributeNode
	HtmlUnknownNode
)

type HtmlNodeList []HtmlNode
type HtmlNode interface {
	InnerHtml() string
	InnerText() string
	Html() string
	GetType() HtmlNodeType
	GetChildren() HtmlNodeList
}

func (nodetype HtmlNodeType) String() string {
	switch nodetype {
	case HtmlCommentNode:
		return "Comment"
	case HtmlDocumentNode:
		return "Document"
	case HtmlTextNode:
		return "Text"
	case HtmlDoctypeNode:
		return "Doctype"
	case HtmlAttributeNode:
		return "Attribute"
	case HtmlElementNode:
		return "Element"
	}
	return "unknown"
}
func (nodelist HtmlNodeList) String() string {
	var msgs string

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

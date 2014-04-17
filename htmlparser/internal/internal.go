package internal

import (
	"github.com/mkboudreau/loggo"
	"reflect"
	"strings"
)

var logger *loggo.LevelLogger = loggo.DefaultLevelLogger()

type HtmlParsingEngine struct {
	Data *ParseData
}

func NewHtmlParsingEngine(html string) *HtmlParsingEngine {
	engine := new(HtmlParsingEngine)
	engine.Data = &ParseData{Index: 0, Data: html}
	return engine
}

func (parser *HtmlParsingEngine) Parse() HtmlNode {

	doc := new(DocumentNode)
	doc.ProcessAllParts(parser.Data)

	return doc
}

/////////////////////////////////////////

const (
	StartOpenTagToken         string = "<"
	StartCloseTagToken               = "</"
	EndNoTagToken                    = "/>"
	EndTagToken                      = ">"
	DoctypeOpenToken                 = "<!DOCTYPE"
	DoctypeTokenPart2                = "DOCTYPE"
	DoctypeOrCommentOpenToken        = "<!"
	CommentToken                     = "!--"
	CommentTokenPart2And3            = "-"
	CommentCloseToken                = "-->"
	CommentOpenToken                 = "<!--"
)

type ParsingHtmlNode interface {
	HtmlNode
	IsDone(data *ParseData) bool
	AddChild(child ParsingHtmlNode)

	ProcessAllParts(data *ParseData)
	ProcessContent(data *ParseData)
	ProcessOpenTag(data *ParseData)
	ProcessCloseTag(data *ParseData)

	AddHtmlCharacter(ch string)
	AddContentCharacter(ch string)
}

func EnsureCapacityForLength(nodes *[]HtmlNode, neededLength int) {
	if neededLength > cap(*nodes) {
		oldValues := *nodes
		*nodes = make([]HtmlNode, neededLength, neededLength*2)
		copy(*nodes, oldValues)
	} else if neededLength > len(*nodes) {
		oldValues := *nodes
		*nodes = oldValues[:neededLength]
	}
}
func appendToString(originalString string, ch string) string {
	both := []string{originalString, ch}
	return strings.Join(both, "")
}

func getNode(data *ParseData) ParsingHtmlNode {
	var node ParsingHtmlNode
	if !data.HasNext() {
		logger.Warn("Returning a nil ParsingHtmlNode")
		return (ParsingHtmlNode)(nil)
	}
	ch := data.PeekCurrent()
	if ch == StartOpenTagToken {
		if data.IsNext(CommentOpenToken) {
			node = new(CommentNode)
		} else if data.IsNextCaseInsensitive(DoctypeOpenToken) {
			node = new(DocTypeNode)
		} else if data.IsNext(DoctypeOrCommentOpenToken) {
			// dunno : <!, but not <!DOCTYPE or <!--
			panic("invalid character sequence found <! that doesn't match <!DOCTYPE or <!--")
		} else {
			//element
			node = new(ElementNode)
		}
		//node.AddHtmlCharacter(ch)
	} else {
		//text
		node = new(TextNode)
		//node.AddContentCharacter(ch)
	}

	logger.Debug("Returning New Node:", reflect.TypeOf(node))
	return node
}

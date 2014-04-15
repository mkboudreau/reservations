package internal

import (
	"log"
	"os"
	s "strings"
)

type InternalParseType int

const (
	InternalParseContent InternalParseType = iota
	InternalParseEndTag
	InternalParseStartTag
	InternalParseNoContentTag
)

type InternalParseNode struct {
	Type    InternalParseType
	Content string
}

type InternalParseNodes []InternalParseNode

type InternalHtmlParsingEngine struct {
	StartOpenTagToken  string
	StartCloseTagToken string
	EndTagToken        string
	EndNoTagToken      string
	debugFlag          bool
	logger             *log.Logger
}

func (parseType InternalParseType) String() string {
	switch parseType {
	case InternalParseContent:
		return "InternalParseType"
	case InternalParseEndTag:
		return "InternalParseEndTag"
	case InternalParseStartTag:
		return "InternalParseStartTag"
	case InternalParseNoContentTag:
		return "InternalParseNoContentTag"
	}
	return ""
}

func NewInternalHtmlParsingEngine() *InternalHtmlParsingEngine {
	engine := new(InternalHtmlParsingEngine)
	engine.StartOpenTagToken = "<"
	engine.StartCloseTagToken = "</"
	engine.EndNoTagToken = "/>"
	engine.EndTagToken = ">"
	engine.debugFlag = false
	engine.logger = buildDefaultLogger()

	return engine
}
func buildDefaultLogger() *log.Logger {
	return log.New(os.Stdout, "Internal Html Parsing >> ", log.Lmicroseconds)
}

func (engine *InternalHtmlParsingEngine) TurnOnDebugging() *InternalHtmlParsingEngine {
	engine.debugFlag = true
	return engine
}
func (engine *InternalHtmlParsingEngine) TurnOffDebugging() *InternalHtmlParsingEngine {
	engine.debugFlag = false
	return engine
}
func (engine *InternalHtmlParsingEngine) debug(msg ...interface{}) *InternalHtmlParsingEngine {
	if engine.debugFlag {
		engine.logger.Println(msg)
	}

	return engine
}

func (engine *InternalHtmlParsingEngine) TokenizeHtmlStringIntoParseNodes(html string) InternalParseNodes {
	engine.debug("HTML [", html, "]")
	token := engine.EndNoTagToken
	splitAfterEnds := s.SplitAfter(html, token)
	if len(splitAfterEnds) == 1 {
		engine.debug("Did not find [", token, "] for splitting main html string")
		return engine.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(html)
	} else {
		engine.debug("Found ["+token+"] for splitting main html string into [", len(splitAfterEnds), "] parts")
		return engine.buildParseNodesFromNoContentTagSplit(splitAfterEnds, s.Count(html, token))
	}
}

func (engine *InternalHtmlParsingEngine) buildParseNodesFromNoContentTagSplit(noContentEndSplit []string, estimatedTotal int) InternalParseNodes {
	parseNodes := make(InternalParseNodes, estimatedTotal)
	currentIndex := 0
	for _, stringWithStartChar := range noContentEndSplit {
		var tag, content string

		engine.debug("stringWithStartChar:", stringWithStartChar)
		if len(stringWithStartChar) == 0 || len(s.TrimSpace(stringWithStartChar)) == 0 {
			engine.debug("   - ignoring")
			continue
		}

		if openTagIndex := s.LastIndex(stringWithStartChar, engine.StartOpenTagToken); openTagIndex != -1 {
			tag = stringWithStartChar[openTagIndex:]
			content = stringWithStartChar[:openTagIndex]

			tokenizedRegularNodes := engine.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(string(content))
			currentIndex = parseNodes.addSliceAtIndex(tokenizedRegularNodes, currentIndex)
			engine.debug("   - tag:", tag)
			parseNodes[currentIndex] = InternalParseNode{Type: InternalParseNoContentTag, Content: tag}
			currentIndex++
		} else {
			engine.debug("PANIC!", "stringWithStartChar:", stringWithStartChar, "currentIndex:", currentIndex, "noContentEndSplit length:", len(noContentEndSplit))
			panic("no start tag found for end tag of contentless element")
		}
	}

	return parseNodes[:currentIndex]
}

func (slice *InternalParseNodes) addSliceAtIndex(newSlice InternalParseNodes, index int) (nextIndex int) {
	if len(newSlice) > 0 {
		slice.ensureCapacity(len(newSlice) + index)
		tmpSlice := *slice
		//*slice = tmpSlice[index:]
		copy(tmpSlice[index:], newSlice)
		nextIndex = index + len(newSlice)
	} else {
		nextIndex = index
	}

	return nextIndex
}

func (slice *InternalParseNodes) ensureCapacity(neededLength int) {
	if neededLength > cap(*slice) {
		oldValues := *slice
		*slice = make(InternalParseNodes, neededLength, neededLength*2)
		copy(*slice, oldValues)
	} else if neededLength > len(*slice) {
		oldValues := *slice
		*slice = oldValues[:neededLength]
	}
}

func (engine *InternalHtmlParsingEngine) TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(html string) InternalParseNodes {
	engine.debug("HTML [", html, "]")
	token := engine.EndTagToken
	splitAfterEnds := s.SplitAfter(html, token)

	if len(splitAfterEnds) == 1 {
		engine.debug("Did not find [", token, "] for splitting regular tags")
		return make(InternalParseNodes, 0)
	} else {
		engine.debug("Found [", token, "] for splitting regular tags into [", len(splitAfterEnds), "] parts")
		return engine.buildParseNodesFromContentTagSplit(splitAfterEnds)
	}
}

func (engine *InternalHtmlParsingEngine) buildParseNodesFromContentTagSplit(contentEndSplit []string) InternalParseNodes {
	parseNodes := make(InternalParseNodes, len(contentEndSplit)*2)
	currentIndex := 0
	for _, stringWithStartChar := range contentEndSplit {
		var tag, content string
		var tagType InternalParseType

		engine.debug("stringWithStartChar:", stringWithStartChar)
		if len(stringWithStartChar) == 0 || len(s.TrimSpace(stringWithStartChar)) == 0 {
			engine.debug("   - ignoring")
			continue
		}

		if openTagIndex := s.LastIndex(stringWithStartChar, engine.StartCloseTagToken); openTagIndex != -1 {
			tag = stringWithStartChar[openTagIndex:]
			content = stringWithStartChar[:openTagIndex]
			tagType = InternalParseEndTag
		} else if openTagIndex := s.LastIndex(stringWithStartChar, engine.StartOpenTagToken); openTagIndex != -1 {
			tag = stringWithStartChar[openTagIndex:]
			content = stringWithStartChar[:openTagIndex]
			tagType = InternalParseStartTag
		} else {
			engine.debug("PANIC!", "stringWithStartChar:", stringWithStartChar, "currentIndex:", currentIndex, "contentEndSplit length:", len(contentEndSplit))
			panic("no start tag found for end tag for content element")
		}

		if len(content) != 0 {
			engine.debug("   - content:", content)
			parseNodes[currentIndex] = InternalParseNode{Type: InternalParseContent, Content: content}
			currentIndex++
		}
		engine.debug("   - tag:", tag)
		parseNodes[currentIndex] = InternalParseNode{Type: tagType, Content: tag}
		currentIndex++
	}

	return parseNodes[:currentIndex]
}

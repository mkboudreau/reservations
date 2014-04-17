package internal_test

import (
	. "github.com/mkboudreau/reservations/htmlparser/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"github.com/mkboudreau/loggo"
	"reflect"
	"testing"
)

func TestInternalHtmlParsing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html Parsing Internal Test Suite")
}

var logger *loggo.LevelLogger = loggo.DefaultLevelLogger()

var _ = Describe("Html Parsing Internal Tests", func() {

	logger.SetTrace()
	logger.LoggingPrefix("[Test] ")

	Context("", func() {

		It("PLACE HOLDER TEST: DOC WITH 1 ELEMENT NODE", func() {
			parser := NewHtmlParsingEngine(HtmlSingleRegularNodeWithContent)
			document := parser.Parse()

			nodes := document.GetChildren()

			Expect(nodes).ToNot(BeNil())
			fmt.Println(document)
			//fmt.Println(nodes)
		})
	})

})

const HtmlSingleRegularNodeWithContent = "<div>Hello</div>"
const HtmlSingleRegularNodeWithIdAndContent = `<div id="world">Hello</div>`
const HtmlSingleNoContentNodeWithId = `<div id="world" />`
const HtmlSingleNoContentNode = `<div />`
const HtmlMultipleNoContentNode = `<div id="hello"><div /></div>`

const InvalidHtmlNoOpeningStartTag = `div id="hello">Hello World</div>`
const InvalidHtmlNoOpeningEndTag = `<div id="hello"Hello World</div>`
const InvalidHtmlNoClosingStartTag = `<div id="hello">Hello World/div>`
const InvalidHtmlNoClosingEndTag = `<div id="hello">Hello World</div`
const InvalidHtmlNoOpeningEndTagSlash = `<div id="hello">Hello World<div>`
const InvalidHtmlPlainString = `html`

const HtmlWithDocTypeTest = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

<html>
<head><title>Hello</title></head>
<body>
<h1>Heading</h1>
</body>
</html>
`

const HtmlWithCommentsTest = `
<html>
<head><title>Hello</title></head>
<body>
<!-- Watch Out! -->
<h1>Heading</h1>

<!-- <h2>This should not be a node</h2> -->
</body>
</html>

`

const HtmlAllVariationTest = `

<html>
<head><title>Hello</title>

</head>

<body>

<div id="id1" class="someclass"> <h1>Test H1</h1></div>

<div id="outer" >
	<div id="middle1" abc="aaa">Test</div>
	<div id="middle2" abc="bbb">Test 2</div>
	<div id="middle3" abc="ccc ddd">Test 3</div>
	<div id="middle4" ><div id="nocontent" /></div>
</div>
<div></div>

<div>       </div>
</body>
</html>

`

/* private helper functions */

func expectFunctionToPanic(done Done, fn interface{}) {
	expectFunctionToPanicAsSpecified(done, fn, true)
}
func expectFunctionToNotPanic(done Done, fn interface{}) {
	expectFunctionToPanicAsSpecified(done, fn, false)
}
func expectFunctionToPanicAsSpecified(done Done, fn interface{}, shouldPanic bool) {
	go func() {
		defer GinkgoRecover()

		panicFunc := func() {
			val := reflect.ValueOf(fn)
			if val.Kind() == reflect.Func {
				in := make([]reflect.Value, 0)
				val.Call(in)
			}
		}

		if shouldPanic {
			Expect(panicFunc).To(Panic())
		} else {
			Expect(panicFunc).ToNot(Panic())
		}

		close(done)
	}()
}

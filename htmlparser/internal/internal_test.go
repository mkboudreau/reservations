package internal_test

import (
	. "github.com/mkboudreau/reservations/htmlparser/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"reflect"
	"testing"
)

func TestGoTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html Parsing INTERNALS Test Suite")
}

var _ = Describe("Html Parsing Internal Tests", func() {

	parser := NewInternalHtmlParsingEngine()
	parser.TurnOnDebugging()

	Context("INTERNAL FLAT STRUCTURE: Valid HTML Text", func() {
		It("Should contain 3 nodes: start, content, end", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleRegularNodeWithContent)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[1].Type).To(Equal(InternalParseContent))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
		})
		It("Should contain 3 nodes: start, content, end", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleRegularNodeWithIdAndContent)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[1].Type).To(Equal(InternalParseContent))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
		})
		It("Should contain 1 node", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleNoContentNodeWithId)
			Expect(nodes).To(HaveLen(1))
			// this is obviously not correct; however, this method is only supposed to be called internally after these tags have been processed (note the name of the method)
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
		})
		It("Should contain 1 node", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlSingleNoContentNode)
			Expect(nodes).To(HaveLen(1))
			// this is obviously not correct; however, this method is only supposed to be called internally after these tags have been processed (note the name of the method)
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
		})
		It("Should contain 3 nodes: start, content, end", func() {
			nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlMultipleNoContentNode)
			Expect(nodes).To(HaveLen(3))
			Expect(nodes[0].Type).To(Equal(InternalParseStartTag))
			// this is obviously not correct; however, this method is only supposed to be called internally after these tags have been processed (note the name of the method)
			Expect(nodes[1].Type).To(Equal(InternalParseStartTag))
			Expect(nodes[2].Type).To(Equal(InternalParseEndTag))
			fmt.Println(nodes)
		})
			It("Should correctly process the full html set", func() {
					nodes := parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(HtmlAllVariationTest)
					Expect(nodes).ToNot(BeNil())
					Expect(nodes).ToNot(HaveLen(0))
				})
	})
	Context("Invalid HTML Text", func() {
		It("Should panic on no opening start tag", func(done Done) {
			expectFunctionToPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoOpeningStartTag) })
		})
		It("Should panic on no opening end tag", func(done Done) {
			//invalid, but this particular method should not panic
			expectFunctionToNotPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoOpeningEndTag) })
		})
		It("Should panic on no closing start tag", func(done Done) {
			//invalid, but this particular method should not panic
			expectFunctionToPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoClosingStartTag) })
		})
		It("Should not panic on no closing end tag", func(done Done) {
				//invalid, but this particular method should not panic
			expectFunctionToNotPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoClosingEndTag) })
		})
		It("Should panic on no opening slash inside end tag", func(done Done) {
			//invalid, but this particular method should not panic
			expectFunctionToNotPanic(done, func() { parser.TokenizeHtmlStringWithOnlyRegularTagsIntoParseNodes(InvalidHtmlNoOpeningEndTagSlash) })
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

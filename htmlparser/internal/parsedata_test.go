package internal_test

import (
	. "github.com/mkboudreau/reservations/htmlparser/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//"fmt"
	//"reflect"
	//"testing"
)

var _ = Describe("Parse Data Tests", func() {

	var data *ParseData

	BeforeEach(func() {
		data = new(ParseData)
		data.Index = 0
	})

	Context("Next", func() {

		It("Next Tests", func() {
			data.Data = "hello world"

			index := 1
			Expect(data.Next()).To(Equal("h"), "call "+string(index))
			index++
			Expect(data.Next()).To(Equal("e"), "call "+string(index))
			index++
			Expect(data.Next()).To(Equal("l"), "call "+string(index))
			index++
			Expect(data.Next()).To(Equal("l"), "call "+string(index))
			index++
			Expect(data.Next()).To(Equal("o"), "call "+string(index))
		})

		It("Has Next At Beginning", func() {
			data.Data = "hello world"

			Expect(data.HasNext()).To(BeTrue())
		})
		It("Has Next In Middle", func() {
			data.Data = "hello world"

			Expect(data.HasNext()).To(BeTrue())
		})
		It("Has Next At End", func(done Done) {
			data.Data = "hello world"
			data.Index = 9

			Expect(data.PeekCurrent()).To(Equal("l"), "Peek Current")
			Expect(data.HasNext()).To(BeTrue(), "Has Next [1]")
			Expect(data.Next()).To(Equal("l"), "Next l")
			Expect(data.HasNext()).To(BeTrue(), "Has Next [2]")
			Expect(data.Next()).To(Equal("d"), "Next d")
			Expect(data.HasNext()).To(BeFalse(), "Has Next [3]")

			expectFunctionToPanic(done, func() {
				data.Next()
			})
		})
	})
	Context("Peek", func() {

		It("Peek Current Tests", func() {
			data.Data = "hello world"

			Expect(data.PeekCurrent()).To(Equal("h"))
			data.Index++
			Expect(data.PeekCurrent()).To(Equal("e"))
		})
		It("Peek Next Tests", func() {
			data.Data = "hello world"

			Expect(data.PeekNext()).To(Equal("e"))
			data.Index++
			Expect(data.PeekNext()).To(Equal("l"))
		})

		It("Peek To Tests", func() {
			data.Data = "hello world"

			Expect(data.PeekTo(2)).To(Equal("he"))
		})
		It("Peek To Tests at End", func() {
			data.Data = "hello world"
			data.Index = 8

			Expect(data.PeekCurrent()).To(Equal("r"))
			Expect(data.PeekTo(2)).To(Equal("rl"))
			Expect(data.PeekTo(3)).To(Equal("rld"))
			Expect(data.PeekTo(4)).To(Equal("rld"))
		})
		It("Peek By Tests", func() {
			data.Data = "hello world"

			Expect(data.PeekBy(2)).To(Equal("l"))
		})
		It("Peek By Tests At End", func() {
			data.Data = "hello world"
			data.Index = 8

			Expect(data.PeekCurrent()).To(Equal("r"))
			Expect(data.PeekBy(2)).To(Equal("d"))
			Expect(data.PeekBy(3)).To(BeEmpty())
		})

	})
	Context("Seek", func() {

		It("get 2 from beginning", func() {
			data.Data = "hello world"

			Expect(data.PeekCurrent()).To(Equal("h"))
			Expect(data.SeekTo(2)).To(Equal("he"))
			Expect(data.PeekCurrent()).To(Equal("l"))
		})
		It("get 4 chars", func() {
			data.Data = "hello world"
			data.Index = 4

			Expect(data.PeekCurrent()).To(Equal("o"))
			Expect(data.SeekTo(4)).To(Equal("o wo"))
			Expect(data.PeekCurrent()).To(Equal("r"))
		})
		It("should panic", func(done Done) {
			data.Data = "hello world"
			data.Index = 9

			Expect(data.PeekCurrent()).To(Equal("l"), "Peek Current")

			expectFunctionToPanic(done, func() {
				data.SeekTo(3)
			})
		})
	})
	Context("IS NEXT", func() {
		It("check from beginning", func() {
			data.Data = "hello world"
			data.Index = 1

			Expect(data.PeekCurrent()).To(Equal("e"))
			Expect(data.IsNext("el")).To(BeTrue())
			Expect(data.IsNext("ello")).To(BeTrue())
			Expect(data.IsNext("ll")).To(BeFalse())

			Expect(data.Next()).To(Equal("e"))
			Expect(data.Next()).To(Equal("l"))
		})
		It("basic", func() {
			data.Data = "hello world"
			data.Index = 1

			Expect(data.PeekCurrent()).To(Equal("e"))
			Expect(data.IsNext("el")).To(BeTrue())
			Expect(data.Next()).To(Equal("e"))
			Expect(data.Next()).To(Equal("l"))
		})
		It("check from end", func() {
			data.Data = "hello world"
			data.Index = 9

			Expect(data.PeekCurrent()).To(Equal("l"))
			Expect(data.IsNext("l")).To(BeTrue())
			Expect(data.IsNext("ld")).To(BeTrue())
			Expect(data.IsNext("ld a")).To(BeFalse())
		})
		It("insensitive check from beginning", func() {
			data.Data = "hello world"
			data.Index = 1

			Expect(data.PeekCurrent()).To(Equal("e"))
			Expect(data.IsNextCaseInsensitive("el")).To(BeTrue())
			Expect(data.IsNextCaseInsensitive("El")).To(BeTrue())
			Expect(data.IsNextCaseInsensitive("eL")).To(BeTrue())
			Expect(data.IsNextCaseInsensitive("EL")).To(BeTrue())
			Expect(data.IsNextCaseInsensitive("elLo")).To(BeTrue())
			Expect(data.IsNextCaseInsensitive("ll")).To(BeFalse())

			Expect(data.Next()).To(Equal("e"))
			Expect(data.Next()).To(Equal("l"))
		})
		It("insensitive check from end", func() {
			data.Data = "hello world"
			data.Index = 9

			Expect(data.PeekCurrent()).To(Equal("l"))
			Expect(data.IsNextCaseInsensitive("ld")).To(BeTrue())
			Expect(data.IsNextCaseInsensitive("lD")).To(BeTrue())
			Expect(data.IsNextCaseInsensitive("ld a")).To(BeFalse())
		})
	})
	Context("Integration", func() {

		It("Peek and Next Combination Tests", func() {
			data.Data = "hello world"

			Expect(data.PeekCurrent()).To(Equal("h"), "1st peek current")
			Expect(data.PeekNext()).To(Equal("e"), "1st peek next")

			Expect(data.Next()).To(Equal("h"), "next")

			Expect(data.PeekCurrent()).To(Equal("e"), "2nd peek current")
			Expect(data.PeekNext()).To(Equal("l"), "2nd peek next")
		})
	})

})

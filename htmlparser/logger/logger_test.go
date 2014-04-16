package logger_test

import (
	. "github.com/mkboudreau/reservations/htmlparser/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"reflect"
	"testing"
)

func TestLogging(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logging Test Suite")
}

var _ = Describe("Level Logger", func() {

	Context("State Testing", func() {
		var log *LevelLogger
		BeforeEach(func() {
			log = DefaultLogger()
		})

		It("Should turn on", func() {
			Expect(log.IsOn()).To(BeTrue())
			Expect(log.IsOff()).To(BeFalse())
		})
		It("Should turn off", func() {
			log.TurnOff()
			Expect(log.IsOff()).To(BeTrue())
			Expect(log.IsOn()).To(BeFalse())
		})

		It("Should set to error", func() {
			log.SetError()
			Expect(log.IsError()).To(BeTrue())
			Expect(log.IsWarn()).To(BeFalse())
			Expect(log.IsInfo()).To(BeFalse())
			Expect(log.IsDebug()).To(BeFalse())
			Expect(log.IsTrace()).To(BeFalse())

			Expect(log.IsOn()).To(BeTrue())
			Expect(log.IsOff()).To(BeFalse())
		})
		It("Should set to warn", func() {
			log.SetWarn()
			Expect(log.IsError()).To(BeTrue())
			Expect(log.IsWarn()).To(BeTrue())
			Expect(log.IsInfo()).To(BeFalse())
			Expect(log.IsDebug()).To(BeFalse())
			Expect(log.IsTrace()).To(BeFalse())

			Expect(log.IsOn()).To(BeTrue())
			Expect(log.IsOff()).To(BeFalse())
		})
		It("Should set to info", func() {
			log.SetInfo()
			Expect(log.IsError()).To(BeTrue())
			Expect(log.IsWarn()).To(BeTrue())
			Expect(log.IsInfo()).To(BeTrue())
			Expect(log.IsDebug()).To(BeFalse())
			Expect(log.IsTrace()).To(BeFalse())

			Expect(log.IsOn()).To(BeTrue())
			Expect(log.IsOff()).To(BeFalse())
		})
		It("Should set to debug", func() {
			log.SetDebug()
			Expect(log.IsError()).To(BeTrue())
			Expect(log.IsWarn()).To(BeTrue())
			Expect(log.IsInfo()).To(BeTrue())
			Expect(log.IsDebug()).To(BeTrue())
			Expect(log.IsTrace()).To(BeFalse())

			Expect(log.IsOn()).To(BeTrue())
			Expect(log.IsOff()).To(BeFalse())
		})
		It("Should set to trace", func() {
			log.SetTrace()
			Expect(log.IsError()).To(BeTrue())
			Expect(log.IsWarn()).To(BeTrue())
			Expect(log.IsInfo()).To(BeTrue())
			Expect(log.IsDebug()).To(BeTrue())
			Expect(log.IsTrace()).To(BeTrue())

			Expect(log.IsOn()).To(BeTrue())
			Expect(log.IsOff()).To(BeFalse())
		})
	})

	Context("State Testing", func() {
		var log *LevelLogger

		BeforeEach(func() {
			log = DefaultLogger()
		})

		It("Should print trace", func() {
			log.SetTrace()
			log.Error("ERROR : SHOULD SEE ME?", true, 5)
			log.Warn("WARN : SHOULD SEE ME?", true, 5)
			log.Info("INFO : SHOULD SEE ME?", true, 5)
			log.Debug("DEBUG : SHOULD SEE ME?", true, 5)
			log.Trace("TRACE : SHOULD SEE ME?", true, 5)
		})
		It("Should print debug", func() {
			log.SetDebug()
			log.Error("ERROR : SHOULD SEE ME?", true, 4)
			log.Warn("WARN : SHOULD SEE ME?", true, 4)
			log.Info("INFO : SHOULD SEE ME?", true, 4)
			log.Debug("DEBUG : SHOULD SEE ME?", true, 4)
			log.Trace("TRACE : SHOULD SEE ME?", false, 4)
		})
		It("Should print error", func() {
			log.SetError()
			log.Error("ERROR : SHOULD SEE ME?", true, 1)
			log.Warn("WARN : SHOULD SEE ME?", false, 1)
			log.Info("INFO : SHOULD SEE ME?", false, 1)
			log.Debug("DEBUG : SHOULD SEE ME?", false, 1)
			log.Trace("TRACE : SHOULD SEE ME?", false, 1)
		})
		It("Should print warn", func() {
			log.SetWarn()
			log.Error("ERROR : SHOULD SEE ME?", true, 2)
			log.Warn("WARN : SHOULD SEE ME?", true, 2)
			log.Info("INFO : SHOULD SEE ME?", false, 2)
			log.Debug("DEBUG : SHOULD SEE ME?", false, 2)
			log.Trace("TRACE : SHOULD SEE ME?", false, 2)
		})
		It("Should print info", func() {
			log.SetInfo()
			log.Error("ERROR : SHOULD SEE ME?", true, 3)
			log.Warn("WARN : SHOULD SEE ME?", true, 3)
			log.Info("INFO : SHOULD SEE ME?", true, 3)
			log.Debug("DEBUG : SHOULD SEE ME?", false, 3)
			log.Trace("TRACE : SHOULD SEE ME?", false, 3)
		})
	})

})

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

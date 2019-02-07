package cef_test

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	. "github.com/ArthurHlt/logrus-cef-formatter"
)

var _ = Describe("CefFormatter", func() {
	var logger = logrus.New()
	var buf = &bytes.Buffer{}
	var formatter = NewCEFFormatter("arthurhlt", "logrus-cef", "1.0")
	logger.Formatter = formatter
	logger.SetOutput(buf)
	BeforeEach(func() {
		formatter.DisableTimestamp = false
		buf.Reset()
	})
	Context("Format", func() {
		It("should format for CEF", func() {

			logger.Info("test")

			Expect(buf.String()).To(ContainSubstring("CEF:0|arthurhlt|logrus-cef|1.0|test|test|0|rt="))
		})
		It("should format for CEF with signature id when given", func() {
			logger.WithField(KeySignatureID, "my-signature").Info("test")

			Expect(buf.String()).To(ContainSubstring("CEF:0|arthurhlt|logrus-cef|1.0|my-signature|test|0|rt="))
		})
		When("Disable timestamp", func() {
			It("should format for CEF without timestamp", func() {
				formatter.DisableTimestamp = true

				logger.Info("test")

				Expect(buf.String()).To(Equal("CEF:0|arthurhlt|logrus-cef|1.0|test|test|0|\n"))
			})
		})
		When("With field", func() {
			It("should format for CEF with simple value field", func() {

				logger.
					WithField("string", "bar").
					WithField("number", 1).
					WithField("bool", true).
					Info("test")

				Expect(buf.String()).To(ContainSubstring("string=bar"))
				Expect(buf.String()).To(ContainSubstring("number=1"))
				Expect(buf.String()).To(ContainSubstring("bool=true"))
			})
			It("should format for CEF with complex value field", func() {

				logger.
					WithField("struct", struct {
						Foo string
						Bar string
					}{"toto", "titi"}).
					WithField("slice", []string{"foo", "bar"}).
					Info("test")

				Expect(buf.String()).To(ContainSubstring(`slice=["foo","bar"]`))
				Expect(buf.String()).To(ContainSubstring(`struct={"Foo":"toto","Bar":"titi"}`))
			})
		})
	})
})

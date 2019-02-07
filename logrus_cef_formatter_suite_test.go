package cef_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLogrusCefFormatter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LogrusCefFormatter Suite")
}

package sender_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRiemannCli(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sender Suite")
}

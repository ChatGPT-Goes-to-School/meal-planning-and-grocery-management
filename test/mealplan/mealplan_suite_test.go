package mealplan_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMealplan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mealplan Suite")
}

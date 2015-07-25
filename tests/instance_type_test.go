package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"github.com/fgrosse/goldi"
	"github.com/fgrosse/goldi/tests/testAPI"
)

var _ = Describe("InstanceType", func() {
	var resolver *goldi.ParameterResolver

	BeforeEach(func() {
		config := map[string]interface{}{}
		registry := goldi.NewTypeRegistry()
		resolver = goldi.NewParameterResolver(config, registry)
	})

	It("should panic if NewInstanceType is called with nil", func() {
		Expect(func() { goldi.NewInstanceType(nil) }).To(Panic())
	})

	Describe("Arguments()", func() {
		It("should return an empty list", func() {
			typeDef := goldi.NewInstanceType(testAPI.NewFoo())
			Expect(typeDef.Arguments()).To(BeEmpty())
		})
	})

	Describe("Generate", func() {
		It("should always return the given instance", func() {
			instance := testAPI.NewFoo()
			factory := goldi.NewInstanceType(instance)

			for i := 0; i < 3; i++ {
				generateResult := factory.Generate(resolver)
				Expect(generateResult == instance).To(BeTrue(),
					fmt.Sprintf("generateResult (%p) should point to the same instance as instance (%p)", generateResult, instance),
				)
			}
		})

		It("should panic if is called with nil", func() {
			factory := &goldi.InstanceType{}
			Expect(func() { factory.Generate(resolver) }).To(Panic())
		})
	})

	It("should implement the TypeFactory interface", func() {
		var factory goldi.TypeFactory
		factory = goldi.NewInstanceType("foo")
		// if this compiles the test passes (next expectation only to make compiler happy)
		Expect(factory).NotTo(BeNil())
	})
})
/*
Name: Roy Sanchez
Class: CS480
Boilerplate code

Article for getting started with bdd in go using Ginkgo
https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo

*/

package cart_test

import (
	"testing"
	
	. "github.com/onsi/ginkgo" // Allows for direc usage without package prefix
	. "github.com/onsi/gomega" // "ginkgo.Describe -> Describe"
	)
	
	
	/* Test Function */
	/*
		RegisterFailHandler(Fail) -> Creates coupling between Ginkgo and Gomega
		var _ = Describe("Shopping cart", func() {}) -> Evaluates the Describe block. 
		(Alternatively, you can also wrap the Describe block in an init() function)
	*/
	func TestCart(t *testing.T) {
	
		RegisterFailHandler(Fail) // Registers a handler, (a call to a function within the ginkgo package) Fail Function, with Gomega 
		RunSpecs(t, "Shopping Cart Suite") // Hands over the control to Ginkgo which runs the test suite.
		
	}
	
	var _ = Describe("Shopping cart", func() {
		Context("initially", func() {
			It("has 0 items", func() {})
			It("has 0 units", func() {})
			Specify("the total amount is 0.00", func() {})
		}) 
		
		
		Context("when a new item is added", func() {
			Context("the shopping cart", func() {
				It("has 1 more unique item than it had earlier", func() {})
				It("has 1 more unit than it had earlier", func() {})
				Specify("total amount increases by item price", func() {})
			})
		})
		
		Context("that has 0 unit of item A", func() {
			Context("removing item A", func() {
				It("should not change the number of items", func() {})
				It("should not change the number of units", func() {})
				It("should not change the amount", func() {})
			})
		})
		
		Context("that has 1 unit of item A", func() {
			Context("removing 1 unit item A", func() {
				It("should reduce the number of items by 1", func() {})
				It("should reduce the number of units by 1", func() {})
				It("should reduce the amount by item price", func() {})
			})
		})
		
		Context("that has 2 units of item A", func() {
			
			Context("removing 1 unit of item A", func() {
				It("should not reduce the number of items", func() {})
				It("should reduce the number of units by 1", func() {})
				It("should reduce the amount by the item price", func() {})
			})
			
			Context("removing 2 units of item A", func() {
				It("should reduce the number of items by 1", func() {})
				It("should reduce the number of units by 2", func() {})
				It("should reduce the amount by twice the item price", func() {})
			})
		
		})
		
	})

/*
Name: Roy Sanchez
Class: CS480
Program: cart.go
Building shopping cart implementation to satisfy the tests written
in cart_test.go
*/

package cart

// Cart represents the state of a buyer's shopping cart
type Cart struct {
	items map[string] Item
}

// Item represents any item available for sale
type Item struct {
	ID string
	Name string
	Price float64
	Qty int
}

/*
init(): 
	Instantiates items map Method will be called from all 
	other methods which need to access the items map directly.
*/

func (c *Cart) init() {
	if c.items == nil {
		c.items = map[string]Item{}
	}
}

/*
AddItem: 
	Checks if an item currently exists in cart.
	If true, increase quantity of item, else, add item to cart
*/

func (c *Cart) AddItem(i Item) {
	c.init()
	if existingItem, ok := c.items[i.ID]; ok {
		existingItem.Qty++
		c.items[i.ID] = existingItem
	} else {
		i.Qty = 1
		c.items[i.ID] = i
	}
}

// RemoveItem removes n number of items with give id from the cart
/*
RemoveItem:
	Reduces quantity of the item if the item exists in excess
	quantity from # passed in. Else removes item from cart.
*/

func (c *Cart) RemoveItem(id string, n int) {
	c.init()
	if existingItem, ok := c.items[id]; ok {
		if existingItem.Qty <= n {
			delete(c.items, id)
		} else {
			existingItem.Qty -= n
			c.items[id] = existingItem
		}
	}
}

// TotalUnits returns the total number of units across all items in the cart
/*
TotalAmount: 
	Goes through all items, calculates the total based on quantity and price of indiv 
	items and returns the amount
*/

func (c *Cart) TotalUnits() int {
	c.init()
	totalUnits := 0
	for _, i := range c.items {
		totalUnits += i.Qty
	}
	return totalUnits
}

// TotalUnitsItems returns the number of unique items in the cart
/*
TotalUniqueItems:
	Should return the total number of unique items.
*/
func (c *Cart) TotalUniqueItems() int {
	return len(c.items)
}

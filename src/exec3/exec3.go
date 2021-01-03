//Practice using data structures to build object oriented programs
package main

import "fmt"

func main() {
	var p Price =2595
	fmt.Printf("test get Price in euro : %q \n ",p.getPriceInEuro())
	RegisterItem(Prices, "banana", 345)
	fmt.Printf("Prices: %v\n", Prices)
	// Re-register item
	RegisterItem(Prices, "banana", 350)
	fmt.Printf("Prices: %v\n", Prices)

	c := new(Cart)
	c.AddItem("eggs")
	c.AddItem("banana")
	fmt.Printf("c.HasItem(%v) = %v\n", "bread", c.HasItem("bread"))

	c.AddItem("chocolat")
	c.AddItem("milk")
	c.Checkout()
}

// Price is the cost of something.
type Price int64

// getPriceInEuro is the string representation of a Price in euro
// Example: 2595 centimes => €25.95
func (p Price) getPriceInEuro() string {
	// TODO
	var pr = fmt.Sprint(p)
	var pr_rune = []rune(pr)
	return "€"+string(pr_rune[0:len(pr)-2])+"."+string(pr_rune[len(pr)-2:len(pr)])
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          519,
	"bread":         119,
	"apples":        595,
	"chips":         245,
	"milk":          150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
func RegisterItem(prices map[string]Price, item string, price Price) {
	// TODO
	var p = prices[item]
	if p == 0 {
		println("New item")
		prices[item]=price
	}
	if p != 0{
		println("Warning! item already in prices! the price will be changed")
		prices[item]=price
	}

}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	// TODO
	for i := 0; i < len(c.Items); i++ {
        if c.Items[i] == item {
            return true
        }
    }
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
func (c *Cart) AddItem(item string) {
	// TODO
	for key, _ := range Prices {
        if key == item {
            c.Items = append(c.Items, item)
        	c.TotalPrice += Prices[item]
        	fmt.Println("Item " + item + " added")
        }
    }
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	// TODO
	fmt.Println(c.Items)
    fmt.Println("Total : " + c.TotalPrice.getPriceInEuro())
    c = new(Cart)
}

package main

func main() {
	g1 := game{
		title: "Minecraft",
		price: 5,
	}
	g2 := game{
		title: "World of warcraft",
		price: 19,
	}
	g3 := game{
		title: "Elite Dangerous",
		price: 54,
	}
	b1 := book{
		title: "Candle in the tomb",
		price: 20,
	}
	b2 := book{
		title: "Barney and Friends",
		price: 10,
	}
	ca1 := computerAccessories{
		title: "Razer BT earpiece",
		price: 159,
	}
	ca2 := computerAccessories{
		title: "Razer keyboard",
		price: 110,
	}
	ca3 := computerAccessories{
		title: "Logitech Mouse",
		price: 80,
	}
	var store List
	store = append(store, &g1, &g2, &g3, &b1, &b2, &ca1, &ca2, &ca3)
	store.Print()
}

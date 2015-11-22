package items
import "fmt"

type item struct {
	Key string `json:"key"`
	ItemGroupKey string `json:"item_group_key"`
	BasicPrice int `json:"basic_price"`
	VipPrice int `json:"vip_price"`
	Effects map[string]int `json:"effects"`
}

type itemList []*item

var items itemList

func setupItems() {
	items = itemList {
		&item{
			Key: "item_1",
			ItemGroupKey: "item_group_1",
			BasicPrice: 10,
		},

		&item{
			Key: "item_2",
			ItemGroupKey: "item_group_1",
			BasicPrice: 20,
		},

		&item{
			Key: "item_3",
			ItemGroupKey: "item_group_1",
			BasicPrice: 30,
		},
	}
}

func All() itemList {
	if len(items) == 0 {
		setupItems()

		fmt.Println("Loaded Items!!!")
	}

	return items
}

func (items itemList) Select(key string) itemList{
	newItems := itemList{}

	for _, item := range items {
		if item.Key == key {
			newItems = append(newItems, item)
		}
	}

	return newItems
}

func AsJson() {

}



package items

type ItemGroup struct {
	key string
}

var ItemGroups []*ItemGroup

func SetupItemGroups() {
	ItemGroups = []*ItemGroup{
		&ItemGroup{key: "item_group_1"},
		&ItemGroup{key: "item_group_2"},
	}
}


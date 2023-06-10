package main

type PancakeHouseMenu struct {
	menuItems []*MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	var menuItems []*MenuItem

	ret := &PancakeHouseMenu{
		menuItems: menuItems,
	}

	ret.addItem("K&B 팬케이크 세트", "스크램블 에그와 토스트가 곁들여진 팬케이크", true, 2.99)
	ret.addItem("레귤러 팬케이크 세트", "달걀 프라이와 소시지가 곁들여진 팬케이크", false, 2.99)
	ret.addItem("블루베리 팬케이크", "신선한 블루베리와 블루베리 시럽으로 만든 팬케이크", true, 3.49)
	ret.addItem("와플", "취향에 따라 블루베리나 딸기를 얹을 수 있는 와플", true, 3.59)

	return ret
}

func (p *PancakeHouseMenu) addItem(name, description string, vegetarian bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarian, price)
	p.menuItems = append(p.menuItems, menuItem)
}

func (p *PancakeHouseMenu) MenuItems() []*MenuItem {
	return p.menuItems
}

func (p *PancakeHouseMenu) CreateIterator() Iterator[MenuItem] {
	return NewPancakeHouseMenuIterator(p.menuItems)
}

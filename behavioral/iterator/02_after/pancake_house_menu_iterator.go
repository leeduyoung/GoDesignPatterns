package main

type PancakeHouseMenuIterator struct {
	menuItems []*MenuItem
	position  int
}

func NewPancakeHouseMenuIterator(menuItems []*MenuItem) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{
		menuItems: menuItems,
		position:  0,
	}
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	if p.position >= len(p.menuItems)-1 {
		return false
	}

	return true
}

func (p *PancakeHouseMenuIterator) Next() *MenuItem {
	menuItem := p.menuItems[p.position]
	p.position = p.position + 1
	return menuItem
}

package main

type DinnerIterator struct {
	menuItems [4]*MenuItem
	position  int
}

func NewDinnerIterator(menuItems [4]*MenuItem) *DinnerIterator {
	return &DinnerIterator{
		menuItems: menuItems,
		position:  0,
	}
}

func (d *DinnerIterator) HasNext() bool {
	if d.position >= len(d.menuItems)-1 || d.menuItems[d.position] == nil {
		return false
	}

	return true
}

func (d *DinnerIterator) Next() *MenuItem {
	menuItem := d.menuItems[d.position]
	d.position = d.position + 1
	return menuItem
}

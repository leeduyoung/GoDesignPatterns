package main

import "fmt"

type DinnerMenu struct {
	maxItems      int
	menuItems     [4]*MenuItem
	numberOfItems int
}

func NewDinnerMenu() *DinnerMenu {
	var menuItems [4]*MenuItem

	ret := &DinnerMenu{
		maxItems:      6,
		menuItems:     menuItems,
		numberOfItems: 0,
	}

	ret.addItem("채식주의자용 BLT", "통밀 위에 콩고기 베이컨, 상추, 토마토를 얹은 메뉴", true, 2.99)
	ret.addItem("BLT", "통밀 위에 베이컨, 상추, 토마토를 얹은 메뉴", false, 2.99)
	ret.addItem("오늘의 스프", "감자 샐로드를 곁들인 오늘의 스프", false, 3.29)
	ret.addItem("핫도그", "사워크라우트, 갖은 양념, 양파, 치즈가 곁들여진 핫도그", false, 3.05)

	return ret
}

func (d *DinnerMenu) addItem(name, description string, vegetarian bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarian, price)
	if d.numberOfItems >= d.maxItems {
		fmt.Println("죄송합니다, 메뉴가 꽉 찼습니다. 더 이상 추가할 수 없습니다")
	} else {
		d.menuItems[d.numberOfItems] = menuItem
		d.numberOfItems = d.numberOfItems + 1
	}
}

func (d *DinnerMenu) MenuItems() [4]*MenuItem {
	return d.menuItems
}

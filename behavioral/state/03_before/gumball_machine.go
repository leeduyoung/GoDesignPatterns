package main

import "fmt"

const (
	SOLD_OUT    = 0
	NO_QUARTER  = 1
	HAS_QUARTER = 2
	SOLD        = 3
)

type gumballMachine struct {
	state int
	count int
}

type GumballMachine interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
}

func NewGumballMachine(count int) GumballMachine {
	gumballMachine := &gumballMachine{
		state: SOLD_OUT,
		count: count,
	}

	if count > 0 {
		gumballMachine.state = NO_QUARTER
	}

	return gumballMachine
}

func (gm *gumballMachine) InsertQuarter() {
	if gm.state == HAS_QUARTER {
		fmt.Println("동전은 한개만 넣어주세요.")
	} else if gm.state == NO_QUARTER {
		gm.state = HAS_QUARTER
		fmt.Println("동전을 넣으셨습니다.")
	} else if gm.state == SOLD_OUT {
		fmt.Println("매진 되었습니다. 다음에 이용해 주세요.")
	} else if gm.state == SOLD {
		fmt.Println("캡슐을 내보내고 있습니다.")
	}
}

func (gm *gumballMachine) EjectQuarter() {
	if gm.state == HAS_QUARTER {
		gm.state = NO_QUARTER
		fmt.Println("동전이 반환되었습니다.")
	} else if gm.state == NO_QUARTER {
		fmt.Println("반환할 동전이 없습니다.")
	} else if gm.state == SOLD_OUT {
		fmt.Println("매진 되어 동전을 넣을 수 없습니다.")
	} else if gm.state == SOLD {
		fmt.Println("이미 캡슐이 반환되었습니다.")
	}
}

func (gm *gumballMachine) TurnCrank() {
	if gm.state == HAS_QUARTER {
		fmt.Println("손잡이를 돌리셨습니다.")
		gm.state = SOLD
		gm.dispense()
	} else if gm.state == NO_QUARTER {
		fmt.Println("먼저 동전을 넣어주세요.")
	} else if gm.state == SOLD_OUT {
		fmt.Println("매진입니다.")
	} else if gm.state == SOLD {
		fmt.Println("손잡이는 한번만 돌려주세요.")
	}
}

func (gm *gumballMachine) dispense() {
	if gm.state == HAS_QUARTER {
		fmt.Println("오류")
	} else if gm.state == NO_QUARTER {
		fmt.Println("오류")
	} else if gm.state == SOLD_OUT {
		fmt.Println("오류")
	} else if gm.state == SOLD {
		fmt.Println("캡슐을 내보냅니다.")

		gm.count = gm.count - 1
		if gm.count == 0 {
			fmt.Println("캡슐이 매진되었습니다.")
			gm.state = SOLD_OUT
		} else {
			gm.state = NO_QUARTER
		}
	}
}

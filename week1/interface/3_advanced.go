package main

import "fmt"

// -------------

type ApplePay struct {
	Money		int
	AppleID		string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Не хватает денег на ApplePay\n")
	}

	a.Money -= amount
	return nil
}

// ----------------

type Payer interface {
	Pay(int) error
}

type Ringer interface {
	Ring(string) error
}

// Композитный интерфейс - интерфейс, сочетающий в себе два интерфейса
// В данном случае, чтобы соответствовать инетрфейсу NFCPhone, нужно чтобы структура умела и в Pay и в Ring
type NFCPhone interface {
	Payer
	Ringer
}

// ----------------

// Пустой интерфейс приводится для того чтобы в эту функцию можно было передать любую структуру
// но работать она могла нормально только с интерфейсом Payer
func Buy(in interface{}) {
	var p Payer
	var ok bool

	// Проверяем, можем ли привести пустой интерфейс к Payer
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не является платежным средством\n\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}


func main()  {
	myMoney := &ApplePay{Money: 100}
	Buy(myMoney)

	t1 := []int{1,2,3}
	Buy(t1)

	Buy("Hello")
}


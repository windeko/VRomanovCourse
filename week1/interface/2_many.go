package main

import "fmt"

// --------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("В кошельке нехватает денег\n")
	}

	w.Cash -= amount
	return nil
}

// -------------

type Card struct {
	Balance			int
	ValidUntil		string
	CardHolder		string
	cvv				string
	Number			string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("На карте нехватает денег\n")
	}

	c.Balance -= amount
	return nil
}

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

// ----------------

func Buy(p Payer) {
	// Чтобы узнать тип переданной структуры и иметь возможность обратиться к полям оригинальной структуры
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличными")
	case *Card:
		// Эта конструкция преобразовывает переменнут типа Payer в Card
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("Не удалось преобразовать к типу *Card")
		}
		fmt.Println("Вставляйте карту,", plasticCard.CardHolder)
	default:
		fmt.Println("Что-то новенькое")
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}


func main()  {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{
		Balance: 200,
		CardHolder: "vtraigel",
	}
	Buy(myMoney)

	myMoney = &ApplePay{}
	Buy(myMoney)
}
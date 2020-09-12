package main

import "fmt"

type payer interface {
	Pay(int) error
}

type wallet struct {
	Cash int
}

func (w *wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("В кошельке нехватает денег")
	}
	
	w.Cash -= amount
	return nil
}

func Buy(p payer)  {
	err := p.Pay(30)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T \n\n", p)
}

func main()  {
	myWallet := wallet{Cash: 100}
	Buy(&myWallet)
	Buy(&myWallet)
	Buy(&myWallet)
	Buy(&myWallet)
}
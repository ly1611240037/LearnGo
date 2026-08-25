package main

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(amount float64) error
	GetBalance() float64
}

type Alipay struct {
	balance float64
}

func (a *Alipay) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("支付金额必须大于 0")
	}
	if amount > a.balance {
		return errors.New("支付宝余额不足")
	}
	a.balance -= amount
	return nil
}

func (a *Alipay) GetBalance() float64 {
	return a.balance
}

type WeChatPay struct {
	balance float64
}

func (w *WeChatPay) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("支付金额必须大于 0")
	}
	if amount > w.balance {
		return errors.New("微信余额不足")
	}
	w.balance -= amount
	return nil
}

func (w *WeChatPay) GetBalance() float64 {
	return w.balance
}

func main() {
	payments := []Payment{
		&Alipay{balance: 10},
		&WeChatPay{balance: 200},
	}

	for _, payment := range payments {
		if err := payment.Pay(50); err != nil {
			fmt.Println("支付失败：", err)
			continue
		}
		fmt.Printf("支付成功，剩余余额：%.2f\n", payment.GetBalance())
	}
}

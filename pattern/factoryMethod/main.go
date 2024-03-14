package main

import (
	"errors"
	"fmt"
	"log"
)

type IBankAccout interface {
	PutMoney(float32)
	WriteOffMoney(float32)
	GetBalance() float32
}

func main() {
	ya, err := NewBankAccount("commercial", "Yandex")
	if err != nil {
		log.Fatal(err)
	}

	v, err := NewBankAccount("personal", "Vasya")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ya)
	fmt.Println(v)
}

type BankAccout struct {
	balance float32
}

func NewBankAccount(accountType, ownerName string) (IBankAccout, error) {
	if accountType == "commercial" {
		return NewBusynessAccount(ownerName), nil
	} else if accountType == "personal" {
		return NewPersonalAccount(ownerName), nil
	} else {
		return nil, errors.New("provided type is incorrect")
	}
}

func (b *BankAccout) PutMoney(increment float32) {
	b.balance += increment
}

func (b *BankAccout) WriteOffMoney(decrement float32) {
	b.balance -= decrement
}

func (b *BankAccout) GetBalance() float32 {
	return b.balance
}

type PersonalBankAccout struct {
	OwnerName string
	BankAccout
}

func NewPersonalAccount(name string) *PersonalBankAccout {
	return &PersonalBankAccout{OwnerName: name}
}

func (p PersonalBankAccout) TakeCredit() {
	fmt.Println("credit was taken")
}

type BusynessBankAccount struct {
	OrganisationName string
	BankAccout
}

func NewBusynessAccount(name string) *BusynessBankAccount {
	return &BusynessBankAccount{OrganisationName: name}
}

func (b BusynessBankAccount) MakeDeposit() {
	fmt.Println("depo was made")
}

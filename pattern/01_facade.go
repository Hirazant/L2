package main

import (
	"fmt"
	"log"
)

/*
Паттерн «Фасад» предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы.
Фасад определяет интерфейс более высокого уровня, который упрощает использование подсистемы.
*/

/*
Фасад — это структурный паттерн проектирования,
который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
*/

/*
«Фасад» — некоторый объект,
аккумулирующий в себе высокоуровневый набор операций для работы с некоторой сложной подсистемой.

Клиент, при этом, не лишается более низкоуровневого доступа к классам подсистемы.
Фасад упрощает выполнение некоторых операций с подсистемой, но не навязывает клиенту свое использование.
*/

/*Плюсами такого подхода является:
- Упрощение работы клиента с подсистемой - меньше кода, меньше ошибок, быстрее разработка.
- Уменьшении зависимости от подсистемы - проще внести изменения, проще тестировать.
- Упрощение внешней документации - упрощение работы с подсистемой для клиента - проще клиентская документация
Минусами подхода является:
- Требуется дополнительная реализация необходимых интерфейсов - дополнительная разработка.
- Нужно хорошо продумать реализуемый набор интерфейсов для клиента, чтобы вся функциональность, ему
необходимая, была у него доступна (при доработках подсистемы нужно поддерживать и фасад).
- Фасад может стать божествееным объектом
*/

// Класс кошелька
// -----------------------------------------------------------------------------------------------------------
type wallet struct {
	balance float64
}

// Конструктор кошелька
func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

// Метод добавления денег в кошелек
func (w *wallet) addMoney(money float64) {
	w.balance += money
	fmt.Println("money add!")
}

// -----------------------------------------------------------------------------------------------------------

// Класс Юзера
// -----------------------------------------------------------------------------------------------------------
type user struct {
	name string
}

// Конструктор Юзера
func newUser(name string) *user {
	return &user{
		name: name,
	}
}

// Метод проверки Юзера
func (user *user) checkUser(userName string) error {
	if user.name != userName {
		return fmt.Errorf("user name is incorrect")
	}
	fmt.Println("User verified")
	return nil
}

// -----------------------------------------------------------------------------------------------------------

// Класс фасада
// -----------------------------------------------------------------------------------------------------------
type walletFacade struct {
	wallet *wallet
	user   *user
}

//Конструктор фасада
func newWalletFacade(user string) *walletFacade {
	return &walletFacade{
		wallet: newWallet(),
		user:   newUser(user),
	}
}

// Метод фасада по проверке юзера и добавлении ему денег
func (facade *walletFacade) addMoney(user string, money float64) error {
	err := facade.user.checkUser(user)
	if err != nil {
		return err
	}
	facade.wallet.addMoney(money)
	fmt.Println("Facade has done its job")
	return nil
}

// -----------------------------------------------------------------------------------------------------------

func main() {
	facade := newWalletFacade("Super")
	err := facade.addMoney("Super", 100)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("New balance: ", facade.wallet.balance)
}

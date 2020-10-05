package types

//Money type for money
type Money int

//PaymentCategory for example(car, purchases)
type PaymentCategory string

//PaymentStatus for status of out payment
type PaymentStatus string

//Currency for currency
type Currency string

//The staus of paymets can be
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

//Payment представляет информацию о платеже
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

//Phone type for phone number
type Phone string

//Account discribes the structure of account
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

//Favorite structure
type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}

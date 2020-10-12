package wallet

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/JovidYnwa/wallet/pkg/types"
	"github.com/google/uuid"
)

//Service The structure of our service
type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
	favorites     []*types.Favorite
}

//Error type
type Error string

func (e Error) Error() string {
	return string(e)
}

//ErrPhoneRegistred custom error
var ErrPhoneRegistred = errors.New("phone already registrated")

//ErrAmountMustBePositive custom error amount should be greater then zero
var ErrAmountMustBePositive = errors.New("amount must be greater then zero")

//ErrAccountNotFound custom error if account was not found
var ErrAccountNotFound = errors.New("account not found")

//ErrPaymentNotFound for case if payment is not found
var ErrPaymentNotFound = errors.New("payment not found")

//ErrFavoriteNotFound favorite not found
var ErrFavoriteNotFound = errors.New("favorite not found")

//ErrFileNotFound for finding writing file
var ErrFileNotFound = errors.New("There is no such a file")

//RegisterAccount Fuction for registration of users
//func RegisterAccount(service *Service, phone types.Phone) {
//	for _, account := range service.accounts {
//		if account.Phone == phone {
//			return
//		}
//	}
//	service.nextAccountID++
//	service.accounts = append(service.accounts, &types.Account{
//		ID:      service.nextAccountID,
//		Phone:   phone,
//		Balance: 0,
//	})
//}

//RegisterAccount Fuction for registration of users
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistred
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)

	return account, nil
}

//Deposit method for depositing the account
func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return ErrPhoneRegistred
	}

	account.Balance += amount
	return nil

}

//Pay functio for payment
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}

	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: accountID,
		Amount:    amount,
		Category:  category,
		Status:    types.PaymentStatusInProgress,
	}

	s.payments = append(s.payments, payment)
	return payment, nil
}

//FindAccountByID first task
func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account

	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}
	return account, nil
}

//FindPaymentByID the following just looks for paymet by payId
func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error) {
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound

}

//Reject rejecting a payment
func (s *Service) Reject(paymentID string) error {

	payment, err := s.FindPaymentByID(paymentID)

	if err != nil {
		return err
	}

	account, err := s.FindAccountByID(payment.AccountID)

	if err != nil {
		return err
	}

	account.Balance += payment.Amount
	payment.Status = types.PaymentStatusFail

	return nil
}

//Repeat the following repeat payme by paymetID
func (s *Service) Repeat(paymentID string) (*types.Payment, error) {

	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}
	paymentNew, err := s.Pay(payment.AccountID, payment.Amount, payment.Category)
	if err != nil {
		return nil, err
	}
	return paymentNew, nil
}

//FavoritePayment creates Favorite payments
func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error) {

	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	favoriteID := uuid.New().String()
	favorite := &types.Favorite{
		ID:        favoriteID,
		AccountID: payment.AccountID,
		Name:      name,
		Amount:    payment.Amount,
		Category:  payment.Category,
	}

	s.favorites = append(s.favorites, favorite)

	return favorite, nil
}

//PayFromFavorite the following just makes paymet from Favorite
func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error) {

	var favorite *types.Favorite
	for _, v := range s.favorites {
		if v.ID == favoriteID {
			favorite = v
			break
		}
	}
	if favorite == nil {
		return nil, ErrFavoriteNotFound
	}

	payment, err := s.Pay(favorite.AccountID, favorite.Amount, favorite.Category)

	if err != nil {
		return nil, err
	}
	return payment, nil
}

// ExportToFile the following just writes down
func (s *Service) ExportToFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()
	addingPart := ""

	for _, acc := range s.accounts {
		ID := strconv.Itoa(int(acc.ID)) + ";"
		phone := string(acc.Phone) + ";"
		balance := strconv.Itoa(int(acc.Balance))

		addingPart = addingPart + ID + phone + balance + "|"

	}

	_, err = file.Write([]byte(addingPart))
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}
	return nil
}

package wallet

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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

//Service1 The structure of our service
type Service1 struct {
	nextAccountID int64
	Accounts      []*types.Account
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

// FindFavoriteByID desctiption
func (s *Service) FindFavoriteByID(favoriteID string) (*types.Favorite, error) {
	for _, favorite := range s.favorites {
		if favoriteID == favorite.ID {
			return favorite, nil
		}
	}
	return nil, ErrFavoriteNotFound
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

//ImportFromFile the following
func (s *Service) ImportFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()

	content := make([]byte, 0)
	buf := make([]byte, 4)
	for {
		read, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
			return ErrFileNotFound
		}
		content = append(content, buf[:read]...)
	}
	data := string(content)

	accounts := strings.Split(string(data), "|")
	accounts = accounts[:len(accounts)-1]
	for _, account := range accounts {
		items := strings.Split(account, ";")
		ID, err := strconv.Atoi(items[0])
		if err != nil {
			return err
		}
		balance, err := strconv.Atoi(items[2])
		if err != nil {
			return err
		}
		newAccount := &types.Account{
			ID:      int64(ID),
			Phone:   types.Phone(items[1]),
			Balance: types.Money(balance),
		}
		s.accounts = append(s.accounts, newAccount)
	}
	return nil
}

//Export for first home work of lecture 17
func (s *Service) Export(dir string) error {
	//Codintion in case if accounts array is not empty
	if s.accounts != nil {
		accountFile := "/accounts.dump"
		file, err := os.Create(dir + accountFile)
		for _, account := range s.accounts {
			txtitemue := []byte(strconv.FormatInt(int64(account.ID), 10) + string(";") + string(account.Phone) + string(";") + strconv.FormatInt(int64(account.Balance), 10) + string(";") + string('\n'))
			_, err = file.Write(txtitemue)
			if err != nil {
				return err
			}
		}
		log.Print("Accounts data exported")
	} else {
		log.Print("Accounts are empy nothing to be exported")
	}
	//Codintion in case if paymets array is not empty
	if s.payments != nil {
		paymetsFile := "/payments.dump"
		file, err := os.Create(dir + paymetsFile)
		for _, payment := range s.payments {
			txtitemue := []byte(string(payment.ID) + string(";") + strconv.FormatInt(int64(payment.AccountID), 10) + string(";") + strconv.FormatInt(int64(payment.Amount), 10) + string(";") + string(payment.Category) + string(";") + string(payment.Status) + string(";") + string('\n'))
			_, err = file.Write(txtitemue)
			if err != nil {
				return err
			}
		}
		log.Print("Payments data exported")
	} else {
		log.Print("Payments are empy nothing to be exported")
	}

	//Codintion in case if favorites array is not empty
	if s.favorites != nil {
		favoriteFile := "/favorites.dump"
		favFile, err := os.Create(dir + favoriteFile)
		if err != nil {
			log.Print(err)
			return err
		}

		for _, fav := range s.favorites {
			text := []byte(fav.ID + ";" + strconv.FormatInt(int64(fav.AccountID), 10) + ";" + fav.Name + ";" + strconv.FormatInt(int64(fav.Amount), 10) + ";" + string(fav.Category) + ";" + string('\n'))
			_, err := favFile.Write(text)
			if err != nil {
				log.Print(err)
				return err
			}
		}
		log.Print("Favorites data exported")
	} else {
		log.Print("Favorites are empy nothing to be exported")
	}

	return nil
}

// Import for first home work of lecture 17
func (s *Service) Import(dir string) error {
	//For accounts
	accountFile := "/accounts.dump"
	src, err := os.Open(dir + accountFile)
	if err != nil {
		log.Print("There is no %w file", accountFile)
	} else {
		defer func() {
			if cerr := src.Close(); cerr != nil {
				log.Print(cerr)
			}
		}()

		reader := bufio.NewReader(src)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				log.Print(line)
				break
			}
			if err != nil {
				log.Print(err)
				return err
			}

			item := strings.Split(line, ";")

			id, err := strconv.ParseInt(item[0], 10, 64)
			if err != nil {
				log.Print(err)
				return err
			}

			phone := item[1]

			balance, err := strconv.ParseInt(item[2], 10, 64)
			if err != nil {
				log.Print(err)
				return err
			}

			findAccount, _ := s.FindAccountByID(id)
			if findAccount != nil {
				findAccount.Phone = types.Phone(phone)
				findAccount.Balance = types.Money(balance)
			} else {
				s.nextAccountID = id
				newAcc := &types.Account{
					ID:      s.nextAccountID,
					Phone:   types.Phone(phone),
					Balance: types.Money(balance),
				}

				s.accounts = append(s.accounts, newAcc)
			}
		}
		log.Print("Imported")

	}

	// For Payments
	paymentsFile := "/payments.dump"
	paySrc, err := os.Open(dir + paymentsFile)
	if err != nil {
		log.Print("There is no %w file", paymentsFile)
	} else {
		defer func() {
			if cerr := paySrc.Close(); cerr != nil {
				log.Print(cerr)
			}
		}()

		payReader := bufio.NewReader(paySrc)
		for {
			payLine, err := payReader.ReadString('\n')
			if err == io.EOF {
				log.Print(payLine)
				break
			}
			if err != nil {
				log.Print(err)
				return err
			}

			item := strings.Split(payLine, ";")

			id := string(item[0])
			accID, err := strconv.ParseInt(item[1], 10, 64)
			if err != nil {
				log.Print(err)
				return err
			}

			amount, err := strconv.ParseInt(item[2], 10, 64)
			if err != nil {
				log.Print(err)
				return err
			}

			category := item[3]

			status := item[4]

			findPay, _ := s.FindPaymentByID(id)
			if findPay != nil {
				findPay.AccountID = accID
				findPay.Amount = types.Money(amount)
				findPay.Category = types.PaymentCategory(category)
				findPay.Status = types.PaymentStatus(status)
			} else {
				newPay := &types.Payment{
					ID:        id,
					AccountID: accID,
					Amount:    types.Money(amount),
					Category:  types.PaymentCategory(category),
					Status:    types.PaymentStatus(status),
				}

				s.payments = append(s.payments, newPay)
			}
		}
		log.Print("Imported")
	}

	//For favorites
	favoritesFile := "/payments.dump"
	favFile, err := os.Open(dir + favoritesFile)
	if err != nil {
		log.Print("There is no %w file", favoritesFile)
	} else {
		reader := bufio.NewReader(favFile)
		for {
			favLine, err := reader.ReadString('\n')
			if err == io.EOF {
				log.Print(favLine)
				break
			}
			if err != nil {
				log.Print(err)
				return err
			}

			item := strings.Split(favLine, ";")

			id := item[0]
			accID, err := strconv.ParseInt(item[1], 10, 64)
			if err != nil {
				log.Print(err)
				return err
			}
			name := item[2]
			amount, err := strconv.ParseInt(item[3], 10, 64)
			if err != nil {
				log.Print(err)
				return err
			}
			category := item[4]

			findFav, _ := s.FindFavoriteByID(id)
			if findFav != nil {
				findFav.AccountID = accID
				findFav.Amount = types.Money(amount)
				findFav.Name = name
				findFav.Category = types.PaymentCategory(category)
			} else {
				newFav := &types.Favorite{
					ID:        id,
					AccountID: accID,
					Name:      name,
					Amount:    types.Money(amount),
					Category:  types.PaymentCategory(category),
				}
				s.favorites = append(s.favorites, newFav)
			}
		}
		log.Print("Imported")
	}

	return nil
}

//ExperMy exp
func (s *Service1) ExperMy() types.Phone {
	inst1 := Service1{Accounts: []*types.Account{{Phone: "9010001000"}}}
	return inst1.Accounts[0].Phone

}

package wallet

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/JovidYnwa/wallet/pkg/types"
)

type testService struct {
	*Service
}
type testAccount struct {
	phone    types.Phone
	balance  types.Money
	payments []struct {
		amount   types.Money
		category types.PaymentCategory
	}
}

var someAccountForTest = testAccount{
	phone:   "+992888888888",
	balance: 1000,
	payments: []struct {
		amount   types.Money
		category types.PaymentCategory
	}{
		{amount: 1000, category: "auto"},
	},
}

func newTestService() *testService {
	return &testService{Service: &Service{}}
}

func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error) {
	account, err := s.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}

	err = s.Deposit(account.ID, data.balance)
	if err != nil {
		return nil, nil, fmt.Errorf("can't deposit account, error = %v", err)
	}

	payments := make([]*types.Payment, len(data.payments))

	for i, payment := range data.payments {
		payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
		if err != nil {
			return nil, nil, fmt.Errorf("can't make payment, error = %v", err)
		}
	}

	return account, payments, nil
}

func TestService_FindAccountByID_found(t *testing.T) {
	s := newTestService()

	account, _, err := s.addAccount(someAccountForTest)
	if err != nil {
		t.Error(err)
		return
	}

	got, err := s.FindAccountByID(account.ID)
	if err != nil {
		t.Error("FindAccountByID(): error", err)
		return
	}

	if !reflect.DeepEqual(account, got) {
		t.Error("Returned inccorect account  by FindAccountByID", err)
		return
	}
}

func TestService_FindAccountByID_success_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992907013487")

	account, err := svc.FindAccountByID(1)

	if err != nil {
		t.Errorf("method returned not nil error, account => %v", account)
	}

}

func TestService_FindAccountByID_notFound_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992907013487")

	account, err := svc.FindAccountByID(2)

	if err == nil {
		t.Errorf("method returned nil error, account => %v", account)
	}
}

func TestService_Reject_success_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992907013487")
	account, err := svc.FindAccountByID(1)

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, error => %v", err)
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 1000, "auto")

	if err != nil {
		t.Errorf("method Pay returned not nil error, error => %v", err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("method FindPaymentByID returned not nil error, error => %v", err)
	}

	err = svc.Reject(pay.ID)

	if err != nil {
		t.Errorf("method Reject returned not nil error, error => %v", err)
	}
}

func TestService_Reject_fail_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992000000001")
	account, err := svc.FindAccountByID(1)

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 1000, "auto")

	if err != nil {
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	err = svc.Reject(pay.ID + "uu")

	if err == nil {
		t.Errorf("method Reject returned not nil error, pay => %v", pay)
	}
}

func TestService_Repeat_success_user(t *testing.T) {
	var svc Service

	account, err := svc.RegisterAccount("+992907013487")

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 10_00, "auto")

	if err != nil {
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	paymentNew, err := svc.Repeat(pay.ID)

	if err != nil {
		t.Errorf("method Repat returned not nil error, paymentNew => %v", paymentNew)
	}
}

func TestService_Favorite_success_user(t *testing.T) {
	var svc Service

	account, err := svc.RegisterAccount("+992907013487")

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 10_00, "auto")

	if err != nil {
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	favorite, err := svc.FavoritePayment(payment.ID, "Favorite")

	if err != nil {
		t.Errorf("method FavoritePayment returned not nil error, favorite => %v", favorite)
	}

	paymentFavorite, err := svc.PayFromFavorite(favorite.ID)
	if err != nil {
		t.Errorf("method PayFromFavorite returned not nil error, paymentFavorite => %v", paymentFavorite)
	}
}

func TestService_SumPayments_success(t *testing.T) {
	want := types.Money(10_00)

	var svc Service

	account, err := svc.RegisterAccount("+992907013487")

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	_, err = svc.Pay(account.ID, 10_00, "auto")

	if err != nil {
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	got := svc.SumPayments(1)

	if want != got {
		t.Errorf("SumPayments(): want: %v got: %v", want, got)
		return
	}
}

func BenchmarkSumPayments(b *testing.B) {
	want := types.Money(10_00)

	var svc Service

	account, _ := svc.RegisterAccount("+992907013487")
	svc.Deposit(account.ID, 100_00)
	svc.Pay(account.ID, 10_00, "auto")

	for i := 0; i < b.N; i++ {
		result := svc.SumPayments(1)
		if result != want {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
	}
}

func TestService_Export_success(t *testing.T) {
	var svc Service

	account, _ := svc.RegisterAccount("+992907013487")
	svc.Deposit(account.ID, 100_00)
	svc.Pay(account.ID, 10_00, "auto")

	err := svc.Export("../../data")

	if err != nil {
		t.Errorf("Export(): error=%v", err)
		return
	}
}

func TestService_Import_success(t *testing.T) {
	var svc Service

	account, _ := svc.RegisterAccount("+992907013487")
	svc.Deposit(account.ID, 100_00)
	svc.Pay(account.ID, 10_00, "auto")

	err := svc.Import("../../data")

	if err != nil {
		t.Errorf("Import(): error=%v", err)
		return
	}
}

func TestService_ExportToFile_success(t *testing.T) {
	var svc Service

	account, _ := svc.RegisterAccount("+992907013487")
	svc.Deposit(account.ID, 100_00)
	svc.Pay(account.ID, 10_00, "auto")

	err := svc.ExportToFile("data")

	if err != nil {
		t.Error("error")
		return
	}
}

func TestService_ImportFromFile_success(t *testing.T) {
	var svc Service

	account, _ := svc.RegisterAccount("+992907013487")
	svc.Deposit(account.ID, 100_00)
	svc.Pay(account.ID, 10_00, "auto")

	err := svc.ImportFromFile("data")

	if err != nil {
		t.Error("error")
		return
	}
}

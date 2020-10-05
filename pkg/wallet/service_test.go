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

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

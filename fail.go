package functional_options

import "github.com/cloudwan/gohan/db/transaction"

// DB OMIT
type DB interface {
	BeginTx(options ...Option) (Transaction, error)
}
// DB# OMIT

// Want OMIT
mockDB.EXPECT().BeginTx(gomock.Eq(IsolationLevel(Serializable))).
	Return(mockTx, nil)
// Want# OMIT

// Have OMIT
mockDB.EXPECT().BeginTx(gomock.Any()).
	DoAndReturn(func(opt ...Option) (Transaction, error) {
		Expect(len(opt)).To(Equal(1))
		params := &TxParams{}
		opt[0](params)
		Expect(params.IsolationLevel).To(Equal(Serializable))
		return mockTx, nil
	})
// Have# OMIT
MyRpc interface {
	Balance(userName string) (count uint64, error string)
	Deposit(userName string, amount uint64)(deposit_amt uint64, error string)
	Withdraw(userName string, amount uint64)(withdraw_amt uint64, error string)
	Transfer(userName string, targetName string,amount uint64)(transfer_amt uint64, target string,error string)
}

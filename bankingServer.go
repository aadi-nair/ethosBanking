package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/myRpc"
	"log"
)

// var myRpc_increment_counter uint64 = 0
// var account_balance uint64 = 0
var Bank map[string]uint64
const BANK_ACCESS_ERROR = "BankAccessError"
const ERROR_NONE = "ErrorNone"
const LOW_BALANCE_ERROR="LowBalanceError"




func init() {
	myRpc.SetupMyRpcBalance(balance)
	myRpc.SetupMyRpcDeposit(deposit)
	myRpc.SetupMyRpcWithdraw(withdraw)
	myRpc.SetupMyRpcTransfer(transfer)

}

func balance(userName string) (myRpc.MyRpcProcedure) {
	// log.Println("Service: %s: called balance",userName)
	if account_balance,ok := Bank[userName]; ok {
		log.Println("Service: %s called balance",userName)
		return &myRpc.MyRpcBalanceReply{account_balance,ERROR_NONE}
	}
	log.Println("Service: %s called balance but failed",userName)
	return &myRpc.MyRpcBalanceReply{0,BANK_ACCESS_ERROR}
}
func deposit(userName string, amount uint64)(myRpc.MyRpcProcedure){

	if _,ok := Bank[userName]; ok {
		log.Println("Service: %s called deposit",userName)
		Bank[userName]+=amount
		return &myRpc.MyRpcDepositReply{amount,ERROR_NONE}
	}
	log.Println("Service: %s called deposit but failed",userName)
	return &myRpc.MyRpcDepositReply{0,BANK_ACCESS_ERROR}
	// account_balance += amount
}

func withdraw(userName string, amount uint64)(myRpc.MyRpcProcedure){

	if account_balance,ok := Bank[userName]; ok {
		 
		if amount > account_balance{
			log.Println("Service: %s called withdraw - insufficient funds",userName)
			return &myRpc.MyRpcWithdrawReply{amount,LOW_BALANCE_ERROR}
		}
		log.Println("Service: %s called withdraw",userName)
		Bank[userName]-=amount
		return &myRpc.MyRpcWithdrawReply{amount,ERROR_NONE}
	}
	log.Println("Service: %s called withdraw but failed",userName)
	return &myRpc.MyRpcWithdrawReply{0,BANK_ACCESS_ERROR}


}

func transfer(userName string, targetName string,amount uint64)(myRpc.MyRpcProcedure){
	if account_balance,ok := Bank[userName]; ok {
		if amount > account_balance{
			log.Println("Service: %s called transfer to %s- insufficient funds",userName, targetName)
			return &myRpc.MyRpcTransferReply{amount,targetName,LOW_BALANCE_ERROR}
		}
		if _,ok := Bank[targetName]; ok {
			log.Println("Service: %s called transfer to %s",userName, targetName)
			Bank[userName]-=amount
			Bank[targetName]+=amount
			return &myRpc.MyRpcTransferReply{amount,targetName,ERROR_NONE}
		}
		log.Println("Service: %s called transfer to %s but failed",userName, targetName)
		return &myRpc.MyRpcTransferReply{0,targetName,BANK_ACCESS_ERROR}

	}

	log.Println("Service: %s called transfer to %s but failed",userName, targetName)
	return &myRpc.MyRpcTransferReply{0,targetName,BANK_ACCESS_ERROR}

}


func main () {

	altEthos.LogToDirectory("test/bankingServer")
	Bank = make(map[string]uint64)
	
	Bank["me"] = 0
	Bank["nobody"] = 0
	Bank["pat"] = 0
	Bank["mike"] = 0



	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}
}

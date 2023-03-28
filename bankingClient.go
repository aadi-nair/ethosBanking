package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/myRpc"
	"ethos/defined"
	"log"
)
var userName string = ""
const BANK_ACCESS_ERROR = "BankAccessError"
const ERROR_NONE = "ErrorNone"
const LOW_BALANCE_ERROR="LowBalanceError"

func init() {

	myRpc.SetupMyRpcBalanceReply(balanceReply)
	myRpc.SetupMyRpcDepositReply(depositReply)
	myRpc.SetupMyRpcWithdrawReply(withdrawReply)
	myRpc.SetupMyRpcTransferReply(transferReply)


}

func balanceReply(count uint64, errorStr string) (myRpc.MyRpcProcedure) {

	switch errorStr{
		case BANK_ACCESS_ERROR:
			log.Printf("%s: Received Balance Reply - Error accessing bank",userName)
		default:
			log.Printf("%s: Received Balance Reply - Account Balance: $ %v\n",userName, count)
	}

	return nil
}
func depositReply(deposit_amt uint64,errorStr string) (myRpc.MyRpcProcedure) {
	switch errorStr{
		case BANK_ACCESS_ERROR:
			log.Printf("%s: Received Deposit Reply - Error accessing bank",userName)
		default:
			log.Printf("%s: Received Deposit Reply - Amount Deposited $ %v\n", userName, deposit_amt)
	}
	return nil
}
func withdrawReply(withdraw_amt uint64,errorStr string) (myRpc.MyRpcProcedure) {
	switch errorStr{
		case BANK_ACCESS_ERROR:
			log.Printf("%s: Received Withdraw Reply - Error accessing bank",userName)
		case LOW_BALANCE_ERROR:
			log.Printf("%s: Received Withdraw Reply - Insufficient Funds: $ %v\n",userName, withdraw_amt)
		default:
			log.Printf("%s: Received Withdraw Reply - Amount Withdrawn $ %v\n", userName, withdraw_amt)
	}
	return nil
}

func transferReply(transfer_amt uint64, target string, errorStr string) (myRpc.MyRpcProcedure){
	switch errorStr{
	case BANK_ACCESS_ERROR:
		log.Printf("%s: Received Transfer Reply - Error accessing bank",userName)
	case LOW_BALANCE_ERROR:
		log.Printf("%s: Received Transfer Reply - Insufficient Funds: $ %v\n",userName, transfer_amt)
	default:
		log.Printf("%s: Received Transfer Reply - $ %v transferred to %s\n", userName, transfer_amt, target)
}
return nil

}

func make_rpc_call(rpc_call defined.Rpc) {
	ipc_fd, ipc_status := altEthos.IpcRepeat("myRpc", "", nil)
	if ipc_status != syscall.StatusOk {
		log.Printf("IPC failed: %v\n", ipc_status)
		altEthos.Exit(ipc_status)
	}

	callStatus := altEthos.ClientCall(ipc_fd, rpc_call)
	if callStatus != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", callStatus)
		altEthos.Exit(callStatus)
	}
}







func main () {

	altEthos.LogToDirectory("test/bankingClient")
	
	log.Println("before call")
	userName = altEthos.GetUser()

	make_rpc_call(&(myRpc.MyRpcBalance{userName}))
	make_rpc_call(&(myRpc.MyRpcDeposit{userName,100}))
	make_rpc_call(&(myRpc.MyRpcBalance{userName}))
	make_rpc_call(&(myRpc.MyRpcWithdraw{userName,50}))
	make_rpc_call(&(myRpc.MyRpcBalance{userName}))
	make_rpc_call(&(myRpc.MyRpcWithdraw{userName,150}))
	make_rpc_call(&(myRpc.MyRpcBalance{userName}))
	make_rpc_call(&(myRpc.MyRpcTransfer{userName, "pat", 20}))
	make_rpc_call(&(myRpc.MyRpcBalance{userName}))






	log.Println("banking-client: done")
}

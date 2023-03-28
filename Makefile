export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=ethos
export GOLINUXINCLUDE=linux
export BUILD=ethos

export ETHOSROOT=client/rootfs
export MINIMALTDROOT=client/minimaltdfs


.PHONY: all install clean
all: bankingClient bankingServer

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

myRpc.go: myRpc.t
	$(ETN2GO) . myRpc $^

myRpc.goo.ethos : myRpc.go ethos
	ethosGoPackage  myRpc ethos myRpc.go

bankingServer: bankingServer.go myRpc.goo.ethos
	ethosGo bankingServer.go

bankingClient: bankingClient.go myRpc.goo.ethos
	ethosGo bankingClient.go

# install types, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	ethosTypeInstall myRpc
	ethosDirCreate $(ETHOSROOT)/services/myRpc   $(ETHOSROOT)/types/spec/myRpc/MyRpc all
	install -D  bankingClient bankingServer                   $(ETHOSROOT)/programs
	ethosStringEncode /programs/bankingServer    > $(ETHOSROOT)/etc/init/services/bankingServer

# remove build artifacts
clean:
	sudo rm -rf myRpc/ myRpcIndex/ ethos client
	rm -f myRpc.go
	rm -f bankingClient
	rm -f bankingClient.goo.ethos
	rm -f bankingServer
	rm -f bankingServer.goo.ethos
	rm -f myRpc.goo.ethos

record:
	ethosLog $(ETHOSROOT)/log/test/bankingServer > banking_server_op.txt


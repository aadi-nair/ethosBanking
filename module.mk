userspace.myRpc.dir := userspace/myRpc

# remote procedure interface between client and server
userspace.myRpc.types        = $(patsubst %, $(userspace.myRpc.dir)/%, myRpc.t)
userspace.myRpc.packgage = $(userspace.myRpc.dir)/myRpc.go


# client build definitions
userspace.myRpcClient.go     = myRpcClient.go
userspace.myRpcClient.src    = $(patsubst %, $(userspace.myRpc.dir)/%, $(userspace.myRpcClient.go))
userspace.myRpcClient.target = $(UDIR)/$(userspace.myRpc.dir)/myRpcClient


# server build definitions
userspace.myRpcService.go    = myRpcService.go
userspace.myRpcService.src   = $(patsubst %, $(userspace.myRpc.dir)/%, $(userspace.myRpcService.go))
userspace.myRpcService.target= $(UDIR)/$(userspace.myRpc.dir)/myRpcService

# install these executables
INSTALL_ETHOSROOT_PROGRAMS        += $(userspace.myRpcClient.target)
INSTALL_ETHOSROOT_SYSTEM_PROGRAMS += $(userspace.myRpcService.target)

# compile the types
$(userspace.myRpc.dir)/myRpc.go: $(userspace.myRpc.types)
	$(ETN2GO) $(@D) myRpc $^

$(userspace.myRpc.package) : $(userspace.myRpc.dir)/myRpc.go


# compile the client
$(userspace.myRpcClient.target): $(userspace.myRpcClient.src) $(userspace.myRpc.package)
	sh bin/ethosGo  $(userspace.myRpcClient.src)

# compile the server
$(userspace.myRpcService.target): $(userspace.myRpcService.src) $(userspace.myRpc.package)
	sh bin/ethosGo $(userspace.myRpcService.src)


# build everything
userspace.myRpc.all: $(userspace.myRpcClient.target) $(userspace.myRpcService.target)

# install types, service,
userspace.myRpc.installfs:
	bin/ethosTypeInstall $(userspace.myRpc.dir)/myRpc
	bin/ethosServiceInstall myRpc MyRpc
	ethosStringEncode /system/programs/myRpcService    > $(install.ethosRoot.init.services)/myRpcService

# remove build artifacts
userspace.myRpc.clean:
	rm -rf $(userspace.myRpc.dir)/myRpc
	rm -f  $(userspace.myRpc.dir)/myRpc.go


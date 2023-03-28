# Homework 2 (Aditya Nair - anair38)

## How to Run this project

1. Copy the following files into a folder in common directory like `/tmp/`
   - `myRpc.t`
   - `Makefile`
   - `bankingClient.go`
   - `bankingServer.go`
2. Change directory to the project folder and run `make clean && make install`
3. Change directory to `client` and run `sudo -E ethosRun` to start the banking-server
4. Open another terminal, login as `me`, change directory to `client` and run `etAl client.ethos`
5. Once the above command runs successfully you'll be logged into an inner VM where you type `bankingClient`. This will spawn a banking client that will execute a set of commands as specified in `bankingClient.go`
6. Repeat steps 4 and 5, logged in as `mike` and `pat`

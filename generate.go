package gopay

//go:generate ./solc-static-linux-8-7 --overwrite --bin --abi -o mainnet/contracts mainnet/contracts/channel.sol
//go:generate ./solc-static-linux-8-7 --overwrite --bin --abi -o l2/contracts l2/contracts/channel.sol
//go:generate abigen --pkg tests --type Channel --abi mainnet/contracts/Channel.abi --bin mainnet/contracts/Channel.bin --out mainnet/tests/channel.go
//go:generate abigen --pkg l2 --type Channel --abi l2/contracts/Channel.abi --bin l2/contracts/Channel.bin --out l2/channel.go

package gopay

//go:generate abigen --pkg tests --sol mainnet/contracts/channel.sol --out mainnet/tests/channel.go --solc ./solc-static-linux-8-7
//go:generate abigen --pkg tests --sol l2/contracts/channel.sol --out l2/tests/channel.go --solc ./solc-static-linux-8-7

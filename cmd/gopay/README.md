## Manual test

Start Clef
```
~/go/src/github.com/ethereum/go-ethereum/build/bin/clef --advanced --chainid 69
```

Start two terminals for both nodes
```
a > ./gopay
a > Start Listener
```

```
b > ./gopay
b > Connect to Peer
c > 127.0.0.1
```

Clef endpoint: /home/matematik/.clef/clef.ipc
Optimism endpoint: https://goerli.optimism.io

Contract on kovan optimism: 0x67fC157cf79AF3d66A24D97A8fec285Abbdd9497
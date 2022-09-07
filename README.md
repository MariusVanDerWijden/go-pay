# SSPC (Stupid Simple Payment Channels)

This repo contains an implementation for very simple payment channels.
This implementation aims to provide a building block for developers to integrate Payment Channels into their own application.

## Rationale
Over the years I came to the conclusion that Rollups are still not good enough for microtransactions.
There have been multiple companies building Payment and State Channel Networks.
These Channel networks allow to route transactions over multiple people. 
However this increases the complexity of the protocol as well as hugely increasing the complexity of the node software.
In my opinion 95% of use cases where a Payment Channel Network is useful can also be solved by simple bidirectional Payment Channels.
As long as the opening and closing of the Payment Channel is cheap (e.g. by opening and closing via an L2).
SPCC should be a building block enabling developers to quickly integrate it into their own application.

## Wrappers
The project currently supplies a golang API that can be used by golang application to start using Payment Channels.
The API is designed to be as simple as possible.
With `NewChannel` a channel object is instantiated, with `Open` you can open a channel etc.
An example on how to use the golang wrapper is provided in `cmd/gopay`.
I hope to implement JS and Android/IOS wrappers soon, so the contracts can be used from multiple different languages.

## Contracts
There are two different types of contracts.
The L2 contracts in package `l2/contracts` are designed to be optimal on a rollup (where state is less expensive than calldata).
The mainnet contracts in package `mainnet/contracts` are designed to be optimal on mainnet (where state is costly but calldata is cheap).
**Neither of the contracts are audited, use at your own risk**

## Sample project
The folder `cmd/deploy` contains a small program to deploy the contracts on a blockchain.
The folder `cmd/gopay` contains a small sample project that should show the usage of the golang API.
It showcases a full node implementation that can be used between two peers to send funds between each other.
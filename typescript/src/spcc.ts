import { ethers } from "ethers";

type Round = {
    number: bigint
    valueA: bigint
    valueB: bigint
    peerSigs: ethers.Signature[]
}

export function NewChannel(
    backend: ethers.Signer, 
    addrA: ethers.Addressable, 
    addrB: ethers.Addressable, 
    valueA: bigint, 
    valueB: bigint): Channel 
{
    let ch: Channel
    ch.backend = backend
    ch.addrA = addrA
    ch.addrB = addrB
    ch.valueA = valueA
    ch.valueB = valueB
    ch.sumFunds = valueA + valueB
    let firstRound: Round
    firstRound.number = 0n
    firstRound.valueA = valueA
    firstRound.valueB = valueB
    ch.rounds = new Map<bigint, Round>([
        [0n, firstRound]
    ]);
    return ch
}

class Channel {
    contract: ethers.Contract;
    backend: ethers.Signer;
    ID: string
    addrA: ethers.Addressable
    addrB: ethers.Addressable
    valueA: bigint
    valueB: bigint
    sumFunds: bigint
    rounds: Map<bigint, Round>
    userIsProposer: boolean

    Open(this: Channel): string {
        let id = ""
        // let id = crypto.randomBytes
        return id;
    };
}
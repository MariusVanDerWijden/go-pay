import { ethers, makeError } from "ethers";

// CoopCloseRound 
// CoopCloseRound, _ = new(big.Int).SetString("0xffffffffffffffffffffffffffffffff", 0)
var CoopCloseRound: bigint

export class Round {
    number: bigint
    valueA: bigint
    valueB: bigint
    peerSigs: string[] | undefined

    constructor(number: bigint, valueA: bigint, valueB: bigint, sig: string) {
        this.number = number
        this.valueA = valueA
        this.valueB = valueB
        this.peerSigs = new Array
        this.peerSigs[0] = sig
    }
}

export class Channel {
    contract: ethers.Contract | undefined; // TODO remove undefined
    backend: ethers.Signer;
    ID: string | undefined
    addrA: string
    addrB: string
    valueA: bigint
    valueB: bigint
    round: bigint
    sumFunds: bigint
    rounds: Map<bigint, Round>
    userIsProposer: boolean


    constructor(
        backend: ethers.Signer, 
        addrA: string, 
        addrB: string, 
        valueA: bigint, 
        valueB: bigint)
    {
        this.backend = backend
        this.addrA = addrA
        this.addrB = addrB
        this.valueA = valueA
        this.valueB = valueB
        this.round = 0n
        this.userIsProposer = false;
        this.sumFunds = valueA + valueB
        let firstRound = new Round(0n, valueA, valueB, "");
        this.rounds = new Map<bigint, Round>([
            [0n, firstRound]
        ]);
}


    CurrentState(this: Channel): Round | undefined {
        return this.rounds.get(this.round)
    }

    Open(this: Channel): string {
        let id = ""
        // let id = crypto.randomBytes
        this.userIsProposer = true
        return id;
    };

    Accept(this: Channel): string {
        this.userIsProposer = false
        return ""
    }

    CoopClose(this: Channel): string {
        return ""
    }

    CreateCoopClose(this: Channel): string {
        return ""
    }

    StartForceClose(this: Channel): string {
        return ""
    }

    DisputeForceClose(this: Channel): string {
        return ""
    }

    FinishForceClose(this: Channel): string {
        return ""
    }

    SendMoney(this: Channel): string {
        return ""
    }

    ReceivedMoney(this: Channel, valueA: bigint, valueB: bigint, round: bigint, sig: string): string {
        if(valueA + valueB > this.sumFunds) {
            return "sum of funds not equal" // TODO error handling
        }
        if(round <= this.round || this.round != CoopCloseRound) {
            return "invalid round" // TODO error handling
        }
        var hash = "TODO"

        var signer: string
        if(this.userIsProposer) {
            if(this.valueA < valueA) {
                return "valueA decreased by B" // TODO
            }
            signer = this.addrB
        } else {
            if(this.valueB < valueB) {
                return "valueB decreased by A" // TODO
            }
            signer = this.addrA
        }
        let recovered = ethers.recoverAddress(hash, sig)
        if(recovered != signer) {
            return "invalid signer"
        }
        var newRound = new Round(round, valueA, valueB, sig)
        this.rounds.set(round, newRound)
        return ""  
    }
    
}
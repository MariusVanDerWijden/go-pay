// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./ECDSA.sol";

contract Channel {

    using ECDSA for bytes32;
	uint64 disputePeriod = 1 days;
    uint128 cooperativeCloseRound = uint128(0xffffffffffffffffffffffffffffffff);

	enum Prog {NONEXISTANT,PROPOSED,ACCEPTED,DISPUTED,CLOSED}

	struct ChannelState {
		address a;
		address b;
		uint256 valueA;
		uint256 valueB;
		Prog progression;
		uint128 round;
	}

	struct Dispute {
		uint64 time;
		address closer;
	}

	// channels is a mapping of channel identifiers
	// to hashes of channel state.
	mapping(bytes32 => ChannelState) public channels;
	mapping(bytes32 => Dispute) public disputes;

	event Open(bytes32 indexed ID);
	event Accepted(bytes32 indexed ID);
	event Closing(bytes32 indexed ID, uint128 round, uint64 time);
	event Closed(bytes32 indexed ID);

	modifier inState(bytes32 id, Prog p){
		require(channels[id].progression == p, "Channel is in invalid state");
		_;
	}

    modifier validState(bytes32 id, uint256 valueA, uint256 valueB) {
		require(channels[id].valueA + channels[id].valueB == valueA + valueB, "invalid state provided");
		_; 
	}

    function hashState(bytes32 id, ChannelState memory state, uint256 valueA, uint256 valueB, uint128 round) public pure returns (bytes32) {
		bytes memory encoded = abi.encode(id, state.a, state.b, valueA, valueB, round);
		return keccak256(encoded);
	}

	function other(bytes32 id, address _address) private view returns (address) {
		if(_address == channels[id].a)
			return channels[id].b;
		if(_address == channels[id].a)
			return channels[id].a;
		return address(0);
	}

	function open(
		bytes32 id, address addrB, uint256 valueA, uint256 valueB) 
		public payable inState(id, Prog.NONEXISTANT) {
			require(msg.value == valueA, "proposer must add enough value");
			channels[id].a = msg.sender;
            channels[id].b = addrB;
            channels[id].valueA = valueA;
            channels[id].valueB = valueB;
            channels[id].progression = Prog.PROPOSED;
			emit Open(id);
	}

    function accept(bytes32 id) 
        public payable inState(id, Prog.PROPOSED) {
            require(msg.sender == channels[id].b, "acceptor must be sender");
            require(msg.value == channels[id].valueB, "acceptor must add enough value");
            channels[id].progression = Prog.ACCEPTED;
			emit Accepted(id);
    }

	function challenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes memory sig) 
        public inState(id, Prog.ACCEPTED) validState(id, valueA, valueB) {
            address aorb = other(id, msg.sender);
			require(aorb != address(0), "invalid other");
			bytes32 hash = hashState(id, channels[id], valueA, valueB, round);
			require(ECDSA.recover(hash, sig) == aorb, "invalid signature");
            channels[id].progression = Prog.DISPUTED;
            disputes[id].closer = msg.sender;
            disputes[id].time = uint64(block.timestamp);
			emit Closing(id, round, uint64(block.timestamp));
    }

    function disputeChallenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes memory sig)
        public inState(id, Prog.DISPUTED) validState(id, valueA, valueB) {
            address aorb = other(id, msg.sender);
			require(aorb != address(0), "invalid other");
            require(disputes[id].closer == aorb, "invalid dispute challenger");
			bytes32 hash = hashState(id, channels[id], valueA, valueB, round);
			require(ECDSA.recover(hash, sig) == aorb, "invalid signature");
			require(channels[id].round < channels[id].round, "disputed challenge with old state");
            disputes[id].closer = msg.sender;
            disputes[id].time = uint64(block.timestamp);
			emit Closing(id, round, uint64(block.timestamp));
        }
    

	function payout(bytes32 id, uint256 valueA, uint256 valueB) internal {
		// brick the channel
		channels[id].progression = Prog.CLOSED;
		// ignore failure conditions to prevent griefing
		payable(channels[id].a).send(valueA);
		payable(channels[id].b).send(valueB);
	}

    function CooperativeClose(bytes32 id, uint256 valueA, uint256 valueB, bytes memory sig) 
        public inState(id, Prog.ACCEPTED) validState(id, valueA, valueB) {
            address aorb = other(id, msg.sender);
			bytes32 hash = hashState(id, channels[id], valueA, valueB, cooperativeCloseRound);
			require(ECDSA.recover(hash, sig) == aorb, "invalid signature");
			payout(id, valueA, valueB);
			emit Closed(id);
    }

    function ForceClose(bytes32 id) 
        public inState(id, Prog.DISPUTED) {
            require(disputes[id].time != 0, "no dispute available");
			require(disputes[id].time + disputePeriod < block.timestamp , "force close before dispute period");
			payout(id, channels[id].valueA, channels[id].valueB);
			emit Closed(id);
    }
}
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./helper.sol";

contract Channel {

	uint64 disputePeriod = 1 days;

	enum Progression {PROPOSED,ACCEPTED,CLOSED}

	struct ChannelState {
		address a;
		address b;
		uint256 valueA;
		uint256 valueB;
		Progression progression;
		uint128 round;
	}

	struct Dispute {
		uint64 time;
		address closer;
	}

	// channels is a mapping of channel identifiers
	// to hashes of channel state.
	mapping(bytes32 => bytes32) public channels;
	// disputes is a mapping of channel identifiers
	// to disputes
	mapping(bytes32 => Dispute) public disputes;

	event Open(bytes32 indexed ID);
	event Accepted(bytes32 indexed ID);
	event Closing(bytes32 indexed ID, uint128 round, uint64 time);
	event Closed(bytes32 indexed ID);

	modifier inState(ChannelState memory s, Progression p){
		require(s.progression == p, "Channel is in invalid state");
		_;
	}

	modifier validState(ChannelState memory oldState, ChannelState memory newState) {
		require(oldState.a == newState.a, "invalid state provided");
		require(oldState.b == newState.b, "invalid state provided");
		require(oldState.valueA + oldState.valueB == newState.valueB + newState.valueA, "invalid state provided");
		_; 
	}

	function hashState(ChannelState memory state) public pure returns (bytes32) {
		// TODO hash the state
		return 0;
	}

	function other(ChannelState memory state, address _address) private pure returns (address) {
		if(_address == state.a)
			return state.b;
		if(_address == state.b)
			return state.a;
		return address(0);
	}

	function open(
		ChannelState memory state,
		bytes32 id) 
		public payable {
			require(msg.sender == state.a, "proposer must be sender");
			require(msg.value == state.valueA, "proposer must add enough value");
			require(channels[id] == 0, "channel already in use");
			require(state.round == 0, "round set in channel opening");
			bytes32 hash = hashState(state);
			channels[id] = hash;
			emit Open(id);
	}

	function accept(
		ChannelState memory oldState, 
		ChannelState memory newState, 
		bytes32 id) 
		public payable inState(oldState, Progression.PROPOSED) inState(newState, Progression.ACCEPTED) validState(oldState, newState){
			require(channels[id] == hashState(oldState), "invalid old state");
			require(oldState.valueA == newState.valueA, "can not steal money from proposer");
			require(msg.sender == oldState.b, "acceptor must be sender");
			require(msg.value == oldState.valueB, "proposer must add enough value");
			require(newState.round == 0, "round set in channel accepting");
			bytes32 hash = hashState(newState);
			channels[id] = hash;
			emit Accepted(id);
	}

	function challenge(
		ChannelState memory oldState, 
		ChannelState memory newState,  
		bytes[] calldata sig, 
		bytes32 id)
		public payable inState(oldState, Progression.ACCEPTED) validState(oldState, newState) {
			require(channels[id] == hashState(oldState), "invalid old state");
			address aorb = other(oldState, msg.sender);
			require(aorb != address(0), "invalid other");
			bytes32 hash = hashState(newState);
			require(helper.verifySig(hash, sig, aorb), "invalid signature");
			require(disputes[id].time == 0, "dispute already in progress");
			channels[id] = hash;
			disputes[id].time = uint64(block.timestamp);
			disputes[id].closer = msg.sender;
			emit Closing(id, newState.round, uint64(block.timestamp));
	}
	
	function disputeChallenge(
		ChannelState calldata oldState, 
		ChannelState memory newState,  
		bytes[] calldata sig, 
		bytes32 id)
		public payable inState(oldState, Progression.ACCEPTED) validState(oldState, newState) {
			require(channels[id] == hashState(oldState), "invalid old state");
			address aorb = other(oldState, msg.sender);
			require(aorb == disputes[id].closer, "invalid closer");
			bytes32 hash = hashState(newState);
			require(helper.verifySig(hash, sig, aorb), "invalid signature");
			require(oldState.round < newState.round, "disputed challenge with old state");
			channels[id] = hash;
			disputes[id].time = uint64(block.timestamp);
			disputes[id].closer = msg.sender;
			emit Closing(id, newState.round, uint64(block.timestamp));
	}

	function payout(ChannelState memory newState, bytes32 id) internal {
		// brick the channel
		channels[id] = 0;
		// TODO transfer the funds
	}

	function CooperativeClose(
		ChannelState calldata oldState, 
		ChannelState memory newState,  
		bytes[] calldata sig, 
		bytes32 id)
		public payable inState(oldState, Progression.ACCEPTED) validState(oldState, newState) {
			require(channels[id] == hashState(oldState), "invalid old state");
			require(newState.progression == Progression.CLOSED, "new state not closed");
			address aorb = other(oldState, msg.sender);
			bytes32 hash = hashState(newState);
			require(helper.verifySig(hash, sig, aorb), "invalid signature");
			payout(newState, id);
			emit Closed(id);
	}

	function ForceClose(
		ChannelState calldata oldState,
		bytes32 id)
		public payable inState(oldState, Progression.ACCEPTED) {
			require(channels[id] == hashState(oldState), "invalid old state");
			require(disputes[id].time != 0, "no dispute available");
			require(disputes[id].time + disputePeriod < block.timestamp , "force close before dispute period");
			payout(oldState, id);
			emit Closed(id);
	}
}
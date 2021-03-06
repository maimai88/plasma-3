Deposit: __log__({depositor: address, value: int128(wei)})

authority: public(address)
last_child_block: public(int128)
last_parent_block: public(int128)
child_chain: public({
    root: bytes32,
    created_at: timestamp
}[int128])
exits: public({
    utxo: int128[3],
    owner: address,
    amount: int128,
}[int128])
# for testing
exitsOrder: int128[100]
currentExit: int128

@private
def membership(leaf: bytes32, utxo: int128[3], proof: bytes32[16]) -> bool:
    result: bytes32 = leaf
    index: int128 = utxo[1]
    for i in range(16):
        if index % 2 == 0:
            result = keccak256(concat(result, proof[i]))
        else:
            result = keccak256(concat(proof[i], result))
        index = floor(index/2)
    return result == self.child_chain[utxo[0]].root

@private
def ecrecover_bytes(h: bytes32, sig: bytes <= 66) -> address:
    # vyper fails to parse null byte into integer
    # v: int128 = convert(slice(sig, start=64, len=1), 'int128')
    # create empty bytes32 array, replace last byte with signature last byte
    # and parse it with extract32
    v: int128 = extract32(
        concat(
            slice(
                concat(convert(0, 'bytes32'), ''), start=0, len=31),
            slice(sig, start=64, len=1)),
        0, type=int128)
    if v < 27:
        v += 27
    rst: address = ecrecover(h,
                     convert(v, 'uint256'),
                     extract32(sig, 0, type=uint256),
                     extract32(sig, 32, type=uint256))
    return rst

@public
def __init__():
    self.authority = msg.sender
    self.last_child_block = 1
    self.last_parent_block = block.number
    self.currentExit = 0

@public
def submitBlock(root: bytes32):
    assert msg.sender == self.authority
    self.child_chain[self.last_child_block] = {
        root: root,
        created_at: block.timestamp
    }
    self.last_child_block += 1
    self.last_parent_block = block.number

@public
@payable
def deposit(tx: bytes <= 1024):
    # vyper compiler fails with  ValueError: source code string cannot contain null bytes when list contains many(?) zero integers
    #tx_list = RLPList(tx, [int128, int128, int128, int128, int128, int128,
    #                       address, int128, address, int128, int128])
    #assert tx_list[0] == 0
    #assert tx_list[3] == 0
    #assert tx_list[7] == convert(msg.value, 'int128')
    zero_bytes: bytes32
    nei: bytes <= 130
    root: bytes32 = keccak256(concat(tx, nei))
    for i in range(16):
        root = keccak256(concat(root, zero_bytes))
        zero_bytes = keccak256(concat(zero_bytes, zero_bytes))
    self.child_chain[self.last_child_block] = {
        root: root,
        created_at: block.timestamp
    }
    self.last_child_block += 1
    self.last_parent_block = block.number
    log.Deposit(msg.sender, msg.value)

@public
@payable
def withdraw(utxo: int128[3],
             tx: bytes <= 1024, proof: bytes32[16],
             sigs: bytes <= 132, confirmSigs: bytes <= 132):
    # TODO if child chain block is older then 7 days use block
    prio: int128 = utxo[0] * 1000000000 + utxo[1] * 10000 + utxo[2]
    # do i care if same exit will be added multiple times?
    # it looks like it won't have any effect, but cost more gas to a sender
    assert self.exits[prio].amount == 0
    tx_list = RLPList(tx, [int128, int128, int128, int128, int128, int128,
                           address, int128, address, int128, int128])
    assert self.membership(keccak256(concat(tx, sigs)), utxo, proof)
    txHash: bytes32 = keccak256(tx)
    confirmationHash: bytes32 = keccak256(concat(
        txHash, self.child_chain[utxo[0]].root,
    ))
    if tx_list[0] == 0 and tx_list[3] == 0:
        assert msg.sender == self.ecrecover_bytes(
            confirmationHash, slice(confirmSigs, start=0, len=65))
    if tx_list[0] != 0:
        assert self.ecrecover_bytes(txHash, slice(sigs, start=0, len=65)) == self.ecrecover_bytes(confirmationHash, slice(confirmSigs, start=0, len=65))
    if tx_list[3] != 0:
        assert self.ecrecover_bytes(txHash, slice(sigs, start=65, len=65)) == self.ecrecover_bytes(confirmationHash, slice(confirmSigs, start=65, len=65))
    # how to access static array with utxo[2]*2 + 6 ?
    if utxo[2] == 0:
        self.exits[prio] = {
            utxo: utxo,
            owner: tx_list[6],
            amount: tx_list[7],
        }
    else:
        self.exits[prio] = {
            utxo: utxo,
            owner: tx_list[8],
            amount: tx_list[9],
        }
    self.exitsOrder[self.currentExit] = prio
    self.currentExit += 1

@public
def challenge(prio: int128, utxo: int128[3], tx: bytes <= 1024,
              proof: bytes32[16], sigs: bytes <= 132, confirmSig: bytes <= 66):
    # no need to send utxo here, oindex from tx will be enough
    assert self.exits[prio].amount != 0
    tx_list = RLPList(tx, [int128, int128, int128, int128, int128, int128,
                           address, int128, address, int128, int128])
    exitutxo: int128[3] = self.exits[prio].utxo
    if utxo[2] == 0:
        assert exitutxo[0] == tx_list[0]
        assert exitutxo[1] == tx_list[1]
        assert exitutxo[2] == tx_list[2]
    else:
        assert exitutxo[0] == tx_list[3]
        assert exitutxo[1] == tx_list[4]
        assert exitutxo[2] == tx_list[5]
    assert self.exits[prio].owner == self.ecrecover_bytes(
        keccak256(concat(keccak256(tx),self.child_chain[utxo[0]].root)),
        confirmSig)
    assert self.membership(keccak256(concat(tx, sigs)), utxo, proof)
    # how to delete an item from mapping?
    self.exits[prio].amount = 0

@public
def finalize():
    # i definitely need dynamic array for priority queue
    for prio in self.exitsOrder:
        if prio == 0:
            continue
        if self.exits[prio].amount == 0:
            continue
        send(self.exits[prio].owner, as_wei_value(self.exits[prio].amount, 'wei'))
        self.exits[prio].amount = 0

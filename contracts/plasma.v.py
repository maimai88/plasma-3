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
def ecrecover_bytes(h: bytes32, sig: bytes <= 65) -> address:
    return ecrecover(h,
                     extract32(sig, 0, type=uint256),
                     extract32(sig, 32, type=uint256),
                     convert(
                         convert(
                             slice(sig, start=64, len=1), 'int128'), 'uint256'))

@public
def __init__():
    self.authority = msg.sender
    self.last_child_block = 1
    self.last_parent_block = block.number

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
    # vyper compiler fails with  ValueError: source code string cannot contain null bytes
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
             sigs: bytes <= 130, confirmSigs: bytes <= 130):
    # TODO if child chain block is older then 7 days use block
    prio: int128 = utxo[0] * 1000000000 + utxo[1] * 10000 + utxo[2]
    # do i care if same exit will be added multiple times?
    # it looks like it won't have any effect, but cost more gas to a sender
    assert self.exits[prio].amount == 0
    tx_list = RLPList(tx, [int128, int128, int128, int128, int128, int128,
                           address, int128, address, int128, int128])
    assert self.membership(keccak256(concat(tx, sigs)), utxo, proof)
    confirmationHash: bytes32 = keccak256(concat(
        keccak256(tx),
        sigs,
        self.child_chain[utxo[0]].root,
    ))
    txHash: bytes32 = keccak256(tx)
    if tx_list[0] == 0 and tx_list[3] == 0:
        assert msg.sender == self.ecrecover_bytes(
            confirmationHash,
            slice(confirmSigs, start=0, len=65))
    if tx_list[0] != 0:
        assert self.ecrecover_bytes(txHash, slice(sigs, start=0, len=65)) == self.ecrecover_bytes(
            confirmationHash, slice(confirmSigs, start=0, len=65))
    if tx_list[3] != 0:
        assert self.ecrecover_bytes(txHash, slice(sigs, start=65, len=65)) == self.ecrecover_bytes(
            confirmationHash, slice(confirmSigs, start=65, len=65))
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

@public
def challenge(prio: int128, utxo: int128[3], tx: bytes <= 1024,
              proof: bytes32[16], sigs: bytes <= 130, confirmSig: uint256[3]):
    assert self.exits[prio].amount != 0
    tx_list = RLPList(tx, [int128, int128, int128, int128, int128, int128,
                           address, int128, address, int128, int128])
    exitutxo: int128[3] = self.exits[prio].utxo
    # exit utxo must be used as an input for another tx
    if utxo[2] == 0:
        assert exitutxo[0] == tx_list[0]
        assert exitutxo[1] == tx_list[1]
        assert exitutxo[2] == tx_list[2]
    else:
        assert exitutxo[0] == tx_list[3]
        assert exitutxo[1] == tx_list[4]
        assert exitutxo[2] == tx_list[5]
    # challenger must prove that signer confirmation is the same as exit owner
    assert self.exits[prio].owner == ecrecover(
        keccak256(concat(
            keccak256(tx),
            sigs,
            self.child_chain[utxo[0]].root,
        )), confirmSig[0], confirmSig[1], confirmSig[2])
    # verify that challenger tx is included in the confirmed block
    assert self.membership(keccak256(concat(
        tx, sigs,
    )), utxo, proof)
    # how to delete mapping in vyper completely?
    self.exits[prio].amount = 0

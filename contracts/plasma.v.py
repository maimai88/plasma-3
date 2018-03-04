Deposit: __log__({depositor: address})

authority: public(address)
last_child_block: public(int128)
last_parent_block: public(int128)
child_chain: public({
    root: bytes32,
    created_at: timestamp
}[int128])

@public
def __init__():
    self.authority = msg.sender
    self.last_child_block = 1
    self.last_parent_block = block.number


@public
def submitBlock(root: bytes32):
    assert msg.sender == self.authority
    assert block.number  == self.last_child_block + 6
    self.child_chain[self.last_child_block] = {
        root: root,
        created_at: block.timestamp
    }
    self.last_child_block += 1
    self.last_parent_block = block.number

@public
@payable
def deposit(tx: bytes <= 1024):
    tx_list = RLPList(tx, [int128, int128, int128, int128, int128, int128,
                           address, int128, address, int128, int128])
    zero_bytes: bytes32
    root: bytes32 = keccak256(concat(tx, zero_bytes))
    for i in range(16):
        root = keccak256(concat(root, zero_bytes))
        zero_bytes = keccak256(concat(zero_bytes, zero_bytes))
    self.child_chain[self.last_child_block] = {
        root: zero_bytes,
        created_at: block.timestamp
    }
    self.last_child_block += 1
    self.last_parent_block = block.number
    log.Deposit(tx_list[6])
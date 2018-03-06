Deposit: __log__({depositor: address, value: int128(wei)})

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
    assert block.number >= self.last_child_block + 6
    self.child_chain[self.last_child_block] = {
        root: root,
        created_at: block.timestamp
    }
    self.last_child_block += 1
    self.last_parent_block = block.number

@public
@payable
def deposit(txHash: bytes32):
    # txHash is a hack to workaround problem with vyper and null bytes
    # this is serious security leak and should not be used by anyone
    zero_bytes: bytes32
    root: bytes32 = txHash
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

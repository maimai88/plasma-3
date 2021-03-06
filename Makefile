gen:
	vyper contracts/plasma.v.py -f json > /tmp/plasma.abi
	vyper contracts/plasma.v.py -f bytecode > /tmp/plasma.bytecode
	abigen -abi /tmp/plasma.abi -bin /tmp/plasma.bytecode -lang go --pkg plasma > pkg/plasma.go
	-rm /tmp/plasma.abi
	-rm /tmp/plasma.bytecode

test:
	go test ./pkg/... -count=3

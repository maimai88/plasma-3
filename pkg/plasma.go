// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plasma

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PlasmaABI is the input ABI used to generate the binding from.
const PlasmaABI = "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"depositor\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"__init__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"submitBlock\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"root\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":81368},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes\",\"name\":\"tx\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":93674},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"int128[3]\",\"name\":\"utxo\"},{\"type\":\"bytes\",\"name\":\"tx\"},{\"type\":\"bytes32[16]\",\"name\":\"proof\"},{\"type\":\"bytes\",\"name\":\"sigs\"},{\"type\":\"bytes\",\"name\":\"confirmSigs\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":172858},{\"name\":\"challenge\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"prio\"},{\"type\":\"int128[3]\",\"name\":\"utxo\"},{\"type\":\"bytes\",\"name\":\"tx\"},{\"type\":\"bytes32[16]\",\"name\":\"proof\"},{\"type\":\"bytes\",\"name\":\"sigs\"},{\"type\":\"uint256[3]\",\"name\":\"confirmSig\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":43794},{\"name\":\"authority\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663},{\"name\":\"last_child_block\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":693},{\"name\":\"last_parent_block\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":723},{\"name\":\"child_chain__root\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1020},{\"name\":\"child_chain__created_at\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1044},{\"name\":\"exits__utxo\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"},{\"type\":\"int128\",\"name\":\"arg1\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1285},{\"name\":\"exits__owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1110},{\"name\":\"exits__amount\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1134}]"

// PlasmaBin is the compiled bytecode used for deploying new contracts.
const PlasmaBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052341561009e57600080fd5b33600055600160015543600255611f1e56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526381da3d3e60005114156102855761028060046101403734156100b557600080fd5b3033146100c157600080fd5b606051602435806040519013585780919012156100dd57600080fd5b50606051604435806040519013585780919012156100fa57600080fd5b506060516064358060405190135857809190121561011757600080fd5b50610140516103c052610160600160200201516103e05261040060006010818352015b600060605160026103e051078060405190135857809190121561015c57600080fd5b14156101bb5760006103c0516020826104a00101526020810190506101c0610400516010811061018b57600080fd5b60200201516020826104a0010152602081019050806104a0526104a090508051602082012090506103c052610210565b60006101c061040051601081106101d157600080fd5b60200201516020826104200101526020810190506103c051602082610420010152602081019050806104205261042090508051602082012090506103c0525b6402540be40060a05160026402540be4006103e05102058060805190135857809190121561023d57600080fd5b056103e0525b815160010180835281141561013a575b5050600161016060006020020151600360c052602060c0200160c052602060c02001546103c0511460005260206000f3005b63178777d660005114156103f857604060046101403734156102a657600080fd5b3033146102b257600080fd5b60616024356004016101803760416024356004013511156102d257600080fd5b610140516103a0526101806020600060208351038113156102f257600080fd5b0460200260200181015190506103c05261018060206020602083510381131561031a57600080fd5b0460200260200181015190506103e05260406001602082066102e00160418284011061034557600080fd5b6041806103008260206020880688030161018001600060046018f150508181528090509050905080602001516000825180602090135857809190121561038a57600080fd5b601f6101000a820481151761039e57600080fd5b606051816020036101000a8304806040519013585780919012156103c157600080fd5b90509050905060008112156103d557600080fd5b61040052602060c060806103a060006001610bb8f15060c05160005260206000f3005b63baa476946000511415610473576020600461014037341561041957600080fd5b600054331461042757600080fd5b600154600360c052602060c0200160c052602060c02042815561014051600182015550600160605160018254018060405190135857809190121561046a57600080fd5b81555043600255005b6398b1e06a6000511415610637576020600461014037610420600435600401610160376104006004356004013511156104ab57600080fd5b6000610160610400806020846106c001018260208501600060046078f15050805182019150506105c06082806020846106c00101826020850160006004601ff1505080518201915050806106c0526106c090508051602082012090506106a052610ba060006010818352015b60006106a051602082610bc00101526020810190506105a051602082610bc001015260208101905080610bc052610bc090508051602082012090506106a05260006105a051602082610c400101526020810190506105a051602082610c4001015260208101905080610c4052610c4090508051602082012090506105a0525b8151600101808352811415610517575b5050600154600360c052602060c0200160c052602060c0204281556106a05160018201555060016060516001825401806040519013585780919012156105eb57600080fd5b8155504360025533610d005234610d20526040610cc05233610d005234610d20527ff7803fc136a91844adf45c625b7530131d2c51b80f640f303a166c2e1a27ba84610cc051610d00a1005b6334b9bb4c60005114156112c9576102c06004610140376060516004358060405190135857809190121561066a57600080fd5b506060516024358060405190135857809190121561068757600080fd5b50606051604435806040519013585780919012156106a457600080fd5b50610420606435600401610400376104006064356004013511156106c757600080fd5b60a261028435600401610840376082610284356004013511156106e957600080fd5b60a26102a4356004016109203760826102a43560040135111561070b57600080fd5b6060516101406002602002015160605160605161271061014060016020020151028060405190135857809190121561074257600080fd5b606051633b9aca0061014060006020020151028060405190135857809190121561076b57600080fd5b018060405190135857809190121561078257600080fd5b018060405190135857809190121561079957600080fd5b610a00526000610a0051600460c052602060c0200160c052602060c02054146107c157600080fd5b610ea0610400610440610a208251602084016000735185d17c44699cecc3133114f8df70753b856709611720f15050610180610a20511461080157600080fd5b610a2051610a200180602001516000825180602090135857809190121561082757600080fd5b601f6101000a820481151761083b57600080fd5b606051816020036101000a83048060405190135857809190121561085e57600080fd5b9050905090508152610a4051610a200180602001516000825180602090135857809190121561088c57600080fd5b601f6101000a82048115176108a057600080fd5b606051816020036101000a8304806040519013585780919012156108c357600080fd5b9050905090508160200152610a6051610a20018060200151600082518060209013585780919012156108f457600080fd5b601f6101000a820481151761090857600080fd5b606051816020036101000a83048060405190135857809190121561092b57600080fd5b9050905090508160400152610a8051610a200180602001516000825180602090135857809190121561095c57600080fd5b601f6101000a820481151761097057600080fd5b606051816020036101000a83048060405190135857809190121561099357600080fd5b9050905090508160600152610aa051610a20018060200151600082518060209013585780919012156109c457600080fd5b601f6101000a82048115176109d857600080fd5b606051816020036101000a8304806040519013585780919012156109fb57600080fd5b9050905090508160800152610ac051610a2001806020015160008251806020901358578091901215610a2c57600080fd5b601f6101000a8204811517610a4057600080fd5b606051816020036101000a830480604051901358578091901215610a6357600080fd5b9050905090508160a001526014610ae051610a20015114610a8357600080fd5b602051610ae051610a340151068160c00152610b0051610a2001806020015160008251806020901358578091901215610abb57600080fd5b601f6101000a8204811517610acf57600080fd5b606051816020036101000a830480604051901358578091901215610af257600080fd5b9050905090508160e001526014610b2051610a20015114610b1257600080fd5b602051610b2051610a34015106816101000152610b4051610a2001806020015160008251806020901358578091901215610b4b57600080fd5b601f6101000a8204811517610b5f57600080fd5b606051816020036101000a830480604051901358578091901215610b8257600080fd5b905090509050816101200152610b6051610a2001806020015160008251806020901358578091901215610bb457600080fd5b601f6101000a8204811517610bc857600080fd5b606051816020036101000a830480604051901358578091901215610beb57600080fd5b9050905090508161014001525060206117c06102846381da3d3e6114e05260006104006104008060208461100001018260208501600060046078f15050805182019150506108406082806020846110000101826020850160006004601ff1505080518201915050806110005261100090508051602082012090506115005261152061014080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015250506115806101c08060006020020151826000602002015280600160200201518260016020020152806002602002015182600260200201528060036020020151826003602002015280600460200201518260046020020152806005602002015182600560200201528060066020020151826006602002015280600760200201518260076020020152806008602002015182600860200201528060096020020151826009602002015280600a602002015182600a602002015280600b602002015182600b602002015280600c602002015182600c602002015280600d602002015182600d602002015280600e602002015182600e602002015280600f602002015182600f602002015250506114fc6000305af1610dba57600080fd5b6117c051610dc757600080fd5b60006104008051602082012090506020826118000101526020810190506108406082806020846118000101826020850160006004601ff1505080518201915050600161014060006020020151600360c052602060c0200160c052602060c0200154602082611800010152602081019050806118005261180090508051602082012090506117e052610400805160208201209050611920526000610f0051146000610ea051141615610f25576020611b6060c4604063178777d6611a40526117e051611a605280611a8052600060416020820661194001608282840110610eac57600080fd5b608280611960826020602088068803016109200160006004601ff15050818152809050905090508080516020018084611a6001828460006004600a8704601201f1610ef657600080fd5b505080518201602001915050611a5c90506000305af1610f1557600080fd5b611b60513314610f2457600080fd5b5b6000610ea05114151561108b576020611fe060c4604063178777d6611ec0526117e051611ee05280611f00526000604160208206611dc001608282840110610f6c57600080fd5b608280611de0826020602088068803016109200160006004601ff15050818152809050905090508080516020018084611ee001828460006004600a8704601201f1610fb657600080fd5b505080518201602001915050611edc90506000305af1610fd557600080fd5b611fe0516020611da060c4604063178777d6611c805261192051611ca05280611cc0526000604160208206611b800160828284011061101357600080fd5b608280611ba0826020602088068803016108400160006004601ff15050818152809050905090508080516020018084611ca001828460006004600a8704601201f161105d57600080fd5b505080518201602001915050611c9c90506000305af161107c57600080fd5b611da0511461108a57600080fd5b5b6000610f00511415156111f157602061246060c4604063178777d6612340526117e0516123605280612380526041604160208206612240016082828401106110d257600080fd5b608280612260826020602088068803016109200160006004601ff1505081815280905090509050808051602001808461236001828460006004600a8704601201f161111c57600080fd5b50508051820160200191505061235c90506000305af161113b57600080fd5b61246051602061222060c4604063178777d6612100526119205161212052806121405260416041602082066120000160828284011061117957600080fd5b608280612020826020602088068803016108400160006004601ff1505081815280905090509050808051602001808461212001828460006004600a8704601201f16111c357600080fd5b50508051820160200191505061211c90506000305af16111e257600080fd5b61222051146111f057600080fd5b5b600061014060026020020151141561126757610a0051600460c052602060c0200160c052602060c020610f80518155610f605160018201556002810160c052602060c0206101408060006020020151600083015580600160200201516001830155806002602002015160028301555050506112c7565b610a0051600460c052602060c0200160c052602060c020610fc0518155610fa05160018201556002810160c052602060c0206101408060006020020151600083015580600160200201516001830155806002602002015160028301555050505b005b6365f52cba6000511415611bdb5761032060046101403734156112eb57600080fd5b6060516004358060405190135857809190121561130757600080fd5b506060516024358060405190135857809190121561132457600080fd5b506060516044358060405190135857809190121561134157600080fd5b506060516064358060405190135857809190121561135e57600080fd5b506104206084356004016104603761040060843560040135111561138157600080fd5b60a26102a4356004016108a03760826102a4356004013511156113a357600080fd5b600061014051600460c052602060c0200160c052602060c0205414156113c857600080fd5b610e006104606104406109808251602084016000735185d17c44699cecc3133114f8df70753b856709611720f15050610180610980511461140857600080fd5b610980516109800180602001516000825180602090135857809190121561142e57600080fd5b601f6101000a820481151761144257600080fd5b606051816020036101000a83048060405190135857809190121561146557600080fd5b90509050905081526109a0516109800180602001516000825180602090135857809190121561149357600080fd5b601f6101000a82048115176114a757600080fd5b606051816020036101000a8304806040519013585780919012156114ca57600080fd5b90509050905081602001526109c051610980018060200151600082518060209013585780919012156114fb57600080fd5b601f6101000a820481151761150f57600080fd5b606051816020036101000a83048060405190135857809190121561153257600080fd5b90509050905081604001526109e0516109800180602001516000825180602090135857809190121561156357600080fd5b601f6101000a820481151761157757600080fd5b606051816020036101000a83048060405190135857809190121561159a57600080fd5b9050905090508160600152610a0051610980018060200151600082518060209013585780919012156115cb57600080fd5b601f6101000a82048115176115df57600080fd5b606051816020036101000a83048060405190135857809190121561160257600080fd5b9050905090508160800152610a20516109800180602001516000825180602090135857809190121561163357600080fd5b601f6101000a820481151761164757600080fd5b606051816020036101000a83048060405190135857809190121561166a57600080fd5b9050905090508160a001526014610a405161098001511461168a57600080fd5b602051610a40516109940151068160c00152610a6051610980018060200151600082518060209013585780919012156116c257600080fd5b601f6101000a82048115176116d657600080fd5b606051816020036101000a8304806040519013585780919012156116f957600080fd5b9050905090508160e001526014610a805161098001511461171957600080fd5b602051610a8051610994015106816101000152610aa0516109800180602001516000825180602090135857809190121561175257600080fd5b601f6101000a820481151761176657600080fd5b606051816020036101000a83048060405190135857809190121561178957600080fd5b905090509050816101200152610ac051610980018060200151600082518060209013585780919012156117bb57600080fd5b601f6101000a82048115176117cf57600080fd5b606051816020036101000a8304806040519013585780919012156117f257600080fd5b90509050905081610140015250610f60600261014051600460c052602060c0200160c052602060c0200160008160c052602060c0200154826000602002015260018160c052602060c0200154826001602002015260028160c052602060c02001548260026020020152505060006101606002602002015114156118bc57610e0051610f60600060200201511461188757600080fd5b610e2051610f60600160200201511461189f57600080fd5b610e4051610f6060026020020151146118b757600080fd5b611905565b610e6051610f6060006020020151146118d457600080fd5b610e8051610f6060016020020151146118ec57600080fd5b610ea051610f60600260200201511461190457600080fd5b5b6000610460805160208201209050602082610fc00101526020810190506108a0608280602084610fc00101826020850160006004601ff1505080518201915050600161016060006020020151600360c052602060c0200160c052602060c0200154602082610fc001015260208101905080610fc052610fc090508051602082012090506110e052610400600060200201516111005261040060016020020151611120526104006002602002015161114052602060c060806110e060006001610bb8f15060c051600161014051600460c052602060c0200160c052602060c0200154146119f057600080fd5b60206119606102846381da3d3e611680526000610460610400806020846111a001018260208501600060046078f15050805182019150506108a06082806020846111a00101826020850160006004601ff1505080518201915050806111a0526111a090508051602082012090506116a0526116c061016080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015250506117206101e08060006020020151826000602002015280600160200201518260016020020152806002602002015182600260200201528060036020020151826003602002015280600460200201518260046020020152806005602002015182600560200201528060066020020151826006602002015280600760200201518260076020020152806008602002015182600860200201528060096020020151826009602002015280600a602002015182600a602002015280600b602002015182600b602002015280600c602002015182600c602002015280600d602002015182600d602002015280600e602002015182600e602002015280600f602002015182600f6020020152505061169c6000305af1611bb257600080fd5b61196051611bbf57600080fd5b600061014051600460c052602060c0200160c052602060c02055005b63bf7e214f6000511415611c01573415611bf457600080fd5b60005460005260206000f3005b6324c5ad9a6000511415611c27573415611c1a57600080fd5b60015460005260206000f3005b635258b0936000511415611c4d573415611c4057600080fd5b60025460005260206000f3005b6324f00f5b6000511415611cb05760206004610140373415611c6e57600080fd5b60605160043580604051901358578091901215611c8a57600080fd5b50600161014051600360c052602060c0200160c052602060c020015460005260206000f3005b6388c73d2d6000511415611d105760206004610140373415611cd157600080fd5b60605160043580604051901358578091901215611ced57600080fd5b5061014051600360c052602060c0200160c052602060c0205460005260206000f3005b631e97cd2f6000511415611daa5760406004610140373415611d3157600080fd5b60605160043580604051901358578091901215611d4d57600080fd5b5060605160243580604051901358578091901215611d6a57600080fd5b506101605160038110611d7c57600080fd5b600261014051600460c052602060c0200160c052602060c0200160c052602060c020015460005260206000f3005b6377790f8d6000511415611e0d5760206004610140373415611dcb57600080fd5b60605160043580604051901358578091901215611de757600080fd5b50600161014051600460c052602060c0200160c052602060c020015460005260206000f3005b6349f1cc336000511415611e6d5760206004610140373415611e2e57600080fd5b60605160043580604051901358578091901215611e4a57600080fd5b5061014051600460c052602060c0200160c052602060c0205460005260206000f3005b5b6100b0611f1e036100b06000396100b0611f1e036000f3`

// DeployPlasma deploys a new Ethereum contract, binding an instance of Plasma to it.
func DeployPlasma(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Plasma, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlasmaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Plasma{PlasmaCaller: PlasmaCaller{contract: contract}, PlasmaTransactor: PlasmaTransactor{contract: contract}, PlasmaFilterer: PlasmaFilterer{contract: contract}}, nil
}

// Plasma is an auto generated Go binding around an Ethereum contract.
type Plasma struct {
	PlasmaCaller     // Read-only binding to the contract
	PlasmaTransactor // Write-only binding to the contract
	PlasmaFilterer   // Log filterer for contract events
}

// PlasmaCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaSession struct {
	Contract     *Plasma           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaCallerSession struct {
	Contract *PlasmaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PlasmaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaTransactorSession struct {
	Contract     *PlasmaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaRaw struct {
	Contract *Plasma // Generic contract binding to access the raw methods on
}

// PlasmaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaCallerRaw struct {
	Contract *PlasmaCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaTransactorRaw struct {
	Contract *PlasmaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasma creates a new instance of Plasma, bound to a specific deployed contract.
func NewPlasma(address common.Address, backend bind.ContractBackend) (*Plasma, error) {
	contract, err := bindPlasma(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Plasma{PlasmaCaller: PlasmaCaller{contract: contract}, PlasmaTransactor: PlasmaTransactor{contract: contract}, PlasmaFilterer: PlasmaFilterer{contract: contract}}, nil
}

// NewPlasmaCaller creates a new read-only instance of Plasma, bound to a specific deployed contract.
func NewPlasmaCaller(address common.Address, caller bind.ContractCaller) (*PlasmaCaller, error) {
	contract, err := bindPlasma(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaCaller{contract: contract}, nil
}

// NewPlasmaTransactor creates a new write-only instance of Plasma, bound to a specific deployed contract.
func NewPlasmaTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaTransactor, error) {
	contract, err := bindPlasma(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaTransactor{contract: contract}, nil
}

// NewPlasmaFilterer creates a new log filterer instance of Plasma, bound to a specific deployed contract.
func NewPlasmaFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaFilterer, error) {
	contract, err := bindPlasma(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaFilterer{contract: contract}, nil
}

// bindPlasma binds a generic wrapper to an already deployed contract.
func bindPlasma(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plasma *PlasmaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Plasma.Contract.PlasmaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plasma *PlasmaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plasma.Contract.PlasmaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plasma *PlasmaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plasma.Contract.PlasmaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plasma *PlasmaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Plasma.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plasma *PlasmaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plasma.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plasma *PlasmaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plasma.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(out address)
func (_Plasma *PlasmaCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(out address)
func (_Plasma *PlasmaSession) Authority() (common.Address, error) {
	return _Plasma.Contract.Authority(&_Plasma.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(out address)
func (_Plasma *PlasmaCallerSession) Authority() (common.Address, error) {
	return _Plasma.Contract.Authority(&_Plasma.CallOpts)
}

// Child_chain__created_at is a free data retrieval call binding the contract method 0x88c73d2d.
//
// Solidity: function child_chain__created_at(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaCaller) Child_chain__created_at(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "child_chain__created_at", arg0)
	return *ret0, err
}

// Child_chain__created_at is a free data retrieval call binding the contract method 0x88c73d2d.
//
// Solidity: function child_chain__created_at(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaSession) Child_chain__created_at(arg0 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Child_chain__created_at(&_Plasma.CallOpts, arg0)
}

// Child_chain__created_at is a free data retrieval call binding the contract method 0x88c73d2d.
//
// Solidity: function child_chain__created_at(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Child_chain__created_at(arg0 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Child_chain__created_at(&_Plasma.CallOpts, arg0)
}

// Child_chain__root is a free data retrieval call binding the contract method 0x24f00f5b.
//
// Solidity: function child_chain__root(arg0 int128) constant returns(out bytes32)
func (_Plasma *PlasmaCaller) Child_chain__root(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "child_chain__root", arg0)
	return *ret0, err
}

// Child_chain__root is a free data retrieval call binding the contract method 0x24f00f5b.
//
// Solidity: function child_chain__root(arg0 int128) constant returns(out bytes32)
func (_Plasma *PlasmaSession) Child_chain__root(arg0 *big.Int) ([32]byte, error) {
	return _Plasma.Contract.Child_chain__root(&_Plasma.CallOpts, arg0)
}

// Child_chain__root is a free data retrieval call binding the contract method 0x24f00f5b.
//
// Solidity: function child_chain__root(arg0 int128) constant returns(out bytes32)
func (_Plasma *PlasmaCallerSession) Child_chain__root(arg0 *big.Int) ([32]byte, error) {
	return _Plasma.Contract.Child_chain__root(&_Plasma.CallOpts, arg0)
}

// Exits__amount is a free data retrieval call binding the contract method 0x49f1cc33.
//
// Solidity: function exits__amount(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaCaller) Exits__amount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "exits__amount", arg0)
	return *ret0, err
}

// Exits__amount is a free data retrieval call binding the contract method 0x49f1cc33.
//
// Solidity: function exits__amount(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaSession) Exits__amount(arg0 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Exits__amount(&_Plasma.CallOpts, arg0)
}

// Exits__amount is a free data retrieval call binding the contract method 0x49f1cc33.
//
// Solidity: function exits__amount(arg0 int128) constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Exits__amount(arg0 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Exits__amount(&_Plasma.CallOpts, arg0)
}

// Exits__owner is a free data retrieval call binding the contract method 0x77790f8d.
//
// Solidity: function exits__owner(arg0 int128) constant returns(out address)
func (_Plasma *PlasmaCaller) Exits__owner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "exits__owner", arg0)
	return *ret0, err
}

// Exits__owner is a free data retrieval call binding the contract method 0x77790f8d.
//
// Solidity: function exits__owner(arg0 int128) constant returns(out address)
func (_Plasma *PlasmaSession) Exits__owner(arg0 *big.Int) (common.Address, error) {
	return _Plasma.Contract.Exits__owner(&_Plasma.CallOpts, arg0)
}

// Exits__owner is a free data retrieval call binding the contract method 0x77790f8d.
//
// Solidity: function exits__owner(arg0 int128) constant returns(out address)
func (_Plasma *PlasmaCallerSession) Exits__owner(arg0 *big.Int) (common.Address, error) {
	return _Plasma.Contract.Exits__owner(&_Plasma.CallOpts, arg0)
}

// Exits__utxo is a free data retrieval call binding the contract method 0x1e97cd2f.
//
// Solidity: function exits__utxo(arg0 int128, arg1 int128) constant returns(out int128)
func (_Plasma *PlasmaCaller) Exits__utxo(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "exits__utxo", arg0, arg1)
	return *ret0, err
}

// Exits__utxo is a free data retrieval call binding the contract method 0x1e97cd2f.
//
// Solidity: function exits__utxo(arg0 int128, arg1 int128) constant returns(out int128)
func (_Plasma *PlasmaSession) Exits__utxo(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Exits__utxo(&_Plasma.CallOpts, arg0, arg1)
}

// Exits__utxo is a free data retrieval call binding the contract method 0x1e97cd2f.
//
// Solidity: function exits__utxo(arg0 int128, arg1 int128) constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Exits__utxo(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Plasma.Contract.Exits__utxo(&_Plasma.CallOpts, arg0, arg1)
}

// Last_child_block is a free data retrieval call binding the contract method 0x24c5ad9a.
//
// Solidity: function last_child_block() constant returns(out int128)
func (_Plasma *PlasmaCaller) Last_child_block(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "last_child_block")
	return *ret0, err
}

// Last_child_block is a free data retrieval call binding the contract method 0x24c5ad9a.
//
// Solidity: function last_child_block() constant returns(out int128)
func (_Plasma *PlasmaSession) Last_child_block() (*big.Int, error) {
	return _Plasma.Contract.Last_child_block(&_Plasma.CallOpts)
}

// Last_child_block is a free data retrieval call binding the contract method 0x24c5ad9a.
//
// Solidity: function last_child_block() constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Last_child_block() (*big.Int, error) {
	return _Plasma.Contract.Last_child_block(&_Plasma.CallOpts)
}

// Last_parent_block is a free data retrieval call binding the contract method 0x5258b093.
//
// Solidity: function last_parent_block() constant returns(out int128)
func (_Plasma *PlasmaCaller) Last_parent_block(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Plasma.contract.Call(opts, out, "last_parent_block")
	return *ret0, err
}

// Last_parent_block is a free data retrieval call binding the contract method 0x5258b093.
//
// Solidity: function last_parent_block() constant returns(out int128)
func (_Plasma *PlasmaSession) Last_parent_block() (*big.Int, error) {
	return _Plasma.Contract.Last_parent_block(&_Plasma.CallOpts)
}

// Last_parent_block is a free data retrieval call binding the contract method 0x5258b093.
//
// Solidity: function last_parent_block() constant returns(out int128)
func (_Plasma *PlasmaCallerSession) Last_parent_block() (*big.Int, error) {
	return _Plasma.Contract.Last_parent_block(&_Plasma.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0x65f52cba.
//
// Solidity: function challenge(prio int128, utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSig uint256[3]) returns()
func (_Plasma *PlasmaTransactor) Challenge(opts *bind.TransactOpts, prio *big.Int, utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSig [3]*big.Int) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "challenge", prio, utxo, tx, proof, sigs, confirmSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x65f52cba.
//
// Solidity: function challenge(prio int128, utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSig uint256[3]) returns()
func (_Plasma *PlasmaSession) Challenge(prio *big.Int, utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSig [3]*big.Int) (*types.Transaction, error) {
	return _Plasma.Contract.Challenge(&_Plasma.TransactOpts, prio, utxo, tx, proof, sigs, confirmSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x65f52cba.
//
// Solidity: function challenge(prio int128, utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSig uint256[3]) returns()
func (_Plasma *PlasmaTransactorSession) Challenge(prio *big.Int, utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSig [3]*big.Int) (*types.Transaction, error) {
	return _Plasma.Contract.Challenge(&_Plasma.TransactOpts, prio, utxo, tx, proof, sigs, confirmSig)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(tx bytes) returns()
func (_Plasma *PlasmaTransactor) Deposit(opts *bind.TransactOpts, tx []byte) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "deposit", tx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(tx bytes) returns()
func (_Plasma *PlasmaSession) Deposit(tx []byte) (*types.Transaction, error) {
	return _Plasma.Contract.Deposit(&_Plasma.TransactOpts, tx)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(tx bytes) returns()
func (_Plasma *PlasmaTransactorSession) Deposit(tx []byte) (*types.Transaction, error) {
	return _Plasma.Contract.Deposit(&_Plasma.TransactOpts, tx)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_Plasma *PlasmaTransactor) SubmitBlock(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "submitBlock", root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_Plasma *PlasmaSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _Plasma.Contract.SubmitBlock(&_Plasma.TransactOpts, root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_Plasma *PlasmaTransactorSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _Plasma.Contract.SubmitBlock(&_Plasma.TransactOpts, root)
}

// Withdraw is a paid mutator transaction binding the contract method 0x34b9bb4c.
//
// Solidity: function withdraw(utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSigs bytes) returns()
func (_Plasma *PlasmaTransactor) Withdraw(opts *bind.TransactOpts, utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSigs []byte) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "withdraw", utxo, tx, proof, sigs, confirmSigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x34b9bb4c.
//
// Solidity: function withdraw(utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSigs bytes) returns()
func (_Plasma *PlasmaSession) Withdraw(utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSigs []byte) (*types.Transaction, error) {
	return _Plasma.Contract.Withdraw(&_Plasma.TransactOpts, utxo, tx, proof, sigs, confirmSigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x34b9bb4c.
//
// Solidity: function withdraw(utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSigs bytes) returns()
func (_Plasma *PlasmaTransactorSession) Withdraw(utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSigs []byte) (*types.Transaction, error) {
	return _Plasma.Contract.Withdraw(&_Plasma.TransactOpts, utxo, tx, proof, sigs, confirmSigs)
}

// PlasmaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Plasma contract.
type PlasmaDepositIterator struct {
	Event *PlasmaDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaDeposit represents a Deposit event raised by the Plasma contract.
type PlasmaDeposit struct {
	Depositor common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xf7803fc136a91844adf45c625b7530131d2c51b80f640f303a166c2e1a27ba84.
//
// Solidity: event Deposit(depositor address, value int128)
func (_Plasma *PlasmaFilterer) FilterDeposit(opts *bind.FilterOpts) (*PlasmaDepositIterator, error) {

	logs, sub, err := _Plasma.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &PlasmaDepositIterator{contract: _Plasma.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xf7803fc136a91844adf45c625b7530131d2c51b80f640f303a166c2e1a27ba84.
//
// Solidity: event Deposit(depositor address, value int128)
func (_Plasma *PlasmaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PlasmaDeposit) (event.Subscription, error) {

	logs, sub, err := _Plasma.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaDeposit)
				if err := _Plasma.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}


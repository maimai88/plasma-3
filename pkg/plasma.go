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
const PlasmaABI = "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"depositor\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"__init__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"submitBlock\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"root\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":81368},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes\",\"name\":\"tx\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":93674},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"int128[3]\",\"name\":\"utxo\"},{\"type\":\"bytes\",\"name\":\"tx\"},{\"type\":\"bytes32[16]\",\"name\":\"proof\"},{\"type\":\"bytes\",\"name\":\"sigs\"},{\"type\":\"bytes\",\"name\":\"confirmSigs\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":234070},{\"name\":\"challenge\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"prio\"},{\"type\":\"int128[3]\",\"name\":\"utxo\"},{\"type\":\"bytes\",\"name\":\"tx\"},{\"type\":\"bytes32[16]\",\"name\":\"proof\"},{\"type\":\"bytes\",\"name\":\"sigs\"},{\"type\":\"bytes\",\"name\":\"confirmSig\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50130},{\"name\":\"finalize\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41567484},{\"name\":\"authority\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":693},{\"name\":\"last_child_block\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":723},{\"name\":\"last_parent_block\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":753},{\"name\":\"child_chain__root\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1050},{\"name\":\"child_chain__created_at\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1074},{\"name\":\"exits__utxo\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"},{\"type\":\"int128\",\"name\":\"arg1\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1315},{\"name\":\"exits__owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1140},{\"name\":\"exits__amount\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1164}]"

// PlasmaBin is the compiled bytecode used for deploying new contracts.
const PlasmaBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052341561009e57600080fd5b33600055600160015543600255600060065561210d56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526381da3d3e60005114156102855761028060046101403734156100b557600080fd5b3033146100c157600080fd5b606051602435806040519013585780919012156100dd57600080fd5b50606051604435806040519013585780919012156100fa57600080fd5b506060516064358060405190135857809190121561011757600080fd5b50610140516103c052610160600160200201516103e05261040060006010818352015b600060605160026103e051078060405190135857809190121561015c57600080fd5b14156101bb5760006103c0516020826104a00101526020810190506101c0610400516010811061018b57600080fd5b60200201516020826104a0010152602081019050806104a0526104a090508051602082012090506103c052610210565b60006101c061040051601081106101d157600080fd5b60200201516020826104200101526020810190506103c051602082610420010152602081019050806104205261042090508051602082012090506103c0525b6402540be40060a05160026402540be4006103e05102058060805190135857809190121561023d57600080fd5b056103e0525b815160010180835281141561013a575b5050600161016060006020020151600360c052602060c0200160c052602060c02001546103c0511460005260206000f3005b63178777d660005114156104c057604060046101403734156102a657600080fd5b3033146102b257600080fd5b60626024356004016101803760426024356004013511156102d257600080fd5b60605160006000601f602082066102e0016020828401106102f257600080fd5b60208061030082602060208806880301600060006020826102800101526020810190508061028052610280905001600060046015f1505081815280905090509050601f8060208461042001018260208501600060046015f150508051820191505060406001602082066103600160428284011061036e57600080fd5b6042806103808260206020880688030161018001600060046018f150508181528090509050905060018060208461042001018260208501600060046012f1505080518201915050806104205261042090506020600060208351038113156103d457600080fd5b046020026020018101519050806040519013585780919012156103f657600080fd5b61022052601b61022051121561042b57610220606051601b8251018060405190135857809190121561042757600080fd5b8152505b610140516104a05261022051600081121561044557600080fd5b6104c05261018060206000602083510381131561046157600080fd5b0460200260200181015190506104e05261018060206020602083510381131561048957600080fd5b04602002602001810151905061050052602060c060806104a060006001610bb8f15060c051610480526104805160005260206000f3005b63baa47694600051141561053b57602060046101403734156104e157600080fd5b60005433146104ef57600080fd5b600154600360c052602060c0200160c052602060c02042815561014051600182015550600160605160018254018060405190135857809190121561053257600080fd5b81555043600255005b6398b1e06a60005114156106ff5760206004610140376104206004356004016101603761040060043560040135111561057357600080fd5b6000610160610400806020846106c001018260208501600060046078f15050805182019150506105c06082806020846106c00101826020850160006004601ff1505080518201915050806106c0526106c090508051602082012090506106a052610ba060006010818352015b60006106a051602082610bc00101526020810190506105a051602082610bc001015260208101905080610bc052610bc090508051602082012090506106a05260006105a051602082610c400101526020810190506105a051602082610c4001015260208101905080610c4052610c4090508051602082012090506105a0525b81516001018083528114156105df575b5050600154600360c052602060c0200160c052602060c0204281556106a05160018201555060016060516001825401806040519013585780919012156106b357600080fd5b8155504360025533610d005234610d20526040610cc05233610d005234610d20527ff7803fc136a91844adf45c625b7530131d2c51b80f640f303a166c2e1a27ba84610cc051610d00a1005b6334b9bb4c60005114156113aa576102c06004610140376060516004358060405190135857809190121561073257600080fd5b506060516024358060405190135857809190121561074f57600080fd5b506060516044358060405190135857809190121561076c57600080fd5b506104206064356004016104003761040060643560040135111561078f57600080fd5b60a461028435600401610840376084610284356004013511156107b157600080fd5b60a46102a4356004016109203760846102a4356004013511156107d357600080fd5b6060516101406002602002015160605160605161271061014060016020020151028060405190135857809190121561080a57600080fd5b606051633b9aca0061014060006020020151028060405190135857809190121561083357600080fd5b018060405190135857809190121561084a57600080fd5b018060405190135857809190121561086157600080fd5b610a00526000610a0051600460c052602060c0200160c052602060c020541461088957600080fd5b610ea0610400610440610a208251602084016000735185d17c44699cecc3133114f8df70753b856709611720f15050610180610a2051146108c957600080fd5b610a2051610a20018060200151600082518060209013585780919012156108ef57600080fd5b601f6101000a820481151761090357600080fd5b606051816020036101000a83048060405190135857809190121561092657600080fd5b9050905090508152610a4051610a200180602001516000825180602090135857809190121561095457600080fd5b601f6101000a820481151761096857600080fd5b606051816020036101000a83048060405190135857809190121561098b57600080fd5b9050905090508160200152610a6051610a20018060200151600082518060209013585780919012156109bc57600080fd5b601f6101000a82048115176109d057600080fd5b606051816020036101000a8304806040519013585780919012156109f357600080fd5b9050905090508160400152610a8051610a2001806020015160008251806020901358578091901215610a2457600080fd5b601f6101000a8204811517610a3857600080fd5b606051816020036101000a830480604051901358578091901215610a5b57600080fd5b9050905090508160600152610aa051610a2001806020015160008251806020901358578091901215610a8c57600080fd5b601f6101000a8204811517610aa057600080fd5b606051816020036101000a830480604051901358578091901215610ac357600080fd5b9050905090508160800152610ac051610a2001806020015160008251806020901358578091901215610af457600080fd5b601f6101000a8204811517610b0857600080fd5b606051816020036101000a830480604051901358578091901215610b2b57600080fd5b9050905090508160a001526014610ae051610a20015114610b4b57600080fd5b602051610ae051610a340151068160c00152610b0051610a2001806020015160008251806020901358578091901215610b8357600080fd5b601f6101000a8204811517610b9757600080fd5b606051816020036101000a830480604051901358578091901215610bba57600080fd5b9050905090508160e001526014610b2051610a20015114610bda57600080fd5b602051610b2051610a34015106816101000152610b4051610a2001806020015160008251806020901358578091901215610c1357600080fd5b601f6101000a8204811517610c2757600080fd5b606051816020036101000a830480604051901358578091901215610c4a57600080fd5b905090509050816101200152610b6051610a2001806020015160008251806020901358578091901215610c7c57600080fd5b601f6101000a8204811517610c9057600080fd5b606051816020036101000a830480604051901358578091901215610cb357600080fd5b9050905090508161014001525060206117c06102846381da3d3e6114e05260006104006104008060208461100001018260208501600060046078f15050805182019150506108406084806020846110000101826020850160006004601ff1505080518201915050806110005261100090508051602082012090506115005261152061014080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015250506115806101c08060006020020151826000602002015280600160200201518260016020020152806002602002015182600260200201528060036020020151826003602002015280600460200201518260046020020152806005602002015182600560200201528060066020020151826006602002015280600760200201518260076020020152806008602002015182600860200201528060096020020151826009602002015280600a602002015182600a602002015280600b602002015182600b602002015280600c602002015182600c602002015280600d602002015182600d602002015280600e602002015182600e602002015280600f602002015182600f602002015250506114fc6000305af1610e8257600080fd5b6117c051610e8f57600080fd5b6104008051602082012090506117e05260006117e051602082611820010152602081019050600161014060006020020151600360c052602060c0200160c052602060c020015460208261182001015260208101905080611820526118209050805160208201209050611800526000610f0051146000610ea051141615610fc2576020611ac060c4604063178777d66119a052611800516119c052806119e05260006041602082066118a001608482840110610f4957600080fd5b6084806118c0826020602088068803016109200160006004601ff150508181528090509050905080805160200180846119c001828460006004600a8704601201f1610f9357600080fd5b5050805182016020019150506119bc90506000305af1610fb257600080fd5b611ac0513314610fc157600080fd5b5b6000610ea051141515611128576020611f4060c4604063178777d6611e205261180051611e405280611e60526000604160208206611d200160848284011061100957600080fd5b608480611d40826020602088068803016109200160006004601ff15050818152809050905090508080516020018084611e4001828460006004600a8704601201f161105357600080fd5b505080518201602001915050611e3c90506000305af161107257600080fd5b611f40516020611d0060c4604063178777d6611be0526117e051611c005280611c20526000604160208206611ae0016084828401106110b057600080fd5b608480611b00826020602088068803016108400160006004601ff15050818152809050905090508080516020018084611c0001828460006004600a8704601201f16110fa57600080fd5b505080518201602001915050611bfc90506000305af161111957600080fd5b611d00511461112757600080fd5b5b6000610f005114151561128e5760206123c060c4604063178777d66122a052611800516122c052806122e05260416041602082066121a00160848284011061116f57600080fd5b6084806121c0826020602088068803016109200160006004601ff150508181528090509050905080805160200180846122c001828460006004600a8704601201f16111b957600080fd5b5050805182016020019150506122bc90506000305af16111d857600080fd5b6123c051602061218060c4604063178777d6612060526117e05161208052806120a0526041604160208206611f600160848284011061121657600080fd5b608480611f80826020602088068803016108400160006004601ff1505081815280905090509050808051602001808461208001828460006004600a8704601201f161126057600080fd5b50508051820160200191505061207c90506000305af161127f57600080fd5b612180511461128d57600080fd5b5b600061014060026020020151141561130457610a0051600460c052602060c0200160c052602060c020610f80518155610f605160018201556002810160c052602060c020610140806000602002015160008301558060016020020151600183015580600260200201516002830155505050611364565b610a0051600460c052602060c0200160c052602060c020610fc0518155610fa05160018201556002810160c052602060c0206101408060006020020151600083015580600160200201516001830155806002602002015160028301555050505b610a00516006546103e8811061137957600080fd5b600560c052602060c020015560066060516001825401806040519013585780919012156113a557600080fd5b815550005b63417807d36000511415611cdc576102e060046101403734156113cc57600080fd5b606051600435806040519013585780919012156113e857600080fd5b506060516024358060405190135857809190121561140557600080fd5b506060516044358060405190135857809190121561142257600080fd5b506060516064358060405190135857809190121561143f57600080fd5b506104206084356004016104203761040060843560040135111561146257600080fd5b60a46102a4356004016108603760846102a43560040135111561148457600080fd5b60626102c4356004016109403760426102c4356004013511156114a657600080fd5b600061014051600460c052602060c0200160c052602060c0205414156114cb57600080fd5b610e606104206104406109e08251602084016000735185d17c44699cecc3133114f8df70753b856709611720f150506101806109e0511461150b57600080fd5b6109e0516109e00180602001516000825180602090135857809190121561153157600080fd5b601f6101000a820481151761154557600080fd5b606051816020036101000a83048060405190135857809190121561156857600080fd5b9050905090508152610a00516109e00180602001516000825180602090135857809190121561159657600080fd5b601f6101000a82048115176115aa57600080fd5b606051816020036101000a8304806040519013585780919012156115cd57600080fd5b9050905090508160200152610a20516109e0018060200151600082518060209013585780919012156115fe57600080fd5b601f6101000a820481151761161257600080fd5b606051816020036101000a83048060405190135857809190121561163557600080fd5b9050905090508160400152610a40516109e00180602001516000825180602090135857809190121561166657600080fd5b601f6101000a820481151761167a57600080fd5b606051816020036101000a83048060405190135857809190121561169d57600080fd5b9050905090508160600152610a60516109e0018060200151600082518060209013585780919012156116ce57600080fd5b601f6101000a82048115176116e257600080fd5b606051816020036101000a83048060405190135857809190121561170557600080fd5b9050905090508160800152610a80516109e00180602001516000825180602090135857809190121561173657600080fd5b601f6101000a820481151761174a57600080fd5b606051816020036101000a83048060405190135857809190121561176d57600080fd5b9050905090508160a001526014610aa0516109e001511461178d57600080fd5b602051610aa0516109f40151068160c00152610ac0516109e0018060200151600082518060209013585780919012156117c557600080fd5b601f6101000a82048115176117d957600080fd5b606051816020036101000a8304806040519013585780919012156117fc57600080fd5b9050905090508160e001526014610ae0516109e001511461181c57600080fd5b602051610ae0516109f4015106816101000152610b00516109e00180602001516000825180602090135857809190121561185557600080fd5b601f6101000a820481151761186957600080fd5b606051816020036101000a83048060405190135857809190121561188c57600080fd5b905090509050816101200152610b20516109e0018060200151600082518060209013585780919012156118be57600080fd5b601f6101000a82048115176118d257600080fd5b606051816020036101000a8304806040519013585780919012156118f557600080fd5b90509050905081610140015250610fc0600261014051600460c052602060c0200160c052602060c0200160008160c052602060c0200154826000602002015260018160c052602060c0200154826001602002015260028160c052602060c02001548260026020020152505060006101606002602002015114156119bf57610e6051610fc0600060200201511461198a57600080fd5b610e8051610fc060016020020151146119a257600080fd5b610ea051610fc060026020020151146119ba57600080fd5b611a08565b610ec051610fc060006020020151146119d757600080fd5b610ee051610fc060016020020151146119ef57600080fd5b610f0051610fc06002602002015114611a0757600080fd5b5b60206111c060c4604063178777d66110a0526000610420805160208201209050602082611020010152602081019050600161016060006020020151600360c052602060c0200160c052602060c0200154602082611020010152602081019050806110205261102090508051602082012090506110c052806110e05261094080805160200180846110c001828460006004600a8704601201f1611aa957600080fd5b5050805182016020019150506110bc90506000305af1611ac857600080fd5b6111c051600161014051600460c052602060c0200160c052602060c020015414611af157600080fd5b60206119a06102846381da3d3e6116c0526000610420610400806020846111e001018260208501600060046078f15050805182019150506108606084806020846111e00101826020850160006004601ff1505080518201915050806111e0526111e090508051602082012090506116e05261170061016080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015250506117606101e08060006020020151826000602002015280600160200201518260016020020152806002602002015182600260200201528060036020020151826003602002015280600460200201518260046020020152806005602002015182600560200201528060066020020151826006602002015280600760200201518260076020020152806008602002015182600860200201528060096020020151826009602002015280600a602002015182600a602002015280600b602002015182600b602002015280600c602002015182600c602002015280600d602002015182600d602002015280600e602002015182600e602002015280600f602002015182600f602002015250506116dc6000305af1611cb357600080fd5b6119a051611cc057600080fd5b600061014051600460c052602060c0200160c052602060c02055005b634bb278f36000511415611dc5573415611cf557600080fd5b61016060006103e8818352015b61016051600560c052602060c0200154610140526000610140511415611d2757611db1565b600061014051600460c052602060c0200160c052602060c020541415611d4c57611db1565b6000600060006000600161014051600460c052602060c0200160c052602060c0205402600161014051600460c052602060c0200160c052602060c02001546000f1611d9657600080fd5b600061014051600460c052602060c0200160c052602060c020555b8151600101808352811415611d02575b5050005b63bf7e214f6000511415611deb573415611dde57600080fd5b60005460005260206000f3005b6324c5ad9a6000511415611e11573415611e0457600080fd5b60015460005260206000f3005b635258b0936000511415611e37573415611e2a57600080fd5b60025460005260206000f3005b6324f00f5b6000511415611e9a5760206004610140373415611e5857600080fd5b60605160043580604051901358578091901215611e7457600080fd5b50600161014051600360c052602060c0200160c052602060c020015460005260206000f3005b6388c73d2d6000511415611efa5760206004610140373415611ebb57600080fd5b60605160043580604051901358578091901215611ed757600080fd5b5061014051600360c052602060c0200160c052602060c0205460005260206000f3005b631e97cd2f6000511415611f945760406004610140373415611f1b57600080fd5b60605160043580604051901358578091901215611f3757600080fd5b5060605160243580604051901358578091901215611f5457600080fd5b506101605160038110611f6657600080fd5b600261014051600460c052602060c0200160c052602060c0200160c052602060c020015460005260206000f3005b6377790f8d6000511415611ff75760206004610140373415611fb557600080fd5b60605160043580604051901358578091901215611fd157600080fd5b50600161014051600460c052602060c0200160c052602060c020015460005260206000f3005b6349f1cc336000511415612057576020600461014037341561201857600080fd5b6060516004358060405190135857809190121561203457600080fd5b5061014051600460c052602060c0200160c052602060c0205460005260206000f3005b5b6100b561210d036100b56000396100b561210d036000f3`

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

// Challenge is a paid mutator transaction binding the contract method 0x417807d3.
//
// Solidity: function challenge(prio int128, utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSig bytes) returns()
func (_Plasma *PlasmaTransactor) Challenge(opts *bind.TransactOpts, prio *big.Int, utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSig []byte) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "challenge", prio, utxo, tx, proof, sigs, confirmSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x417807d3.
//
// Solidity: function challenge(prio int128, utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSig bytes) returns()
func (_Plasma *PlasmaSession) Challenge(prio *big.Int, utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSig []byte) (*types.Transaction, error) {
	return _Plasma.Contract.Challenge(&_Plasma.TransactOpts, prio, utxo, tx, proof, sigs, confirmSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x417807d3.
//
// Solidity: function challenge(prio int128, utxo int128[3], tx bytes, proof bytes32[16], sigs bytes, confirmSig bytes) returns()
func (_Plasma *PlasmaTransactorSession) Challenge(prio *big.Int, utxo [3]*big.Int, tx []byte, proof [16][32]byte, sigs []byte, confirmSig []byte) (*types.Transaction, error) {
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

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Plasma *PlasmaTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plasma.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Plasma *PlasmaSession) Finalize() (*types.Transaction, error) {
	return _Plasma.Contract.Finalize(&_Plasma.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Plasma *PlasmaTransactorSession) Finalize() (*types.Transaction, error) {
	return _Plasma.Contract.Finalize(&_Plasma.TransactOpts)
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


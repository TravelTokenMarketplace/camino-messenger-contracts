// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cmaccount

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PartnerConfigurationService is an auto generated low-level Go binding around an user-defined struct.
type PartnerConfigurationService struct {
	RestrictedRate bool
	Capabilities   []string
}

// CmaccountMetaData contains all meta data concerning the Cmaccount contract.
var CmaccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"latestImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountImplementationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountNoUpgradeNeeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"use\",\"type\":\"uint8\"}],\"name\":\"InvalidPublicKeyUseType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefundLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrefundNotSpentYet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceededForPeriod\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportsOffChainPayment\",\"type\":\"bool\"}],\"name\":\"OffChainPaymentSupportUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceCapabilitiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"ServiceRestrictedRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOKING_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER_BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVICE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasMoney\",\"type\":\"uint256\"}],\"name\":\"addMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"addService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"addServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"addWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"buyBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasMoneyWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGasMoneyWithdrawalForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKeysAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pubKeyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getService\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service\",\"name\":\"service\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"isBotAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"isServiceSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"mintBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offChainPaymentSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"removeMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"removePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"removeServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"removeSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"removeWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGasMoneyWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSupported\",\"type\":\"bool\"}],\"name\":\"setOffChainPaymentSupported\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"setServiceCapabilities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"setServiceRestrictedRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"reason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGasMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161485d62000104600039600081816126340152818161265d0152612978015261485d6000f3fe6080604052600436106103245760003560e01c806301ffc9a71461033057806308564c1914610365578063136f50ca14610387578063150b7a02146103a95780631aca6376146103ed5780631c54f0f71461040f5780631c5db99e1461042f578063241bbbfc1461044f578063248a9ca3146104645780632a119380146104925780632f2ff15d146104b2578063319d13f3146104d257806333746274146104f257806336568abe14610514578063383aba871461053457806339e4c7051461055657806342072bbd146105765780634f1ef2861461058b5780634f3f46391461059e57806351889d6b146105c057806352d1902d146105e0578063581ed290146105f557806358c0a4c2146106155780635c988994146106355780635e07f86914610655578063658db0af146106755780636d69fcaf146106985780636fc22cd1146106b857806372afa328146106d857806374aa2048146106fa57806374fe60e91461071a5780637512e55b1461073a578063763191901461075a5780637eec56c71461077a578063852b3ccb1461079d578063857cdbb8146107bf57806385f438c1146107ec5780638c20f5741461080e5780638f69347d1461082e5780639010d07c1461084e57806391d148541461086e5780639db5dbe41461088e578063a217fddf146108ae578063a31aa039146108c3578063a3246ad3146108e3578063a7d022f814610910578063ad3cb1cc14610930578063b512463514610961578063b82923fb14610981578063be66718814610996578063c162d7da146109b6578063c6640e68146109cb578063ca15c873146109eb578063ccde65dc14610a0b578063cd9ef91414610a2b578063d09445c214610a4b578063d3c7c2c714610a6d578063d547741f14610a82578063da47d85614610aa2578063e0b78add14610acf578063e26a61bb14610aef578063e5a6725c14610b0f578063e7bfce9a14610b2f578063ea79d07a14610b4f578063ebc20d2014610b64578063ee3b641f14610b84578063f3fef3a314610ba4578063f51acaea14610bc4578063f72c0d8b14610be4578063f7e45f0914610c06578063f8c8765e14610c2657600080fd5b3661032b57005b600080fd5b34801561033c57600080fd5b5061035061034b36600461393c565b610c46565b60405190151581526020015b60405180910390f35b34801561037157600080fd5b5061037a610c71565b60405161035c9190613a10565b34801561039357600080fd5b5061039c610d2a565b60405161035c9190613a23565b3480156103b557600080fd5b506103d46103c4366004613b3f565b630a85bd0160e11b949350505050565b6040516001600160e01b0319909116815260200161035c565b3480156103f957600080fd5b5061040d610408366004613baa565b610d4a565b005b34801561041b57600080fd5b5061040d61042a366004613beb565b610df7565b34801561043b57600080fd5b5061040d61044a366004613cab565b610e9b565b34801561045b57600080fd5b50610350610f53565b34801561047057600080fd5b5061048461047f366004613cdf565b610f6b565b60405190815260200161035c565b34801561049e57600080fd5b5061040d6104ad366004613d0f565b610f8b565b3480156104be57600080fd5b5061040d6104cd366004613d4b565b611031565b3480156104de57600080fd5b5061037a6104ed366004613d7b565b611053565b3480156104fe57600080fd5b506104846000805160206147c883398151915281565b34801561052057600080fd5b5061040d61052f366004613d4b565b611061565b34801561054057600080fd5b5061048460008051602061474883398151915281565b34801561056257600080fd5b5061040d610571366004613cab565b611094565b34801561058257600080fd5b5061039c611147565b61040d610599366004613daf565b61115e565b3480156105aa57600080fd5b506105b361117d565b60405161035c9190613dfe565b3480156105cc57600080fd5b5061040d6105db366004613e12565b61119b565b3480156105ec57600080fd5b5061048461126c565b34801561060157600080fd5b5061040d610610366004613e4e565b611289565b34801561062157600080fd5b50610350610630366004613d7b565b6112f8565b34801561064157600080fd5b5061040d610650366004613cdf565b61130b565b34801561066157600080fd5b5061037a610670366004613cdf565b611340565b34801561068157600080fd5b5061068a61143e565b60405161035c929190613ec1565b3480156106a457600080fd5b5061040d6106b3366004613ecf565b611460565b3480156106c457600080fd5b5061040d6106d3366004613beb565b611481565b3480156106e457600080fd5b5061048460008051602061472883398151915281565b34801561070657600080fd5b5061040d610715366004613eec565b6114a3565b34801561072657600080fd5b5061040d610735366004613d0f565b61155b565b34801561074657600080fd5b5061040d610755366004613f32565b611595565b34801561076657600080fd5b5061040d610775366004613ecf565b611611565b34801561078657600080fd5b5061078f611632565b60405161035c929190613fff565b3480156107a957600080fd5b5061048460008051602061480883398151915281565b3480156107cb57600080fd5b506107df6107da366004613ecf565b611779565b60405161035c9190614071565b3480156107f857600080fd5b506104846000805160206147e883398151915281565b34801561081a57600080fd5b5061040d610829366004613f32565b611867565b34801561083a57600080fd5b50610350610849366004613cdf565b6118d6565b34801561085a57600080fd5b506105b3610869366004613beb565b611904565b34801561087a57600080fd5b50610350610889366004613d4b565b611932565b34801561089a57600080fd5b5061040d6108a9366004613baa565b611968565b3480156108ba57600080fd5b50610484600081565b3480156108cf57600080fd5b5061040d6108de366004614084565b6119bb565b3480156108ef57600080fd5b506109036108fe366004613cdf565b6119dc565b60405161035c919061409f565b34801561091c57600080fd5b5061040d61092b3660046140e0565b611a09565b34801561093c57600080fd5b506107df604051806040016040528060058152602001640352e302e360dc1b81525081565b34801561096d57600080fd5b5061035061097c366004613d7b565b611a7a565b34801561098d57600080fd5b5061040d611a88565b3480156109a257600080fd5b5061040d6109b1366004613beb565b611b23565b3480156109c257600080fd5b506105b3611b5d565b3480156109d757600080fd5b5061040d6109e6366004613ecf565b611b78565b3480156109f757600080fd5b50610484610a06366004613cdf565b611c13565b348015610a1757600080fd5b5061040d610a26366004613daf565b611c38565b348015610a3757600080fd5b5061040d610a4636600461412d565b611c5a565b348015610a5757600080fd5b506104846000805160206147a883398151915281565b348015610a7957600080fd5b50610903611d12565b348015610a8e57600080fd5b5061040d610a9d366004613d4b565b611d2c565b348015610aae57600080fd5b50610ac2610abd366004613cdf565b611d48565b60405161035c9190614166565b348015610adb57600080fd5b50610350610aea366004613ecf565b611e69565b348015610afb57600080fd5b5061040d610b0a366004614179565b611e83565b348015610b1b57600080fd5b5061040d610b2a366004613cdf565b611f21565b348015610b3b57600080fd5b5061040d610b4a366004613ecf565b611fad565b348015610b5b57600080fd5b50610903611fce565b348015610b7057600080fd5b5061040d610b7f366004614208565b611fe8565b348015610b9057600080fd5b5061068a610b9f366004613ecf565b612055565b348015610bb057600080fd5b5061040d610bbf366004613e12565b612090565b348015610bd057600080fd5b5061040d610bdf366004613d7b565b612136565b348015610bf057600080fd5b5061048460008051602061476883398151915281565b348015610c1257600080fd5b5061040d610c21366004613eec565b61218a565b348015610c3257600080fd5b5061040d610c41366004614261565b6121c4565b60006001600160e01b03198216635a05180f60e01b1480610c6b5750610c6b82612377565b92915050565b60606000610c7d610d2a565b9050600081516001600160401b03811115610c9a57610c9a613a7c565b604051908082528060200260200182016040528015610ccd57816020015b6060815260200190600190039081610cb85790505b50905060005b8251811015610d2357610cfe838281518110610cf157610cf16142bd565b60200260200101516123ac565b828281518110610d1057610d106142bd565b6020908102919091010152600101610cd3565b5092915050565b60606000610d36612428565b9050610d448160090161244c565b91505090565b6000805160206147e8833981519152610d6281612459565b6001600160a01b038316610d8957604051633a954ecd60e21b815260040160405180910390fd5b604051632142170760e11b81523060048201526001600160a01b038481166024830152604482018490528516906342842e0e90606401600060405180830381600087803b158015610dd957600080fd5b505af1158015610ded573d6000803e3d6000fd5b5050505050505050565b600080516020614808833981519152610e0f81612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6321b87f3a610e3161117d565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018690526044810185905260640160006040518083038186803b158015610e7e57600080fd5b505af4158015610e92573d6000803e3d6000fd5b50505050505050565b6000805160206147a8833981519152610eb381612459565b60005b8251811015610f4e576000610ee3848381518110610ed657610ed66142bd565b6020026020010151612463565b9050610eee816124d9565b838281518110610f0057610f006142bd565b6020026020010151604051610f1591906142d3565b604051908190038120907f50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f890600090a250600101610eb6565b505050565b600080610f5e612428565b6003015460ff1692915050565b600080610f76612517565b60009384526020525050604090206001015490565b600080516020614808833981519152610fa381612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6307e47316610fc561117d565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024810187905261ffff80871660448301528516606482015260840160006040518083038186803b15801561101d57600080fd5b505af4158015610ded573d6000803e3d6000fd5b61103a82610f6b565b61104381612459565b61104d838361253b565b50505050565b6060610c6b6106708361257d565b6001600160a01b038116331461108a5760405163334bd91960e11b815260040160405180910390fd5b610f4e82826125b2565b6000805160206147a88339815191526110ac81612459565b60005b8251811015610f4e5760006110dc8483815181106110cf576110cf6142bd565b602002602001015161257d565b90506110e7816125eb565b8382815181106110f9576110f96142bd565b602002602001015160405161110e91906142d3565b604051908190038120907f0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e90600090a2506001016110af565b60606000611153612428565b9050610d448161244c565b611166612629565b61116f826126b9565b61117982826127fc565b5050565b6000806111886128b0565b600101546001600160a01b031692915050565b6000805160206147c88339815191526111b381612459565b6001600160a01b0383166111da57604051633a954ecd60e21b815260040160405180910390fd5b6111f26000805160206147288339815191528461253b565b5061120b6000805160206148088339815191528461253b565b506112246000805160206147488339815191528461253b565b506040516001600160a01b038416907fdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a87399490600090a2610f4e6001600160a01b038416836128d4565b600061127661296d565b5060008051602061478883398151915290565b6000805160206147a88339815191526112a181612459565b6112b46112ad85612463565b83856129b6565b836040516112c291906142d3565b604051908190038120907f763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae937527990600090a250505050565b6000610c6b6113068361257d565b612a3a565b611313612a51565b60008051602061474883398151915261132b81612459565b61133482612a87565b5061133d612bb4565b50565b6060600061134c612428565b90506113588382612bc5565b6000838152600282016020908152604080832060010180548251818502810185019093528083529193909284015b828210156114325783829060005260206000200180546113a5906142ef565b80601f01602080910402602001604051908101604052809291908181526020018280546113d1906142ef565b801561141e5780601f106113f35761010080835404028352916020019161141e565b820191906000526020600020905b81548152906001019060200180831161140157829003601f168201915b505050505081526020019060010190611386565b50505050915050919050565b600080600061144b612bef565b90508060020154816003015492509250509091565b6000805160206147a883398151915261147881612459565b61117982612c13565b6000805160206147c883398151915261149981612459565b610f4e8383612c89565b6000805160206148088339815191526114bb81612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63fd13a43e6114dd61117d565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018890526044810187905261ffff80871660648301528516608482015260a40160006040518083038186803b15801561153c57600080fd5b505af4158015611550573d6000803e3d6000fd5b505050505050505050565b60008051602061480883398151915261157381612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63793dddac610fc561117d565b6000805160206147a88339815191526115ad81612459565b6115bf6115b98461257d565b83612ce3565b826040516115cd91906142d3565b60405180910390207f498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf836040516116049190614071565b60405180910390a2505050565b6000805160206147a883398151915261162981612459565b61117982612d26565b606080600061163f611147565b9050600081516001600160401b0381111561165c5761165c613a7c565b60405190808252806020026020018201604052801561168f57816020015b606081526020019060019003908161167a5790505b509050600082516001600160401b038111156116ad576116ad613a7c565b6040519080825280602002602001820160405280156116e657816020015b6116d3613848565b8152602001906001900390816116cb5790505b50905060005b835181101561176e5761170a848281518110610cf157610cf16142bd565b83828151811061171c5761171c6142bd565b602002602001018190525061174984828151811061173c5761173c6142bd565b6020026020010151611d48565b82828151811061175b5761175b6142bd565b60209081029190910101526001016116ec565b509094909350915050565b60606000611785612428565b90506117946006820184612d9c565b6117bc578260405163ba650b5f60e01b81526004016117b39190613dfe565b60405180910390fd5b6001600160a01b0383166000908152600882016020526040902080546117e1906142ef565b80601f016020809104026020016040519081016040528092919081815260200182805461180d906142ef565b801561185a5780601f1061182f5761010080835404028352916020019161185a565b820191906000526020600020905b81548152906001019060200180831161183d57829003601f168201915b5050505050915050919050565b6000805160206147a883398151915261187f81612459565b61189161188b8461257d565b83612db1565b8260405161189f91906142d3565b60405180910390207fba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264836040516116049190614071565b6000806118e1612428565b90506118ed8382612bc5565b600092835260020160205250604090205460ff1690565b60008061190f612ee8565b600085815260208290526040902090915061192a9084612f0c565b949350505050565b60008061193d612517565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b6000805160206147e883398151915261198081612459565b6001600160a01b0383166119a757604051633a954ecd60e21b815260040160405180910390fd5b61104d6001600160a01b0385168484612f18565b6000805160206147a88339815191526119d381612459565b61117982612f70565b606060006119e8612ee8565b6000848152602082905260409020909150611a029061244c565b9392505050565b6000805160206147a8833981519152611a2181612459565b611a33611a2d8461257d565b83612fc7565b82604051611a4191906142d3565b6040519081900381208315158252907f23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab90602001611604565b6000610c6b6108498361257d565b6000805160206147a8833981519152611aa081612459565b6000611aaa611632565b50905060005b8151811015610f4e57611ad6611ad18383815181106110cf576110cf6142bd565b612ffe565b818181518110611ae857611ae86142bd565b6020026020010151604051611afd91906142d3565b6040519081900381209060008051602061470883398151915290600090a2600101611ab0565b600080516020614808833981519152611b3b81612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63348e06dd610e3161117d565b600080611b686128b0565b546001600160a01b031692915050565b6000805160206147c8833981519152611b9081612459565b611ba8600080516020614728833981519152836125b2565b50611bc1600080516020614808833981519152836125b2565b50611bda600080516020614748833981519152836125b2565b506040516001600160a01b038316907fd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc6291390600090a25050565b600080611c1e612ee8565b6000848152602082905260409020909150611a029061305f565b6000805160206147a8833981519152611c5081612459565b610f4e8383613069565b611c62612a51565b600080516020614808833981519152611c7a81612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__637adf63b7611c9c61117d565b6040516001600160e01b031960e084901b1681526001600160a01b0391821660048201526024810188905260448101879052908516606482015260840160006040518083038186803b158015611cf157600080fd5b505af4158015611d05573d6000803e3d6000fd5b5050505050610f4e612bb4565b60606000611d1e612428565b9050610d446004820161244c565b611d3582610f6b565b611d3e81612459565b61104d83836125b2565b611d50613848565b6000611d5a612428565b9050611d668382612bc5565b6000838152600282016020908152604080832081518083018352815460ff16151581526001820180548451818702810187019095528085529195929486810194939192919084015b82821015611e5a578382906000526020600020018054611dcd906142ef565b80601f0160208091040260200160405190810160405280929190818152602001828054611df9906142ef565b8015611e465780601f10611e1b57610100808354040283529160200191611e46565b820191906000526020600020905b815481529060010190602001808311611e2957829003601f168201915b505050505081526020019060010190611dae565b50505091525090949350505050565b6000610c6b60008051602061472883398151915283611932565b600080516020614808833981519152611e9b81612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63e4c22569611ebd61117d565b8a8a8a8a8a8a8a6040518963ffffffff1660e01b8152600401611ee7989796959493929190614329565b60006040518083038186803b158015611eff57600080fd5b505af4158015611f13573d6000803e3d6000fd5b505050505050505050505050565b600080516020614808833981519152611f3981612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63c7bffa96611f5b61117d565b846040518363ffffffff1660e01b8152600401611f79929190614387565b60006040518083038186803b158015611f9157600080fd5b505af4158015611fa5573d6000803e3d6000fd5b505050505050565b6000805160206147a8833981519152611fc581612459565b61117982613105565b60606000611fda612428565b9050610d448160060161244c565b6000805160206147a883398151915261200081612459565b61201261200c8461257d565b8361319e565b8260405161202091906142d3565b604051908190038120907fd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c53137190600090a2505050565b6000806000612062612bef565b6001600160a01b03909416600090815260208581526040808320546001909701909152902054939492505050565b612098612a51565b6000805160206147e88339815191526120b081612459565b6001600160a01b0383166120d757604051633a954ecd60e21b815260040160405180910390fd5b6120ea6001600160a01b038416836128d4565b826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161212591815260200190565b60405180910390a250611179612bb4565b6000805160206147a883398151915261214e81612459565b61215a611ad18361257d565b8160405161216891906142d3565b6040519081900381209060008051602061470883398151915290600090a25050565b6000805160206148088339815191526121a281612459565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63b54e72d86114dd61117d565b60006121ce6131db565b805490915060ff600160401b82041615906001600160401b03166000811580156121f55750825b90506000826001600160401b031660011480156122115750303b155b90508115801561221f575080155b1561223d5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b0319166001178555831561226657845460ff60401b1916600160401b1785555b61226e6131ff565b6122766131ff565b61227e613207565b61228960008861253b565b506122a26000805160206147a88339815191528861253b565b506122bb6000805160206147c88339815191528861253b565b506122d46000805160206147688339815191528761253b565b5060006122df6128b0565b80546001600160a01b03808d166001600160a01b0319928316178355600183018054918d16919092161790559050678ac7230489e80000620151806123248282613217565b505050831561155057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050505050565b60006001600160e01b03198216637965db0b60e01b1480610c6b57506301ffc9a760e01b6001600160e01b0319831614610c6b565b60606123b6611b5d565b6001600160a01b031663306ade09836040518263ffffffff1660e01b81526004016123e391815260200190565b600060405180830381865afa158015612400573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c6b91908101906143a0565b7ff2856e5e1b7689dcde1bb551fd115c3cad8d243ea609d47a46b4d22ee58d300090565b60606000611a028361323a565b61133d8133613296565b600061246d611b5d565b6001600160a01b031663352af39a836040518263ffffffff1660e01b81526004016124989190614071565b602060405180830381865afa1580156124b5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6b919061440d565b60006124e3612428565b905060006124f460098301846132c1565b905080610f4e57604051631a1e056960e01b8152600481018490526024016117b3565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b600080612546612ee8565b9050600061255485856132cd565b9050801561192a576000858152602083905260409020612574908561336e565b50949350505050565b6000612587611b5d565b6001600160a01b0316631ca0e943836040518263ffffffff1660e01b81526004016124989190614071565b6000806125bd612ee8565b905060006125cb8585613383565b9050801561192a57600085815260208390526040902061257490856133fb565b60006125f5612428565b905060006126066009830184613410565b905080610f4e57604051637eae59f160e11b8152600481018490526024016117b3565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061269957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661268d61341c565b6001600160a01b031614155b156126b75760405163703e46dd60e11b815260040160405180910390fd5b565b6000805160206147688339815191526126d181612459565b60006126db611b5d565b6001600160a01b0316639d825bc56040518163ffffffff1660e01b8152600401602060405180830381865afa158015612718573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061273c9190614426565b9050600061274861341c565b9050836001600160a01b0316816001600160a01b0316036127805780846040516382afabc160e01b81526004016117b3929190614443565b816001600160a01b0316846001600160a01b0316146127b657818460405163699a021f60e11b81526004016117b3929190614443565b836001600160a01b0316816001600160a01b03167fa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db460405160405180910390a350505050565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612856575060408051601f3d908101601f191682019092526128539181019061440d565b60015b6128755781604051634c9c8ce360e01b81526004016117b39190613dfe565b60008051602061478883398151915281146128a657604051632a87526960e21b8152600481018290526024016117b3565b610f4e8383613438565b7f0c7b73796c7cc89b9f849b9056a93200eba741881e57a1b03b9bedb2c0e0710090565b804710156128f957478160405163cf47918160e01b81526004016117b3929190613ec1565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114612946576040519150601f19603f3d011682016040523d82523d6000602084013e61294b565b606091505b5050905080610f4e5760405163d6bda27560e01b815260040160405180910390fd5b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146126b75760405163703e46dd60e11b815260040160405180910390fd5b60006129c0612428565b905060006129ce82866132c1565b9050806129f0576040516221e3bb60e31b8152600481018690526024016117b3565b604080518082018252841515815260208082018781526000898152600287018352939093208251815460ff191690151517815592518051929392610ded9260018501920190613860565b600080612a45612428565b9050611a02818461348e565b6000612a5b61349a565b805490915060011901612a8157604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000612a91612bef565b90508060020154821115612ac057806002015482604051631728bc5b60e31b81526004016117b3929190613ec1565b6003810154336000908152602083905260409020544291612ae091614473565b811115612b085733600090815260018301602090815260408083208390559084905290208190555b6002820154336000908152600184016020526040902054612b2a908590614473565b1115612b515781600201548360405163d54b188760e01b81526004016117b3929190613ec1565b33600090815260018301602052604081208054859290612b72908490614473565b90915550612b82905033846128d4565b60405183815233907fb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c290602001611604565b6000612bbe61349a565b6001905550565b612bcf818361348e565b61117957604051631e96f6ed60e21b8152600481018390526024016117b3565b7f99a652063088b6badaeb0c7f680676baf720654b4f86f50167944489af637d0090565b6000612c1d612428565b90506000612c2e600483018461336e565b905080612c505782604051632872fbf960e11b81526004016117b39190613dfe565b6040516001600160a01b038416907fa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f90600090a2505050565b6000612c93612bef565b60028101849055600381018390556040519091507f8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e90612cd69085908590613ec1565b60405180910390a1505050565b6000612ced612428565b9050612cf98382612bc5565b600083815260028201602090815260408220600190810180549182018155835291200161104d83826144e3565b6000612d30612428565b90506000612d4160048301846133fb565b905080612d635782604051631532e67160e21b81526004016117b39190613dfe565b6040516001600160a01b038416907f85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af290600090a2505050565b6000611a02836001600160a01b0384166134be565b6000612dbb612428565b9050612dc78382612bc5565b60008381526002820160205260408120600101905b8154811015612ee15783604051602001612df691906142d3565b60405160208183030381529060405280519060200120828281548110612e1e57612e1e6142bd565b90600052602060002001604051602001612e38919061459c565b6040516020818303038152906040528051906020012003612ed95781548290612e6390600190614612565b81548110612e7357612e736142bd565b90600052602060002001828281548110612e8f57612e8f6142bd565b906000526020600020019081612ea59190614625565b5081805480612eb657612eb66146f1565b600190038181906000526020600020016000612ed291906138b6565b9055612ee1565b600101612ddc565b5050505050565b7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200090565b6000611a0283836134d6565b610f4e83846001600160a01b031663a9059cbb8585604051602401612f3e929190614387565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613500565b6000612f7a612428565b60038101805460ff19168415159081179091556040519081529091507fe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e39060200160405180910390a15050565b6000612fd1612428565b9050612fdd8382612bc5565b60009283526002016020526040909120805460ff1916911515919091179055565b6000613008612428565b905060006130168284613410565b90508061303957604051631e96f6ed60e21b8152600481018490526024016117b3565b60008381526002830160205260408120805460ff1916815590612ee160018301826138f0565b6000610c6b825490565b6000613073612428565b90506000613084600683018561336e565b9050806130a65783604051631a6107e360e31b81526004016117b39190613dfe565b6001600160a01b038416600090815260088301602052604090206130ca84826144e3565b506040516001600160a01b038516907f928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe8290600090a250505050565b600061310f612428565b9050600061312060068301846133fb565b905080613142578260405163ba650b5f60e01b81526004016117b39190613dfe565b6001600160a01b03831660009081526008830160205260408120613165916138b6565b6040516001600160a01b038416907fc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf90600090a2505050565b60006131a8612428565b90506131b48382612bc5565b60008381526002820160209081526040909120835161104d92600190920191850190613860565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b6126b7613568565b61320f613568565b6126b761358d565b61321f613568565b6000613229612bef565b600281019390935550600390910155565b60608160000180548060200260200160405190810160405280929190818152602001828054801561328a57602002820191906000526020600020905b815481526020019060010190808311613276575b50505050509050919050565b6132a08282611932565b61117957808260405163e2517d3f60e01b81526004016117b3929190614387565b6000611a028383613595565b6000806132d8612517565b90506132e48484611932565b613364576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561331a3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610c6b565b6000915050610c6b565b6000611a02836001600160a01b038416613595565b60008061338e612517565b905061339a8484611932565b15613364576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610c6b565b6000611a02836001600160a01b0384166135df565b6000611a0283836135df565b600080516020614788833981519152546001600160a01b031690565b613441826136c8565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561348657610f4e8282613724565b61117961379a565b6000611a0283836134be565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b60009081526001919091016020526040902054151590565b60008260000182815481106134ed576134ed6142bd565b9060005260206000200154905092915050565b600080602060008451602086016000885af180613523576040513d6000823e3d81fd5b50506000513d9150811561353b578060011415613548565b6001600160a01b0384163b155b1561104d5783604051635274afe760e01b81526004016117b39190613dfe565b6135706137b9565b6126b757604051631afcd79f60e31b815260040160405180910390fd5b612bb4613568565b60006135a183836134be565b6135d757508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610c6b565b506000610c6b565b60008181526001830160205260408120548015613364576000613603600183614612565b855490915060009061361790600190614612565b905080821461367c576000866000018281548110613637576136376142bd565b906000526020600020015490508087600001848154811061365a5761365a6142bd565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061368d5761368d6146f1565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610c6b565b806001600160a01b03163b6000036136f55780604051634c9c8ce360e01b81526004016117b39190613dfe565b60008051602061478883398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161374191906142d3565b600060405180830381855af49150503d806000811461377c576040519150601f19603f3d011682016040523d82523d6000602084013e613781565b606091505b50915091506137918583836137d3565b95945050505050565b34156126b75760405163b398979f60e01b815260040160405180910390fd5b60006137c36131db565b54600160401b900460ff16919050565b6060826137e8576137e38261381f565b611a02565b81511580156137ff57506001600160a01b0384163b155b15610d235783604051639996b31560e01b81526004016117b39190613dfe565b80511561382f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b60408051808201909152600081526060602082015290565b8280548282559060005260206000209081019282156138a6579160200282015b828111156138a6578251829061389690826144e3565b5091602001919060010190613880565b506138b292915061390a565b5090565b5080546138c2906142ef565b6000825580601f106138d2575050565b601f01602090049060005260206000209081019061133d9190613927565b508054600082559060005260206000209081019061133d91905b808211156138b257600061391e82826138b6565b5060010161390a565b5b808211156138b25760008155600101613928565b60006020828403121561394e57600080fd5b81356001600160e01b031981168114611a0257600080fd5b60005b83811015613981578181015183820152602001613969565b50506000910152565b600081518084526139a2816020860160208601613966565b601f01601f19169290920160200192915050565b60008282518085526020808601955060208260051b8401016020860160005b84811015613a0357601f198684030189526139f183835161398a565b988401989250908301906001016139d5565b5090979650505050505050565b602081526000611a0260208301846139b6565b6020808252825182820181905260009190848201906040850190845b81811015613a5b57835183529284019291840191600101613a3f565b50909695505050505050565b6001600160a01b038116811461133d57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715613aba57613aba613a7c565b604052919050565b60006001600160401b03821115613adb57613adb613a7c565b50601f01601f191660200190565b600082601f830112613afa57600080fd5b8135613b0d613b0882613ac2565b613a92565b818152846020838601011115613b2257600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215613b5557600080fd5b8435613b6081613a67565b93506020850135613b7081613a67565b92506040850135915060608501356001600160401b03811115613b9257600080fd5b613b9e87828801613ae9565b91505092959194509250565b600080600060608486031215613bbf57600080fd5b8335613bca81613a67565b92506020840135613bda81613a67565b929592945050506040919091013590565b60008060408385031215613bfe57600080fd5b50508035926020909101359150565b600082601f830112613c1e57600080fd5b813560206001600160401b0380831115613c3a57613c3a613a7c565b8260051b613c49838201613a92565b9384528581018301938381019088861115613c6357600080fd5b84880192505b85831015613c9f57823584811115613c815760008081fd5b613c8f8a87838c0101613ae9565b8352509184019190840190613c69565b98975050505050505050565b600060208284031215613cbd57600080fd5b81356001600160401b03811115613cd357600080fd5b61192a84828501613c0d565b600060208284031215613cf157600080fd5b5035919050565b803561ffff81168114613d0a57600080fd5b919050565b600080600060608486031215613d2457600080fd5b83359250613d3460208501613cf8565b9150613d4260408501613cf8565b90509250925092565b60008060408385031215613d5e57600080fd5b823591506020830135613d7081613a67565b809150509250929050565b600060208284031215613d8d57600080fd5b81356001600160401b03811115613da357600080fd5b61192a84828501613ae9565b60008060408385031215613dc257600080fd5b8235613dcd81613a67565b915060208301356001600160401b03811115613de857600080fd5b613df485828601613ae9565b9150509250929050565b6001600160a01b0391909116815260200190565b60008060408385031215613e2557600080fd5b8235613e3081613a67565b946020939093013593505050565b80358015158114613d0a57600080fd5b600080600060608486031215613e6357600080fd5b83356001600160401b0380821115613e7a57600080fd5b613e8687838801613ae9565b9450613e9460208701613e3e565b93506040860135915080821115613eaa57600080fd5b50613eb786828701613c0d565b9150509250925092565b918252602082015260400190565b600060208284031215613ee157600080fd5b8135611a0281613a67565b60008060008060808587031215613f0257600080fd5b8435935060208501359250613f1960408601613cf8565b9150613f2760608601613cf8565b905092959194509250565b60008060408385031215613f4557600080fd5b82356001600160401b0380821115613f5c57600080fd5b613f6886838701613ae9565b93506020850135915080821115613f7e57600080fd5b50613df485828601613ae9565b6000604083018251151584526020808401516040602087015282815180855260608801915060608160051b890101945060208301925060005b81811015613ff257605f19898703018352613fe086855161398a565b95509284019291840191600101613fc4565b5093979650505050505050565b60408152600061401260408301856139b6565b6020838203818501528185518084528284019150828160051b85010183880160005b8381101561406257601f19878403018552614050838351613f8b565b94860194925090850190600101614034565b50909998505050505050505050565b602081526000611a02602083018461398a565b60006020828403121561409657600080fd5b611a0282613e3e565b6020808252825182820181905260009190848201906040850190845b81811015613a5b5783516001600160a01b0316835292840192918401916001016140bb565b600080604083850312156140f357600080fd5b82356001600160401b0381111561410957600080fd5b61411585828601613ae9565b92505061412460208401613e3e565b90509250929050565b60008060006060848603121561414257600080fd5b8335925060208401359150604084013561415b81613a67565b809150509250925092565b602081526000611a026020830184613f8b565b600080600080600080600060e0888a03121561419457600080fd5b873561419f81613a67565b965060208801356001600160401b038111156141ba57600080fd5b6141c68a828b01613ae9565b965050604088013594506060880135935060808801356141e581613a67565b925060a088013591506141fa60c08901613e3e565b905092959891949750929550565b6000806040838503121561421b57600080fd5b82356001600160401b038082111561423257600080fd5b61423e86838701613ae9565b9350602085013591508082111561425457600080fd5b50613df485828601613c0d565b6000806000806080858703121561427757600080fd5b843561428281613a67565b9350602085013561429281613a67565b925060408501356142a281613a67565b915060608501356142b281613a67565b939692955090935050565b634e487b7160e01b600052603260045260246000fd5b600082516142e5818460208701613966565b9190910192915050565b600181811c9082168061430357607f821691505b60208210810361432357634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b0389811682528881166020830152610100604083018190526000916143578483018b61398a565b6060850199909952608084019790975250509290931660a083015260c082015290151560e0909101529392505050565b6001600160a01b03929092168252602082015260400190565b6000602082840312156143b257600080fd5b81516001600160401b038111156143c857600080fd5b8201601f810184136143d957600080fd5b80516143e7613b0882613ac2565b8181528560208385010111156143fc57600080fd5b613791826020830160208601613966565b60006020828403121561441f57600080fd5b5051919050565b60006020828403121561443857600080fd5b8151611a0281613a67565b6001600160a01b0392831681529116602082015260400190565b634e487b7160e01b600052601160045260246000fd5b80820180821115610c6b57610c6b61445d565b601f821115610f4e576000816000526020600020601f850160051c810160208610156144af5750805b601f850160051c820191505b81811015611fa5578281556001016144bb565b600019600383901b1c191660019190911b1790565b81516001600160401b038111156144fc576144fc613a7c565b6145108161450a84546142ef565b84614486565b602080601f83116001811461453f576000841561452d5750858301515b61453785826144ce565b865550611fa5565b600085815260208120601f198616915b8281101561456e5788860151825594840194600190910190840161454f565b508582101561458c5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008083546145aa816142ef565b600182811680156145c257600181146145d757614606565b60ff1984168752821515830287019450614606565b8760005260208060002060005b858110156145fd5781548a8201529084019082016145e4565b50505082870194505b50929695505050505050565b81810381811115610c6b57610c6b61445d565b818103614630575050565b61463a82546142ef565b6001600160401b0381111561465157614651613a7c565b61465f8161450a84546142ef565b6000601f82116001811461468d576000831561467b5750848201545b61468584826144ce565b855550612ee1565b600085815260209020601f19841690600086815260209020845b838110156146c757828601548255600195860195909101906020016146a7565b508583101561458c5793015460001960f8600387901b161c19169092555050600190811b01905550565b634e487b7160e01b600052603160045260246000fdfe52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813e9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621ae562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c95189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9a95e87c5af084bf5db8491c3a6515da9dd6da39b24b0eb0af08d7b9cd808d91c6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e43acdf00ba9ef08b5f2c22768276611b9af078bf6c24fa36b34ec5e9f2eb061faa26469706673582212201d56d6e2c01623da24b81522365664ff931f922846664f9aed25eacfaaecc34064736f6c63430008180033",
}

// CmaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use CmaccountMetaData.ABI instead.
var CmaccountABI = CmaccountMetaData.ABI

// CmaccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CmaccountMetaData.Bin instead.
var CmaccountBin = CmaccountMetaData.Bin

// DeployCmaccount deploys a new Ethereum contract, binding an instance of Cmaccount to it.
func DeployCmaccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cmaccount, error) {
	parsed, err := CmaccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CmaccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cmaccount{CmaccountCaller: CmaccountCaller{contract: contract}, CmaccountTransactor: CmaccountTransactor{contract: contract}, CmaccountFilterer: CmaccountFilterer{contract: contract}}, nil
}

// Cmaccount is an auto generated Go binding around an Ethereum contract.
type Cmaccount struct {
	CmaccountCaller     // Read-only binding to the contract
	CmaccountTransactor // Write-only binding to the contract
	CmaccountFilterer   // Log filterer for contract events
}

// CmaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type CmaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CmaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CmaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CmaccountSession struct {
	Contract     *Cmaccount        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CmaccountCallerSession struct {
	Contract *CmaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CmaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CmaccountTransactorSession struct {
	Contract     *CmaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CmaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type CmaccountRaw struct {
	Contract *Cmaccount // Generic contract binding to access the raw methods on
}

// CmaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CmaccountCallerRaw struct {
	Contract *CmaccountCaller // Generic read-only contract binding to access the raw methods on
}

// CmaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CmaccountTransactorRaw struct {
	Contract *CmaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCmaccount creates a new instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccount(address common.Address, backend bind.ContractBackend) (*Cmaccount, error) {
	contract, err := bindCmaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cmaccount{CmaccountCaller: CmaccountCaller{contract: contract}, CmaccountTransactor: CmaccountTransactor{contract: contract}, CmaccountFilterer: CmaccountFilterer{contract: contract}}, nil
}

// NewCmaccountCaller creates a new read-only instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccountCaller(address common.Address, caller bind.ContractCaller) (*CmaccountCaller, error) {
	contract, err := bindCmaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CmaccountCaller{contract: contract}, nil
}

// NewCmaccountTransactor creates a new write-only instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*CmaccountTransactor, error) {
	contract, err := bindCmaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CmaccountTransactor{contract: contract}, nil
}

// NewCmaccountFilterer creates a new log filterer instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*CmaccountFilterer, error) {
	contract, err := bindCmaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CmaccountFilterer{contract: contract}, nil
}

// bindCmaccount binds a generic wrapper to an already deployed contract.
func bindCmaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CmaccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cmaccount *CmaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cmaccount.Contract.CmaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cmaccount *CmaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.Contract.CmaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cmaccount *CmaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cmaccount.Contract.CmaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cmaccount *CmaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cmaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cmaccount *CmaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cmaccount *CmaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cmaccount.Contract.contract.Transact(opts, method, params...)
}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) BOOKINGOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "BOOKING_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) BOOKINGOPERATORROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOOKINGOPERATORROLE(&_Cmaccount.CallOpts)
}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) BOOKINGOPERATORROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOOKINGOPERATORROLE(&_Cmaccount.CallOpts)
}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) BOTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "BOT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) BOTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOTADMINROLE(&_Cmaccount.CallOpts)
}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) BOTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOTADMINROLE(&_Cmaccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.DEFAULTADMINROLE(&_Cmaccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.DEFAULTADMINROLE(&_Cmaccount.CallOpts)
}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) GASWITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "GAS_WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) GASWITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.GASWITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) GASWITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.GASWITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// MESSENGERBOTROLE is a free data retrieval call binding the contract method 0x72afa328.
//
// Solidity: function MESSENGER_BOT_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) MESSENGERBOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "MESSENGER_BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSENGERBOTROLE is a free data retrieval call binding the contract method 0x72afa328.
//
// Solidity: function MESSENGER_BOT_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) MESSENGERBOTROLE() ([32]byte, error) {
	return _Cmaccount.Contract.MESSENGERBOTROLE(&_Cmaccount.CallOpts)
}

// MESSENGERBOTROLE is a free data retrieval call binding the contract method 0x72afa328.
//
// Solidity: function MESSENGER_BOT_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) MESSENGERBOTROLE() ([32]byte, error) {
	return _Cmaccount.Contract.MESSENGERBOTROLE(&_Cmaccount.CallOpts)
}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) SERVICEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "SERVICE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) SERVICEADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.SERVICEADMINROLE(&_Cmaccount.CallOpts)
}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) SERVICEADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.SERVICEADMINROLE(&_Cmaccount.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) UPGRADERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.UPGRADERROLE(&_Cmaccount.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.UPGRADERROLE(&_Cmaccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccount *CmaccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccount *CmaccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Cmaccount.Contract.UPGRADEINTERFACEVERSION(&_Cmaccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccount *CmaccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Cmaccount.Contract.UPGRADEINTERFACEVERSION(&_Cmaccount.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) WITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.WITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.WITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCaller) GetAllServiceHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getAllServiceHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountSession) GetAllServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetAllServiceHashes(&_Cmaccount.CallOpts)
}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCallerSession) GetAllServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetAllServiceHashes(&_Cmaccount.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccount *CmaccountCaller) GetBookingTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getBookingTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccount *CmaccountSession) GetBookingTokenAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetBookingTokenAddress(&_Cmaccount.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccount *CmaccountCallerSession) GetBookingTokenAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetBookingTokenAddress(&_Cmaccount.CallOpts)
}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Cmaccount *CmaccountCaller) GetGasMoneyWithdrawal(opts *bind.CallOpts) (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getGasMoneyWithdrawal")

	outstruct := new(struct {
		WithdrawalLimit  *big.Int
		WithdrawalPeriod *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawalLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawalPeriod = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Cmaccount *CmaccountSession) GetGasMoneyWithdrawal() (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawal(&_Cmaccount.CallOpts)
}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Cmaccount *CmaccountCallerSession) GetGasMoneyWithdrawal() (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawal(&_Cmaccount.CallOpts)
}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Cmaccount *CmaccountCaller) GetGasMoneyWithdrawalForAccount(opts *bind.CallOpts, account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getGasMoneyWithdrawalForAccount", account)

	outstruct := new(struct {
		PeriodStart     *big.Int
		WithdrawnAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PeriodStart = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawnAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Cmaccount *CmaccountSession) GetGasMoneyWithdrawalForAccount(account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawalForAccount(&_Cmaccount.CallOpts, account)
}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Cmaccount *CmaccountCallerSession) GetGasMoneyWithdrawalForAccount(account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawalForAccount(&_Cmaccount.CallOpts, account)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Cmaccount *CmaccountCaller) GetManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Cmaccount *CmaccountSession) GetManagerAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetManagerAddress(&_Cmaccount.CallOpts)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Cmaccount *CmaccountCallerSession) GetManagerAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetManagerAddress(&_Cmaccount.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Cmaccount *CmaccountCaller) GetPublicKey(opts *bind.CallOpts, pubKeyAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getPublicKey", pubKeyAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Cmaccount *CmaccountSession) GetPublicKey(pubKeyAddress common.Address) ([]byte, error) {
	return _Cmaccount.Contract.GetPublicKey(&_Cmaccount.CallOpts, pubKeyAddress)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Cmaccount *CmaccountCallerSession) GetPublicKey(pubKeyAddress common.Address) ([]byte, error) {
	return _Cmaccount.Contract.GetPublicKey(&_Cmaccount.CallOpts, pubKeyAddress)
}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Cmaccount *CmaccountCaller) GetPublicKeysAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getPublicKeysAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Cmaccount *CmaccountSession) GetPublicKeysAddresses() ([]common.Address, error) {
	return _Cmaccount.Contract.GetPublicKeysAddresses(&_Cmaccount.CallOpts)
}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Cmaccount *CmaccountCallerSession) GetPublicKeysAddresses() ([]common.Address, error) {
	return _Cmaccount.Contract.GetPublicKeysAddresses(&_Cmaccount.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccount *CmaccountCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccount *CmaccountSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cmaccount.Contract.GetRoleAdmin(&_Cmaccount.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cmaccount.Contract.GetRoleAdmin(&_Cmaccount.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccount *CmaccountCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccount *CmaccountSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Cmaccount.Contract.GetRoleMember(&_Cmaccount.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccount *CmaccountCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Cmaccount.Contract.GetRoleMember(&_Cmaccount.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccount *CmaccountCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccount *CmaccountSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetRoleMemberCount(&_Cmaccount.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccount *CmaccountCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetRoleMemberCount(&_Cmaccount.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Cmaccount *CmaccountCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Cmaccount *CmaccountSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Cmaccount.Contract.GetRoleMembers(&_Cmaccount.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Cmaccount *CmaccountCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Cmaccount.Contract.GetRoleMembers(&_Cmaccount.CallOpts, role)
}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((bool,string[]) service)
func (_Cmaccount *CmaccountCaller) GetService(opts *bind.CallOpts, serviceHash [32]byte) (PartnerConfigurationService, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getService", serviceHash)

	if err != nil {
		return *new(PartnerConfigurationService), err
	}

	out0 := *abi.ConvertType(out[0], new(PartnerConfigurationService)).(*PartnerConfigurationService)

	return out0, err

}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((bool,string[]) service)
func (_Cmaccount *CmaccountSession) GetService(serviceHash [32]byte) (PartnerConfigurationService, error) {
	return _Cmaccount.Contract.GetService(&_Cmaccount.CallOpts, serviceHash)
}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((bool,string[]) service)
func (_Cmaccount *CmaccountCallerSession) GetService(serviceHash [32]byte) (PartnerConfigurationService, error) {
	return _Cmaccount.Contract.GetService(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCaller) GetServiceCapabilities(opts *bind.CallOpts, serviceName string) ([]string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceCapabilities", serviceName)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Cmaccount *CmaccountSession) GetServiceCapabilities(serviceName string) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCallerSession) GetServiceCapabilities(serviceName string) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCaller) GetServiceCapabilities0(opts *bind.CallOpts, serviceHash [32]byte) ([]string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceCapabilities0", serviceHash)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Cmaccount *CmaccountSession) GetServiceCapabilities0(serviceHash [32]byte) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities0(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCallerSession) GetServiceCapabilities0(serviceHash [32]byte) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities0(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCaller) GetServiceRestrictedRate(opts *bind.CallOpts, serviceHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceRestrictedRate", serviceHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountSession) GetServiceRestrictedRate(serviceHash [32]byte) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCallerSession) GetServiceRestrictedRate(serviceHash [32]byte) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCaller) GetServiceRestrictedRate0(opts *bind.CallOpts, serviceName string) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceRestrictedRate0", serviceName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountSession) GetServiceRestrictedRate0(serviceName string) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate0(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCallerSession) GetServiceRestrictedRate0(serviceName string) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate0(&_Cmaccount.CallOpts, serviceName)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (bool,string[])[] services)
func (_Cmaccount *CmaccountCaller) GetSupportedServices(opts *bind.CallOpts) (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getSupportedServices")

	outstruct := new(struct {
		ServiceNames []string
		Services     []PartnerConfigurationService
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ServiceNames = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Services = *abi.ConvertType(out[1], new([]PartnerConfigurationService)).(*[]PartnerConfigurationService)

	return *outstruct, err

}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (bool,string[])[] services)
func (_Cmaccount *CmaccountSession) GetSupportedServices() (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	return _Cmaccount.Contract.GetSupportedServices(&_Cmaccount.CallOpts)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (bool,string[])[] services)
func (_Cmaccount *CmaccountCallerSession) GetSupportedServices() (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	return _Cmaccount.Contract.GetSupportedServices(&_Cmaccount.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Cmaccount *CmaccountCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Cmaccount *CmaccountSession) GetSupportedTokens() ([]common.Address, error) {
	return _Cmaccount.Contract.GetSupportedTokens(&_Cmaccount.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Cmaccount *CmaccountCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _Cmaccount.Contract.GetSupportedTokens(&_Cmaccount.CallOpts)
}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCaller) GetWantedServiceHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getWantedServiceHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountSession) GetWantedServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetWantedServiceHashes(&_Cmaccount.CallOpts)
}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCallerSession) GetWantedServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetWantedServiceHashes(&_Cmaccount.CallOpts)
}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Cmaccount *CmaccountCaller) GetWantedServices(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getWantedServices")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Cmaccount *CmaccountSession) GetWantedServices() ([]string, error) {
	return _Cmaccount.Contract.GetWantedServices(&_Cmaccount.CallOpts)
}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Cmaccount *CmaccountCallerSession) GetWantedServices() ([]string, error) {
	return _Cmaccount.Contract.GetWantedServices(&_Cmaccount.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccount *CmaccountCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccount *CmaccountSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cmaccount.Contract.HasRole(&_Cmaccount.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccount *CmaccountCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cmaccount.Contract.HasRole(&_Cmaccount.CallOpts, role, account)
}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Cmaccount *CmaccountCaller) IsBotAllowed(opts *bind.CallOpts, bot common.Address) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "isBotAllowed", bot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Cmaccount *CmaccountSession) IsBotAllowed(bot common.Address) (bool, error) {
	return _Cmaccount.Contract.IsBotAllowed(&_Cmaccount.CallOpts, bot)
}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Cmaccount *CmaccountCallerSession) IsBotAllowed(bot common.Address) (bool, error) {
	return _Cmaccount.Contract.IsBotAllowed(&_Cmaccount.CallOpts, bot)
}

// IsServiceSupported is a free data retrieval call binding the contract method 0x58c0a4c2.
//
// Solidity: function isServiceSupported(string serviceName) view returns(bool)
func (_Cmaccount *CmaccountCaller) IsServiceSupported(opts *bind.CallOpts, serviceName string) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "isServiceSupported", serviceName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsServiceSupported is a free data retrieval call binding the contract method 0x58c0a4c2.
//
// Solidity: function isServiceSupported(string serviceName) view returns(bool)
func (_Cmaccount *CmaccountSession) IsServiceSupported(serviceName string) (bool, error) {
	return _Cmaccount.Contract.IsServiceSupported(&_Cmaccount.CallOpts, serviceName)
}

// IsServiceSupported is a free data retrieval call binding the contract method 0x58c0a4c2.
//
// Solidity: function isServiceSupported(string serviceName) view returns(bool)
func (_Cmaccount *CmaccountCallerSession) IsServiceSupported(serviceName string) (bool, error) {
	return _Cmaccount.Contract.IsServiceSupported(&_Cmaccount.CallOpts, serviceName)
}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Cmaccount *CmaccountCaller) OffChainPaymentSupported(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "offChainPaymentSupported")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Cmaccount *CmaccountSession) OffChainPaymentSupported() (bool, error) {
	return _Cmaccount.Contract.OffChainPaymentSupported(&_Cmaccount.CallOpts)
}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Cmaccount *CmaccountCallerSession) OffChainPaymentSupported() (bool, error) {
	return _Cmaccount.Contract.OffChainPaymentSupported(&_Cmaccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccount *CmaccountSession) ProxiableUUID() ([32]byte, error) {
	return _Cmaccount.Contract.ProxiableUUID(&_Cmaccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Cmaccount.Contract.ProxiableUUID(&_Cmaccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccount *CmaccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccount *CmaccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cmaccount.Contract.SupportsInterface(&_Cmaccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccount *CmaccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cmaccount.Contract.SupportsInterface(&_Cmaccount.CallOpts, interfaceId)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactor) AcceptCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "acceptCancellation", tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AcceptCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactorSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AcceptCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Cmaccount *CmaccountTransactor) AddMessengerBot(opts *bind.TransactOpts, bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addMessengerBot", bot, gasMoney)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Cmaccount *CmaccountSession) AddMessengerBot(bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddMessengerBot(&_Cmaccount.TransactOpts, bot, gasMoney)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Cmaccount *CmaccountTransactorSession) AddMessengerBot(bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddMessengerBot(&_Cmaccount.TransactOpts, bot, gasMoney)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Cmaccount *CmaccountTransactor) AddPublicKey(opts *bind.TransactOpts, pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addPublicKey", pubKeyAddress, data)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Cmaccount *CmaccountSession) AddPublicKey(pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddPublicKey(&_Cmaccount.TransactOpts, pubKeyAddress, data)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Cmaccount *CmaccountTransactorSession) AddPublicKey(pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddPublicKey(&_Cmaccount.TransactOpts, pubKeyAddress, data)
}

// AddService is a paid mutator transaction binding the contract method 0x581ed290.
//
// Solidity: function addService(string serviceName, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactor) AddService(opts *bind.TransactOpts, serviceName string, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addService", serviceName, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x581ed290.
//
// Solidity: function addService(string serviceName, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountSession) AddService(serviceName string, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddService(&_Cmaccount.TransactOpts, serviceName, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x581ed290.
//
// Solidity: function addService(string serviceName, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactorSession) AddService(serviceName string, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddService(&_Cmaccount.TransactOpts, serviceName, restrictedRate, capabilities)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactor) AddServiceCapability(opts *bind.TransactOpts, serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addServiceCapability", serviceName, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountSession) AddServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactorSession) AddServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactor) AddSupportedToken(opts *bind.TransactOpts, _supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addSupportedToken", _supportedToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountSession) AddSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactorSession) AddSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactor) AddWantedServices(opts *bind.TransactOpts, serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addWantedServices", serviceNames)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountSession) AddWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactorSession) AddWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Cmaccount *CmaccountTransactor) BuyBookingToken(opts *bind.TransactOpts, tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "buyBookingToken", tokenId, expectedPrice, expectedPaymentToken)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Cmaccount *CmaccountSession) BuyBookingToken(tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.BuyBookingToken(&_Cmaccount.TransactOpts, tokenId, expectedPrice, expectedPaymentToken)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Cmaccount *CmaccountTransactorSession) BuyBookingToken(tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.BuyBookingToken(&_Cmaccount.TransactOpts, tokenId, expectedPrice, expectedPaymentToken)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) CounterCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "counterCancellation", tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Cmaccount *CmaccountSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.CounterCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.CounterCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactor) FinalizeCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "finalizeCancellation", tokenId, refundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountSession) FinalizeCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.FinalizeCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactorSession) FinalizeCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.FinalizeCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.GrantRole(&_Cmaccount.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.GrantRole(&_Cmaccount.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Cmaccount *CmaccountTransactor) Initialize(opts *bind.TransactOpts, manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "initialize", manager, bookingToken, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Cmaccount *CmaccountSession) Initialize(manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.Initialize(&_Cmaccount.TransactOpts, manager, bookingToken, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Cmaccount *CmaccountTransactorSession) Initialize(manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.Initialize(&_Cmaccount.TransactOpts, manager, bookingToken, defaultAdmin, upgrader)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) InitiateCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "initiateCancellation", tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Cmaccount *CmaccountSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.InitiateCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.InitiateCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Cmaccount *CmaccountTransactor) MintBookingToken(opts *bind.TransactOpts, reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "mintBookingToken", reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Cmaccount *CmaccountSession) MintBookingToken(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.MintBookingToken(&_Cmaccount.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Cmaccount *CmaccountTransactorSession) MintBookingToken(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.MintBookingToken(&_Cmaccount.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Cmaccount *CmaccountTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Cmaccount *CmaccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.OnERC721Received(&_Cmaccount.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Cmaccount *CmaccountTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.OnERC721Received(&_Cmaccount.TransactOpts, arg0, arg1, arg2, arg3)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactor) RecordExpiration(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "recordExpiration", tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Cmaccount *CmaccountSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.RecordExpiration(&_Cmaccount.TransactOpts, tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactorSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.RecordExpiration(&_Cmaccount.TransactOpts, tokenId)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) RejectCancellation(opts *bind.TransactOpts, tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "rejectCancellation", tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Cmaccount *CmaccountSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.RejectCancellation(&_Cmaccount.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.RejectCancellation(&_Cmaccount.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Cmaccount *CmaccountTransactor) RemoveAllServices(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeAllServices")
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Cmaccount *CmaccountSession) RemoveAllServices() (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveAllServices(&_Cmaccount.TransactOpts)
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveAllServices() (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveAllServices(&_Cmaccount.TransactOpts)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Cmaccount *CmaccountTransactor) RemoveMessengerBot(opts *bind.TransactOpts, bot common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeMessengerBot", bot)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Cmaccount *CmaccountSession) RemoveMessengerBot(bot common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveMessengerBot(&_Cmaccount.TransactOpts, bot)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveMessengerBot(bot common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveMessengerBot(&_Cmaccount.TransactOpts, bot)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Cmaccount *CmaccountTransactor) RemovePublicKey(opts *bind.TransactOpts, pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removePublicKey", pubKeyAddress)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Cmaccount *CmaccountSession) RemovePublicKey(pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemovePublicKey(&_Cmaccount.TransactOpts, pubKeyAddress)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Cmaccount *CmaccountTransactorSession) RemovePublicKey(pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemovePublicKey(&_Cmaccount.TransactOpts, pubKeyAddress)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Cmaccount *CmaccountTransactor) RemoveService(opts *bind.TransactOpts, serviceName string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeService", serviceName)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Cmaccount *CmaccountSession) RemoveService(serviceName string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveService(&_Cmaccount.TransactOpts, serviceName)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveService(serviceName string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveService(&_Cmaccount.TransactOpts, serviceName)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactor) RemoveServiceCapability(opts *bind.TransactOpts, serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeServiceCapability", serviceName, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountSession) RemoveServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactor) RemoveSupportedToken(opts *bind.TransactOpts, _supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeSupportedToken", _supportedToken)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountSession) RemoveSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactor) RemoveWantedServices(opts *bind.TransactOpts, serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeWantedServices", serviceNames)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountSession) RemoveWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccount *CmaccountTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccount *CmaccountSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RenounceRole(&_Cmaccount.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccount *CmaccountTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RenounceRole(&_Cmaccount.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RevokeRole(&_Cmaccount.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RevokeRole(&_Cmaccount.TransactOpts, role, account)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Cmaccount *CmaccountTransactor) SetGasMoneyWithdrawal(opts *bind.TransactOpts, limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setGasMoneyWithdrawal", limit, period)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Cmaccount *CmaccountSession) SetGasMoneyWithdrawal(limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetGasMoneyWithdrawal(&_Cmaccount.TransactOpts, limit, period)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Cmaccount *CmaccountTransactorSession) SetGasMoneyWithdrawal(limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetGasMoneyWithdrawal(&_Cmaccount.TransactOpts, limit, period)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Cmaccount *CmaccountTransactor) SetOffChainPaymentSupported(opts *bind.TransactOpts, _isSupported bool) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setOffChainPaymentSupported", _isSupported)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Cmaccount *CmaccountSession) SetOffChainPaymentSupported(_isSupported bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetOffChainPaymentSupported(&_Cmaccount.TransactOpts, _isSupported)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Cmaccount *CmaccountTransactorSession) SetOffChainPaymentSupported(_isSupported bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetOffChainPaymentSupported(&_Cmaccount.TransactOpts, _isSupported)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactor) SetServiceCapabilities(opts *bind.TransactOpts, serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setServiceCapabilities", serviceName, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Cmaccount *CmaccountSession) SetServiceCapabilities(serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceCapabilities(&_Cmaccount.TransactOpts, serviceName, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactorSession) SetServiceCapabilities(serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceCapabilities(&_Cmaccount.TransactOpts, serviceName, capabilities)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Cmaccount *CmaccountTransactor) SetServiceRestrictedRate(opts *bind.TransactOpts, serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setServiceRestrictedRate", serviceName, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Cmaccount *CmaccountSession) SetServiceRestrictedRate(serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceRestrictedRate(&_Cmaccount.TransactOpts, serviceName, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Cmaccount *CmaccountTransactorSession) SetServiceRestrictedRate(serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceRestrictedRate(&_Cmaccount.TransactOpts, serviceName, restrictedRate)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactor) TransferERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "transferERC20", token, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Cmaccount *CmaccountSession) TransferERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC20(&_Cmaccount.TransactOpts, token, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactorSession) TransferERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC20(&_Cmaccount.TransactOpts, token, to, amount)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactor) TransferERC721(opts *bind.TransactOpts, token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "transferERC721", token, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Cmaccount *CmaccountSession) TransferERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC721(&_Cmaccount.TransactOpts, token, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactorSession) TransferERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC721(&_Cmaccount.TransactOpts, token, to, tokenId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccount *CmaccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccount *CmaccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.UpgradeToAndCall(&_Cmaccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccount *CmaccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.UpgradeToAndCall(&_Cmaccount.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "withdraw", recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Cmaccount *CmaccountSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.Withdraw(&_Cmaccount.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.Withdraw(&_Cmaccount.TransactOpts, recipient, amount)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) WithdrawCancellation(opts *bind.TransactOpts, tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "withdrawCancellation", tokenId, reason, reasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Cmaccount *CmaccountSession) WithdrawCancellation(tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawCancellation(&_Cmaccount.TransactOpts, tokenId, reason, reasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) WithdrawCancellation(tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawCancellation(&_Cmaccount.TransactOpts, tokenId, reason, reasonVersion)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Cmaccount *CmaccountTransactor) WithdrawGasMoney(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "withdrawGasMoney", amount)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Cmaccount *CmaccountSession) WithdrawGasMoney(amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawGasMoney(&_Cmaccount.TransactOpts, amount)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Cmaccount *CmaccountTransactorSession) WithdrawGasMoney(amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawGasMoney(&_Cmaccount.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cmaccount *CmaccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cmaccount *CmaccountSession) Receive() (*types.Transaction, error) {
	return _Cmaccount.Contract.Receive(&_Cmaccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cmaccount *CmaccountTransactorSession) Receive() (*types.Transaction, error) {
	return _Cmaccount.Contract.Receive(&_Cmaccount.TransactOpts)
}

// CmaccountCMAccountUpgradedIterator is returned from FilterCMAccountUpgraded and is used to iterate over the raw logs and unpacked data for CMAccountUpgraded events raised by the Cmaccount contract.
type CmaccountCMAccountUpgradedIterator struct {
	Event *CmaccountCMAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *CmaccountCMAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountCMAccountUpgraded)
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
		it.Event = new(CmaccountCMAccountUpgraded)
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
func (it *CmaccountCMAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountCMAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountCMAccountUpgraded represents a CMAccountUpgraded event raised by the Cmaccount contract.
type CmaccountCMAccountUpgraded struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCMAccountUpgraded is a free log retrieval operation binding the contract event 0xa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db4.
//
// Solidity: event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccount *CmaccountFilterer) FilterCMAccountUpgraded(opts *bind.FilterOpts, oldImplementation []common.Address, newImplementation []common.Address) (*CmaccountCMAccountUpgradedIterator, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "CMAccountUpgraded", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountCMAccountUpgradedIterator{contract: _Cmaccount.contract, event: "CMAccountUpgraded", logs: logs, sub: sub}, nil
}

// WatchCMAccountUpgraded is a free log subscription operation binding the contract event 0xa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db4.
//
// Solidity: event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccount *CmaccountFilterer) WatchCMAccountUpgraded(opts *bind.WatchOpts, sink chan<- *CmaccountCMAccountUpgraded, oldImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "CMAccountUpgraded", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountCMAccountUpgraded)
				if err := _Cmaccount.contract.UnpackLog(event, "CMAccountUpgraded", log); err != nil {
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

// ParseCMAccountUpgraded is a log parse operation binding the contract event 0xa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db4.
//
// Solidity: event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccount *CmaccountFilterer) ParseCMAccountUpgraded(log types.Log) (*CmaccountCMAccountUpgraded, error) {
	event := new(CmaccountCMAccountUpgraded)
	if err := _Cmaccount.contract.UnpackLog(event, "CMAccountUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Cmaccount contract.
type CmaccountDepositIterator struct {
	Event *CmaccountDeposit // Event containing the contract specifics and raw log

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
func (it *CmaccountDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountDeposit)
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
		it.Event = new(CmaccountDeposit)
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
func (it *CmaccountDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountDeposit represents a Deposit event raised by the Cmaccount contract.
type CmaccountDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Cmaccount *CmaccountFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*CmaccountDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountDepositIterator{contract: _Cmaccount.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Cmaccount *CmaccountFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CmaccountDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountDeposit)
				if err := _Cmaccount.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Cmaccount *CmaccountFilterer) ParseDeposit(log types.Log) (*CmaccountDeposit, error) {
	event := new(CmaccountDeposit)
	if err := _Cmaccount.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountGasMoneyWithdrawalIterator is returned from FilterGasMoneyWithdrawal and is used to iterate over the raw logs and unpacked data for GasMoneyWithdrawal events raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawalIterator struct {
	Event *CmaccountGasMoneyWithdrawal // Event containing the contract specifics and raw log

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
func (it *CmaccountGasMoneyWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountGasMoneyWithdrawal)
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
		it.Event = new(CmaccountGasMoneyWithdrawal)
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
func (it *CmaccountGasMoneyWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountGasMoneyWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountGasMoneyWithdrawal represents a GasMoneyWithdrawal event raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawal struct {
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasMoneyWithdrawal is a free log retrieval operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Cmaccount *CmaccountFilterer) FilterGasMoneyWithdrawal(opts *bind.FilterOpts, withdrawer []common.Address) (*CmaccountGasMoneyWithdrawalIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "GasMoneyWithdrawal", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountGasMoneyWithdrawalIterator{contract: _Cmaccount.contract, event: "GasMoneyWithdrawal", logs: logs, sub: sub}, nil
}

// WatchGasMoneyWithdrawal is a free log subscription operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Cmaccount *CmaccountFilterer) WatchGasMoneyWithdrawal(opts *bind.WatchOpts, sink chan<- *CmaccountGasMoneyWithdrawal, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "GasMoneyWithdrawal", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountGasMoneyWithdrawal)
				if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawal", log); err != nil {
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

// ParseGasMoneyWithdrawal is a log parse operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Cmaccount *CmaccountFilterer) ParseGasMoneyWithdrawal(log types.Log) (*CmaccountGasMoneyWithdrawal, error) {
	event := new(CmaccountGasMoneyWithdrawal)
	if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountGasMoneyWithdrawalUpdatedIterator is returned from FilterGasMoneyWithdrawalUpdated and is used to iterate over the raw logs and unpacked data for GasMoneyWithdrawalUpdated events raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawalUpdatedIterator struct {
	Event *CmaccountGasMoneyWithdrawalUpdated // Event containing the contract specifics and raw log

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
func (it *CmaccountGasMoneyWithdrawalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountGasMoneyWithdrawalUpdated)
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
		it.Event = new(CmaccountGasMoneyWithdrawalUpdated)
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
func (it *CmaccountGasMoneyWithdrawalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountGasMoneyWithdrawalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountGasMoneyWithdrawalUpdated represents a GasMoneyWithdrawalUpdated event raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawalUpdated struct {
	Limit  *big.Int
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGasMoneyWithdrawalUpdated is a free log retrieval operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Cmaccount *CmaccountFilterer) FilterGasMoneyWithdrawalUpdated(opts *bind.FilterOpts) (*CmaccountGasMoneyWithdrawalUpdatedIterator, error) {

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "GasMoneyWithdrawalUpdated")
	if err != nil {
		return nil, err
	}
	return &CmaccountGasMoneyWithdrawalUpdatedIterator{contract: _Cmaccount.contract, event: "GasMoneyWithdrawalUpdated", logs: logs, sub: sub}, nil
}

// WatchGasMoneyWithdrawalUpdated is a free log subscription operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Cmaccount *CmaccountFilterer) WatchGasMoneyWithdrawalUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountGasMoneyWithdrawalUpdated) (event.Subscription, error) {

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "GasMoneyWithdrawalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountGasMoneyWithdrawalUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawalUpdated", log); err != nil {
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

// ParseGasMoneyWithdrawalUpdated is a log parse operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Cmaccount *CmaccountFilterer) ParseGasMoneyWithdrawalUpdated(log types.Log) (*CmaccountGasMoneyWithdrawalUpdated, error) {
	event := new(CmaccountGasMoneyWithdrawalUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cmaccount contract.
type CmaccountInitializedIterator struct {
	Event *CmaccountInitialized // Event containing the contract specifics and raw log

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
func (it *CmaccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountInitialized)
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
		it.Event = new(CmaccountInitialized)
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
func (it *CmaccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountInitialized represents a Initialized event raised by the Cmaccount contract.
type CmaccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccount *CmaccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*CmaccountInitializedIterator, error) {

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CmaccountInitializedIterator{contract: _Cmaccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccount *CmaccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CmaccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountInitialized)
				if err := _Cmaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccount *CmaccountFilterer) ParseInitialized(log types.Log) (*CmaccountInitialized, error) {
	event := new(CmaccountInitialized)
	if err := _Cmaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountMessengerBotAddedIterator is returned from FilterMessengerBotAdded and is used to iterate over the raw logs and unpacked data for MessengerBotAdded events raised by the Cmaccount contract.
type CmaccountMessengerBotAddedIterator struct {
	Event *CmaccountMessengerBotAdded // Event containing the contract specifics and raw log

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
func (it *CmaccountMessengerBotAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountMessengerBotAdded)
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
		it.Event = new(CmaccountMessengerBotAdded)
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
func (it *CmaccountMessengerBotAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountMessengerBotAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountMessengerBotAdded represents a MessengerBotAdded event raised by the Cmaccount contract.
type CmaccountMessengerBotAdded struct {
	Bot common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMessengerBotAdded is a free log retrieval operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Cmaccount *CmaccountFilterer) FilterMessengerBotAdded(opts *bind.FilterOpts, bot []common.Address) (*CmaccountMessengerBotAddedIterator, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "MessengerBotAdded", botRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountMessengerBotAddedIterator{contract: _Cmaccount.contract, event: "MessengerBotAdded", logs: logs, sub: sub}, nil
}

// WatchMessengerBotAdded is a free log subscription operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Cmaccount *CmaccountFilterer) WatchMessengerBotAdded(opts *bind.WatchOpts, sink chan<- *CmaccountMessengerBotAdded, bot []common.Address) (event.Subscription, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "MessengerBotAdded", botRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountMessengerBotAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotAdded", log); err != nil {
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

// ParseMessengerBotAdded is a log parse operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Cmaccount *CmaccountFilterer) ParseMessengerBotAdded(log types.Log) (*CmaccountMessengerBotAdded, error) {
	event := new(CmaccountMessengerBotAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountMessengerBotRemovedIterator is returned from FilterMessengerBotRemoved and is used to iterate over the raw logs and unpacked data for MessengerBotRemoved events raised by the Cmaccount contract.
type CmaccountMessengerBotRemovedIterator struct {
	Event *CmaccountMessengerBotRemoved // Event containing the contract specifics and raw log

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
func (it *CmaccountMessengerBotRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountMessengerBotRemoved)
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
		it.Event = new(CmaccountMessengerBotRemoved)
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
func (it *CmaccountMessengerBotRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountMessengerBotRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountMessengerBotRemoved represents a MessengerBotRemoved event raised by the Cmaccount contract.
type CmaccountMessengerBotRemoved struct {
	Bot common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMessengerBotRemoved is a free log retrieval operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Cmaccount *CmaccountFilterer) FilterMessengerBotRemoved(opts *bind.FilterOpts, bot []common.Address) (*CmaccountMessengerBotRemovedIterator, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "MessengerBotRemoved", botRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountMessengerBotRemovedIterator{contract: _Cmaccount.contract, event: "MessengerBotRemoved", logs: logs, sub: sub}, nil
}

// WatchMessengerBotRemoved is a free log subscription operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Cmaccount *CmaccountFilterer) WatchMessengerBotRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountMessengerBotRemoved, bot []common.Address) (event.Subscription, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "MessengerBotRemoved", botRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountMessengerBotRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotRemoved", log); err != nil {
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

// ParseMessengerBotRemoved is a log parse operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Cmaccount *CmaccountFilterer) ParseMessengerBotRemoved(log types.Log) (*CmaccountMessengerBotRemoved, error) {
	event := new(CmaccountMessengerBotRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountOffChainPaymentSupportUpdatedIterator is returned from FilterOffChainPaymentSupportUpdated and is used to iterate over the raw logs and unpacked data for OffChainPaymentSupportUpdated events raised by the Cmaccount contract.
type CmaccountOffChainPaymentSupportUpdatedIterator struct {
	Event *CmaccountOffChainPaymentSupportUpdated // Event containing the contract specifics and raw log

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
func (it *CmaccountOffChainPaymentSupportUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountOffChainPaymentSupportUpdated)
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
		it.Event = new(CmaccountOffChainPaymentSupportUpdated)
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
func (it *CmaccountOffChainPaymentSupportUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountOffChainPaymentSupportUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountOffChainPaymentSupportUpdated represents a OffChainPaymentSupportUpdated event raised by the Cmaccount contract.
type CmaccountOffChainPaymentSupportUpdated struct {
	SupportsOffChainPayment bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOffChainPaymentSupportUpdated is a free log retrieval operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Cmaccount *CmaccountFilterer) FilterOffChainPaymentSupportUpdated(opts *bind.FilterOpts) (*CmaccountOffChainPaymentSupportUpdatedIterator, error) {

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "OffChainPaymentSupportUpdated")
	if err != nil {
		return nil, err
	}
	return &CmaccountOffChainPaymentSupportUpdatedIterator{contract: _Cmaccount.contract, event: "OffChainPaymentSupportUpdated", logs: logs, sub: sub}, nil
}

// WatchOffChainPaymentSupportUpdated is a free log subscription operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Cmaccount *CmaccountFilterer) WatchOffChainPaymentSupportUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountOffChainPaymentSupportUpdated) (event.Subscription, error) {

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "OffChainPaymentSupportUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountOffChainPaymentSupportUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "OffChainPaymentSupportUpdated", log); err != nil {
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

// ParseOffChainPaymentSupportUpdated is a log parse operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Cmaccount *CmaccountFilterer) ParseOffChainPaymentSupportUpdated(log types.Log) (*CmaccountOffChainPaymentSupportUpdated, error) {
	event := new(CmaccountOffChainPaymentSupportUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "OffChainPaymentSupportUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPaymentTokenAddedIterator is returned from FilterPaymentTokenAdded and is used to iterate over the raw logs and unpacked data for PaymentTokenAdded events raised by the Cmaccount contract.
type CmaccountPaymentTokenAddedIterator struct {
	Event *CmaccountPaymentTokenAdded // Event containing the contract specifics and raw log

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
func (it *CmaccountPaymentTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPaymentTokenAdded)
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
		it.Event = new(CmaccountPaymentTokenAdded)
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
func (it *CmaccountPaymentTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPaymentTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPaymentTokenAdded represents a PaymentTokenAdded event raised by the Cmaccount contract.
type CmaccountPaymentTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenAdded is a free log retrieval operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Cmaccount *CmaccountFilterer) FilterPaymentTokenAdded(opts *bind.FilterOpts, token []common.Address) (*CmaccountPaymentTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PaymentTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPaymentTokenAddedIterator{contract: _Cmaccount.contract, event: "PaymentTokenAdded", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenAdded is a free log subscription operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Cmaccount *CmaccountFilterer) WatchPaymentTokenAdded(opts *bind.WatchOpts, sink chan<- *CmaccountPaymentTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PaymentTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPaymentTokenAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenAdded", log); err != nil {
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

// ParsePaymentTokenAdded is a log parse operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Cmaccount *CmaccountFilterer) ParsePaymentTokenAdded(log types.Log) (*CmaccountPaymentTokenAdded, error) {
	event := new(CmaccountPaymentTokenAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPaymentTokenRemovedIterator is returned from FilterPaymentTokenRemoved and is used to iterate over the raw logs and unpacked data for PaymentTokenRemoved events raised by the Cmaccount contract.
type CmaccountPaymentTokenRemovedIterator struct {
	Event *CmaccountPaymentTokenRemoved // Event containing the contract specifics and raw log

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
func (it *CmaccountPaymentTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPaymentTokenRemoved)
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
		it.Event = new(CmaccountPaymentTokenRemoved)
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
func (it *CmaccountPaymentTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPaymentTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPaymentTokenRemoved represents a PaymentTokenRemoved event raised by the Cmaccount contract.
type CmaccountPaymentTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenRemoved is a free log retrieval operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Cmaccount *CmaccountFilterer) FilterPaymentTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*CmaccountPaymentTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PaymentTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPaymentTokenRemovedIterator{contract: _Cmaccount.contract, event: "PaymentTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenRemoved is a free log subscription operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Cmaccount *CmaccountFilterer) WatchPaymentTokenRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountPaymentTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PaymentTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPaymentTokenRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenRemoved", log); err != nil {
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

// ParsePaymentTokenRemoved is a log parse operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Cmaccount *CmaccountFilterer) ParsePaymentTokenRemoved(log types.Log) (*CmaccountPaymentTokenRemoved, error) {
	event := new(CmaccountPaymentTokenRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPublicKeyAddedIterator is returned from FilterPublicKeyAdded and is used to iterate over the raw logs and unpacked data for PublicKeyAdded events raised by the Cmaccount contract.
type CmaccountPublicKeyAddedIterator struct {
	Event *CmaccountPublicKeyAdded // Event containing the contract specifics and raw log

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
func (it *CmaccountPublicKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPublicKeyAdded)
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
		it.Event = new(CmaccountPublicKeyAdded)
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
func (it *CmaccountPublicKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPublicKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPublicKeyAdded represents a PublicKeyAdded event raised by the Cmaccount contract.
type CmaccountPublicKeyAdded struct {
	PubKeyAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyAdded is a free log retrieval operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) FilterPublicKeyAdded(opts *bind.FilterOpts, pubKeyAddress []common.Address) (*CmaccountPublicKeyAddedIterator, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PublicKeyAdded", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPublicKeyAddedIterator{contract: _Cmaccount.contract, event: "PublicKeyAdded", logs: logs, sub: sub}, nil
}

// WatchPublicKeyAdded is a free log subscription operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) WatchPublicKeyAdded(opts *bind.WatchOpts, sink chan<- *CmaccountPublicKeyAdded, pubKeyAddress []common.Address) (event.Subscription, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PublicKeyAdded", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPublicKeyAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyAdded", log); err != nil {
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

// ParsePublicKeyAdded is a log parse operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) ParsePublicKeyAdded(log types.Log) (*CmaccountPublicKeyAdded, error) {
	event := new(CmaccountPublicKeyAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPublicKeyRemovedIterator is returned from FilterPublicKeyRemoved and is used to iterate over the raw logs and unpacked data for PublicKeyRemoved events raised by the Cmaccount contract.
type CmaccountPublicKeyRemovedIterator struct {
	Event *CmaccountPublicKeyRemoved // Event containing the contract specifics and raw log

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
func (it *CmaccountPublicKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPublicKeyRemoved)
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
		it.Event = new(CmaccountPublicKeyRemoved)
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
func (it *CmaccountPublicKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPublicKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPublicKeyRemoved represents a PublicKeyRemoved event raised by the Cmaccount contract.
type CmaccountPublicKeyRemoved struct {
	PubKeyAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyRemoved is a free log retrieval operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) FilterPublicKeyRemoved(opts *bind.FilterOpts, pubKeyAddress []common.Address) (*CmaccountPublicKeyRemovedIterator, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PublicKeyRemoved", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPublicKeyRemovedIterator{contract: _Cmaccount.contract, event: "PublicKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchPublicKeyRemoved is a free log subscription operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) WatchPublicKeyRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountPublicKeyRemoved, pubKeyAddress []common.Address) (event.Subscription, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PublicKeyRemoved", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPublicKeyRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyRemoved", log); err != nil {
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

// ParsePublicKeyRemoved is a log parse operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) ParsePublicKeyRemoved(log types.Log) (*CmaccountPublicKeyRemoved, error) {
	event := new(CmaccountPublicKeyRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Cmaccount contract.
type CmaccountRoleAdminChangedIterator struct {
	Event *CmaccountRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CmaccountRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountRoleAdminChanged)
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
		it.Event = new(CmaccountRoleAdminChanged)
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
func (it *CmaccountRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountRoleAdminChanged represents a RoleAdminChanged event raised by the Cmaccount contract.
type CmaccountRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccount *CmaccountFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CmaccountRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountRoleAdminChangedIterator{contract: _Cmaccount.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccount *CmaccountFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CmaccountRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountRoleAdminChanged)
				if err := _Cmaccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccount *CmaccountFilterer) ParseRoleAdminChanged(log types.Log) (*CmaccountRoleAdminChanged, error) {
	event := new(CmaccountRoleAdminChanged)
	if err := _Cmaccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Cmaccount contract.
type CmaccountRoleGrantedIterator struct {
	Event *CmaccountRoleGranted // Event containing the contract specifics and raw log

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
func (it *CmaccountRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountRoleGranted)
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
		it.Event = new(CmaccountRoleGranted)
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
func (it *CmaccountRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountRoleGranted represents a RoleGranted event raised by the Cmaccount contract.
type CmaccountRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CmaccountRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountRoleGrantedIterator{contract: _Cmaccount.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CmaccountRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountRoleGranted)
				if err := _Cmaccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) ParseRoleGranted(log types.Log) (*CmaccountRoleGranted, error) {
	event := new(CmaccountRoleGranted)
	if err := _Cmaccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Cmaccount contract.
type CmaccountRoleRevokedIterator struct {
	Event *CmaccountRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CmaccountRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountRoleRevoked)
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
		it.Event = new(CmaccountRoleRevoked)
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
func (it *CmaccountRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountRoleRevoked represents a RoleRevoked event raised by the Cmaccount contract.
type CmaccountRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CmaccountRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountRoleRevokedIterator{contract: _Cmaccount.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CmaccountRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountRoleRevoked)
				if err := _Cmaccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) ParseRoleRevoked(log types.Log) (*CmaccountRoleRevoked, error) {
	event := new(CmaccountRoleRevoked)
	if err := _Cmaccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceAddedIterator is returned from FilterServiceAdded and is used to iterate over the raw logs and unpacked data for ServiceAdded events raised by the Cmaccount contract.
type CmaccountServiceAddedIterator struct {
	Event *CmaccountServiceAdded // Event containing the contract specifics and raw log

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
func (it *CmaccountServiceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceAdded)
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
		it.Event = new(CmaccountServiceAdded)
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
func (it *CmaccountServiceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceAdded represents a ServiceAdded event raised by the Cmaccount contract.
type CmaccountServiceAdded struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceAdded is a free log retrieval operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterServiceAdded(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceAddedIterator{contract: _Cmaccount.contract, event: "ServiceAdded", logs: logs, sub: sub}, nil
}

// WatchServiceAdded is a free log subscription operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchServiceAdded(opts *bind.WatchOpts, sink chan<- *CmaccountServiceAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceAdded", log); err != nil {
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

// ParseServiceAdded is a log parse operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseServiceAdded(log types.Log) (*CmaccountServiceAdded, error) {
	event := new(CmaccountServiceAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceCapabilitiesUpdatedIterator is returned from FilterServiceCapabilitiesUpdated and is used to iterate over the raw logs and unpacked data for ServiceCapabilitiesUpdated events raised by the Cmaccount contract.
type CmaccountServiceCapabilitiesUpdatedIterator struct {
	Event *CmaccountServiceCapabilitiesUpdated // Event containing the contract specifics and raw log

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
func (it *CmaccountServiceCapabilitiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceCapabilitiesUpdated)
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
		it.Event = new(CmaccountServiceCapabilitiesUpdated)
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
func (it *CmaccountServiceCapabilitiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceCapabilitiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceCapabilitiesUpdated represents a ServiceCapabilitiesUpdated event raised by the Cmaccount contract.
type CmaccountServiceCapabilitiesUpdated struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilitiesUpdated is a free log retrieval operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterServiceCapabilitiesUpdated(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceCapabilitiesUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceCapabilitiesUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceCapabilitiesUpdatedIterator{contract: _Cmaccount.contract, event: "ServiceCapabilitiesUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilitiesUpdated is a free log subscription operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchServiceCapabilitiesUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountServiceCapabilitiesUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceCapabilitiesUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceCapabilitiesUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilitiesUpdated", log); err != nil {
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

// ParseServiceCapabilitiesUpdated is a log parse operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseServiceCapabilitiesUpdated(log types.Log) (*CmaccountServiceCapabilitiesUpdated, error) {
	event := new(CmaccountServiceCapabilitiesUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilitiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceCapabilityAddedIterator is returned from FilterServiceCapabilityAdded and is used to iterate over the raw logs and unpacked data for ServiceCapabilityAdded events raised by the Cmaccount contract.
type CmaccountServiceCapabilityAddedIterator struct {
	Event *CmaccountServiceCapabilityAdded // Event containing the contract specifics and raw log

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
func (it *CmaccountServiceCapabilityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceCapabilityAdded)
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
		it.Event = new(CmaccountServiceCapabilityAdded)
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
func (it *CmaccountServiceCapabilityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceCapabilityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceCapabilityAdded represents a ServiceCapabilityAdded event raised by the Cmaccount contract.
type CmaccountServiceCapabilityAdded struct {
	ServiceName common.Hash
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityAdded is a free log retrieval operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) FilterServiceCapabilityAdded(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceCapabilityAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceCapabilityAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceCapabilityAddedIterator{contract: _Cmaccount.contract, event: "ServiceCapabilityAdded", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityAdded is a free log subscription operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) WatchServiceCapabilityAdded(opts *bind.WatchOpts, sink chan<- *CmaccountServiceCapabilityAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceCapabilityAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceCapabilityAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityAdded", log); err != nil {
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

// ParseServiceCapabilityAdded is a log parse operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) ParseServiceCapabilityAdded(log types.Log) (*CmaccountServiceCapabilityAdded, error) {
	event := new(CmaccountServiceCapabilityAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceCapabilityRemovedIterator is returned from FilterServiceCapabilityRemoved and is used to iterate over the raw logs and unpacked data for ServiceCapabilityRemoved events raised by the Cmaccount contract.
type CmaccountServiceCapabilityRemovedIterator struct {
	Event *CmaccountServiceCapabilityRemoved // Event containing the contract specifics and raw log

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
func (it *CmaccountServiceCapabilityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceCapabilityRemoved)
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
		it.Event = new(CmaccountServiceCapabilityRemoved)
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
func (it *CmaccountServiceCapabilityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceCapabilityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceCapabilityRemoved represents a ServiceCapabilityRemoved event raised by the Cmaccount contract.
type CmaccountServiceCapabilityRemoved struct {
	ServiceName common.Hash
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityRemoved is a free log retrieval operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) FilterServiceCapabilityRemoved(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceCapabilityRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceCapabilityRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceCapabilityRemovedIterator{contract: _Cmaccount.contract, event: "ServiceCapabilityRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityRemoved is a free log subscription operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) WatchServiceCapabilityRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountServiceCapabilityRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceCapabilityRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceCapabilityRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityRemoved", log); err != nil {
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

// ParseServiceCapabilityRemoved is a log parse operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) ParseServiceCapabilityRemoved(log types.Log) (*CmaccountServiceCapabilityRemoved, error) {
	event := new(CmaccountServiceCapabilityRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceRemovedIterator is returned from FilterServiceRemoved and is used to iterate over the raw logs and unpacked data for ServiceRemoved events raised by the Cmaccount contract.
type CmaccountServiceRemovedIterator struct {
	Event *CmaccountServiceRemoved // Event containing the contract specifics and raw log

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
func (it *CmaccountServiceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceRemoved)
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
		it.Event = new(CmaccountServiceRemoved)
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
func (it *CmaccountServiceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceRemoved represents a ServiceRemoved event raised by the Cmaccount contract.
type CmaccountServiceRemoved struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceRemoved is a free log retrieval operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterServiceRemoved(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceRemovedIterator{contract: _Cmaccount.contract, event: "ServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceRemoved is a free log subscription operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchServiceRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountServiceRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
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

// ParseServiceRemoved is a log parse operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseServiceRemoved(log types.Log) (*CmaccountServiceRemoved, error) {
	event := new(CmaccountServiceRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceRestrictedRateUpdatedIterator is returned from FilterServiceRestrictedRateUpdated and is used to iterate over the raw logs and unpacked data for ServiceRestrictedRateUpdated events raised by the Cmaccount contract.
type CmaccountServiceRestrictedRateUpdatedIterator struct {
	Event *CmaccountServiceRestrictedRateUpdated // Event containing the contract specifics and raw log

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
func (it *CmaccountServiceRestrictedRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceRestrictedRateUpdated)
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
		it.Event = new(CmaccountServiceRestrictedRateUpdated)
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
func (it *CmaccountServiceRestrictedRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceRestrictedRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceRestrictedRateUpdated represents a ServiceRestrictedRateUpdated event raised by the Cmaccount contract.
type CmaccountServiceRestrictedRateUpdated struct {
	ServiceName    common.Hash
	RestrictedRate bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterServiceRestrictedRateUpdated is a free log retrieval operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Cmaccount *CmaccountFilterer) FilterServiceRestrictedRateUpdated(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceRestrictedRateUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceRestrictedRateUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceRestrictedRateUpdatedIterator{contract: _Cmaccount.contract, event: "ServiceRestrictedRateUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceRestrictedRateUpdated is a free log subscription operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Cmaccount *CmaccountFilterer) WatchServiceRestrictedRateUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountServiceRestrictedRateUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceRestrictedRateUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceRestrictedRateUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceRestrictedRateUpdated", log); err != nil {
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

// ParseServiceRestrictedRateUpdated is a log parse operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Cmaccount *CmaccountFilterer) ParseServiceRestrictedRateUpdated(log types.Log) (*CmaccountServiceRestrictedRateUpdated, error) {
	event := new(CmaccountServiceRestrictedRateUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceRestrictedRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Cmaccount contract.
type CmaccountUpgradedIterator struct {
	Event *CmaccountUpgraded // Event containing the contract specifics and raw log

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
func (it *CmaccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountUpgraded)
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
		it.Event = new(CmaccountUpgraded)
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
func (it *CmaccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountUpgraded represents a Upgraded event raised by the Cmaccount contract.
type CmaccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccount *CmaccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CmaccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountUpgradedIterator{contract: _Cmaccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccount *CmaccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CmaccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountUpgraded)
				if err := _Cmaccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccount *CmaccountFilterer) ParseUpgraded(log types.Log) (*CmaccountUpgraded, error) {
	event := new(CmaccountUpgraded)
	if err := _Cmaccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountWantedServiceAddedIterator is returned from FilterWantedServiceAdded and is used to iterate over the raw logs and unpacked data for WantedServiceAdded events raised by the Cmaccount contract.
type CmaccountWantedServiceAddedIterator struct {
	Event *CmaccountWantedServiceAdded // Event containing the contract specifics and raw log

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
func (it *CmaccountWantedServiceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountWantedServiceAdded)
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
		it.Event = new(CmaccountWantedServiceAdded)
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
func (it *CmaccountWantedServiceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountWantedServiceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountWantedServiceAdded represents a WantedServiceAdded event raised by the Cmaccount contract.
type CmaccountWantedServiceAdded struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceAdded is a free log retrieval operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterWantedServiceAdded(opts *bind.FilterOpts, serviceName []string) (*CmaccountWantedServiceAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "WantedServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountWantedServiceAddedIterator{contract: _Cmaccount.contract, event: "WantedServiceAdded", logs: logs, sub: sub}, nil
}

// WatchWantedServiceAdded is a free log subscription operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchWantedServiceAdded(opts *bind.WatchOpts, sink chan<- *CmaccountWantedServiceAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "WantedServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountWantedServiceAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceAdded", log); err != nil {
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

// ParseWantedServiceAdded is a log parse operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseWantedServiceAdded(log types.Log) (*CmaccountWantedServiceAdded, error) {
	event := new(CmaccountWantedServiceAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountWantedServiceRemovedIterator is returned from FilterWantedServiceRemoved and is used to iterate over the raw logs and unpacked data for WantedServiceRemoved events raised by the Cmaccount contract.
type CmaccountWantedServiceRemovedIterator struct {
	Event *CmaccountWantedServiceRemoved // Event containing the contract specifics and raw log

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
func (it *CmaccountWantedServiceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountWantedServiceRemoved)
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
		it.Event = new(CmaccountWantedServiceRemoved)
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
func (it *CmaccountWantedServiceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountWantedServiceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountWantedServiceRemoved represents a WantedServiceRemoved event raised by the Cmaccount contract.
type CmaccountWantedServiceRemoved struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceRemoved is a free log retrieval operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterWantedServiceRemoved(opts *bind.FilterOpts, serviceName []string) (*CmaccountWantedServiceRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "WantedServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountWantedServiceRemovedIterator{contract: _Cmaccount.contract, event: "WantedServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchWantedServiceRemoved is a free log subscription operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchWantedServiceRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountWantedServiceRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "WantedServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountWantedServiceRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceRemoved", log); err != nil {
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

// ParseWantedServiceRemoved is a log parse operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseWantedServiceRemoved(log types.Log) (*CmaccountWantedServiceRemoved, error) {
	event := new(CmaccountWantedServiceRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Cmaccount contract.
type CmaccountWithdrawIterator struct {
	Event *CmaccountWithdraw // Event containing the contract specifics and raw log

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
func (it *CmaccountWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountWithdraw)
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
		it.Event = new(CmaccountWithdraw)
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
func (it *CmaccountWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountWithdraw represents a Withdraw event raised by the Cmaccount contract.
type CmaccountWithdraw struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Cmaccount *CmaccountFilterer) FilterWithdraw(opts *bind.FilterOpts, receiver []common.Address) (*CmaccountWithdrawIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountWithdrawIterator{contract: _Cmaccount.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Cmaccount *CmaccountFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CmaccountWithdraw, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountWithdraw)
				if err := _Cmaccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Cmaccount *CmaccountFilterer) ParseWithdraw(log types.Log) (*CmaccountWithdraw, error) {
	event := new(CmaccountWithdraw)
	if err := _Cmaccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

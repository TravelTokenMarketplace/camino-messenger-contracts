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
	Fee            *big.Int
	RestrictedRate bool
	Capabilities   []string
}

// CmaccountMetaData contains all meta data concerning the Cmaccount contract.
var CmaccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"latestImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountImplementationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountNoUpgradeNeeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"use\",\"type\":\"uint8\"}],\"name\":\"InvalidPublicKeyUseType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefundLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrefundNotSpentYet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceededForPeriod\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportsOffChainPayment\",\"type\":\"bool\"}],\"name\":\"OffChainPaymentSupportUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceCapabilitiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ServiceFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"ServiceRestrictedRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOKING_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER_BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVICE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasMoney\",\"type\":\"uint256\"}],\"name\":\"addMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"addService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"addServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"addWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"buyBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasMoneyWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGasMoneyWithdrawalForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKeysAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pubKeyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getService\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service\",\"name\":\"service\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"isBotAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"mintBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offChainPaymentSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"removeMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"removePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"removeServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"removeSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"removeWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGasMoneyWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSupported\",\"type\":\"bool\"}],\"name\":\"setOffChainPaymentSupported\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"setServiceCapabilities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"setServiceRestrictedRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"reason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGasMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516149e062000104600039600081816127f30152818161281c0152612b3701526149e06000f3fe60806040526004361061033a5760003560e01c806301ffc9a71461034657806308564c191461037b578063136f50ca1461039d578063150b7a02146103bf57806318274da4146104035780631aca6376146104315780631c54f0f7146104535780631c5db99e14610473578063241bbbfc14610493578063248a9ca3146104a85780632a119380146104c85780632f2ff15d146104e8578063319d13f314610508578063337462741461052857806336568abe1461054a578063383aba871461056a57806339e4c7051461058c57806341bf7c69146105ac57806342072bbd146105cc578063432cf639146105e15780634f1ef286146106015780634f3f46391461061457806351889d6b1461063657806352d1902d146106565780635c9889941461066b5780635e07f8691461068b578063658db0af146106ab5780636d69fcaf146106ce5780636fc22cd1146106ee57806372afa3281461070e57806374aa20481461073057806374fe60e9146107505780637512e55b1461077057806376319190146107905780637eec56c7146107b0578063852b3ccb146107d3578063857cdbb8146107f557806385f438c1146108225780638c20f574146108445780638f69347d146108645780639010d07c1461088457806391d14854146108a45780639db5dbe4146108c4578063a217fddf146108e4578063a31aa039146108f9578063a3246ad314610919578063a7d022f814610946578063ad3cb1cc14610966578063b512463514610997578063b82923fb146109b7578063be667188146109cc578063c162d7da146109ec578063c6640e6814610a01578063ca15c87314610a21578063ccde65dc14610a41578063cd9ef91414610a61578063d09445c214610a81578063d3c7c2c714610aa3578063d547741f14610ab8578063da47d85614610ad8578063e0b78add14610b05578063e26a61bb14610b25578063e5a6725c14610b45578063e7bfce9a14610b65578063ea79d07a14610b85578063eb5ea27314610b9a578063ebc20d2014610bba578063ee3b641f14610bda578063f3fef3a314610bfa578063f51acaea14610c1a578063f72c0d8b14610c3a578063f7e45f0914610c5c578063f8c8765e14610c7c57600080fd5b3661034157005b600080fd5b34801561035257600080fd5b50610366610361366004613a73565b610c9c565b60405190151581526020015b60405180910390f35b34801561038757600080fd5b50610390610cc7565b6040516103729190613b47565b3480156103a957600080fd5b506103b2610d80565b6040516103729190613b5a565b3480156103cb57600080fd5b506103ea6103da366004613c76565b630a85bd0160e11b949350505050565b6040516001600160e01b03199091168152602001610372565b34801561040f57600080fd5b5061042361041e366004613ce1565b610da0565b604051908152602001610372565b34801561043d57600080fd5b5061045161044c366004613d15565b610dae565b005b34801561045f57600080fd5b5061045161046e366004613d56565b610e5b565b34801561047f57600080fd5b5061045161048e366004613e16565b610eff565b34801561049f57600080fd5b50610366610fb7565b3480156104b457600080fd5b506104236104c3366004613e4a565b610fcf565b3480156104d457600080fd5b506104516104e3366004613e7a565b610fef565b3480156104f457600080fd5b50610451610503366004613eb6565b611095565b34801561051457600080fd5b50610390610523366004613ce1565b6110b7565b34801561053457600080fd5b5061042360008051602061494b83398151915281565b34801561055657600080fd5b50610451610565366004613eb6565b6110c5565b34801561057657600080fd5b506104236000805160206148cb83398151915281565b34801561059857600080fd5b506104516105a7366004613e16565b6110f8565b3480156105b857600080fd5b506104516105c7366004613ee6565b6111ab565b3480156105d857600080fd5b506103b2611223565b3480156105ed57600080fd5b506104516105fc366004613f3a565b61123a565b61045161060f366004613fab565b6112ab565b34801561062057600080fd5b506106296112ca565b6040516103729190613ffa565b34801561064257600080fd5b5061045161065136600461400e565b6112e8565b34801561066257600080fd5b506104236113b9565b34801561067757600080fd5b50610451610686366004613e4a565b6113d6565b34801561069757600080fd5b506103906106a6366004613e4a565b61140b565b3480156106b757600080fd5b506106c0611513565b60405161037292919061403a565b3480156106da57600080fd5b506104516106e9366004614048565b611535565b3480156106fa57600080fd5b50610451610709366004613d56565b611556565b34801561071a57600080fd5b506104236000805160206148ab83398151915281565b34801561073c57600080fd5b5061045161074b366004614065565b611578565b34801561075c57600080fd5b5061045161076b366004613e7a565b611630565b34801561077c57600080fd5b5061045161078b3660046140ab565b61166a565b34801561079c57600080fd5b506104516107ab366004614048565b6116d9565b3480156107bc57600080fd5b506107c56116fa565b604051610372929190614182565b3480156107df57600080fd5b5061042360008051602061498b83398151915281565b34801561080157600080fd5b50610815610810366004614048565b611841565b60405161037291906141f4565b34801561082e57600080fd5b5061042360008051602061496b83398151915281565b34801561085057600080fd5b5061045161085f3660046140ab565b61192f565b34801561087057600080fd5b5061036661087f366004613e4a565b61199e565b34801561089057600080fd5b5061062961089f366004613d56565b6119cf565b3480156108b057600080fd5b506103666108bf366004613eb6565b6119fd565b3480156108d057600080fd5b506104516108df366004613d15565b611a33565b3480156108f057600080fd5b50610423600081565b34801561090557600080fd5b50610451610914366004614207565b611a86565b34801561092557600080fd5b50610939610934366004613e4a565b611aa7565b6040516103729190614222565b34801561095257600080fd5b50610451610961366004614263565b611ad4565b34801561097257600080fd5b50610815604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156109a357600080fd5b506103666109b2366004613ce1565b611b45565b3480156109c357600080fd5b50610451611b53565b3480156109d857600080fd5b506104516109e7366004613d56565b611bee565b3480156109f857600080fd5b50610629611c28565b348015610a0d57600080fd5b50610451610a1c366004614048565b611c43565b348015610a2d57600080fd5b50610423610a3c366004613e4a565b611cde565b348015610a4d57600080fd5b50610451610a5c366004613fab565b611d03565b348015610a6d57600080fd5b50610451610a7c3660046142b0565b611d25565b348015610a8d57600080fd5b5061042360008051602061492b83398151915281565b348015610aaf57600080fd5b50610939611ddd565b348015610ac457600080fd5b50610451610ad3366004613eb6565b611df7565b348015610ae457600080fd5b50610af8610af3366004613e4a565b611e13565b60405161037291906142e9565b348015610b1157600080fd5b50610366610b20366004614048565b611f3c565b348015610b3157600080fd5b50610451610b403660046142fc565b611f56565b348015610b5157600080fd5b50610451610b60366004613e4a565b611ff4565b348015610b7157600080fd5b50610451610b80366004614048565b612080565b348015610b9157600080fd5b506109396120a1565b348015610ba657600080fd5b50610423610bb5366004613e4a565b6120bb565b348015610bc657600080fd5b50610451610bd536600461438b565b6120e6565b348015610be657600080fd5b506106c0610bf5366004614048565b612153565b348015610c0657600080fd5b50610451610c1536600461400e565b61218e565b348015610c2657600080fd5b50610451610c35366004613ce1565b612234565b348015610c4657600080fd5b506104236000805160206148eb83398151915281565b348015610c6857600080fd5b50610451610c77366004614065565b612288565b348015610c8857600080fd5b50610451610c973660046143e4565b6122c2565b60006001600160e01b03198216635a05180f60e01b1480610cc15750610cc182612475565b92915050565b60606000610cd3610d80565b9050600081516001600160401b03811115610cf057610cf0613bb3565b604051908082528060200260200182016040528015610d2357816020015b6060815260200190600190039081610d0e5790505b50905060005b8251811015610d7957610d54838281518110610d4757610d47614440565b60200260200101516124aa565b828281518110610d6657610d66614440565b6020908102919091010152600101610d29565b5092915050565b60606000610d8c612526565b9050610d9a8160090161254a565b91505090565b6000610cc1610bb583612557565b60008051602061496b833981519152610dc6816125cd565b6001600160a01b038316610ded57604051633a954ecd60e21b815260040160405180910390fd5b604051632142170760e11b81523060048201526001600160a01b038481166024830152604482018490528516906342842e0e90606401600060405180830381600087803b158015610e3d57600080fd5b505af1158015610e51573d6000803e3d6000fd5b5050505050505050565b60008051602061498b833981519152610e73816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6321b87f3a610e956112ca565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018690526044810185905260640160006040518083038186803b158015610ee257600080fd5b505af4158015610ef6573d6000803e3d6000fd5b50505050505050565b60008051602061492b833981519152610f17816125cd565b60005b8251811015610fb2576000610f47848381518110610f3a57610f3a614440565b60200260200101516125d7565b9050610f528161260c565b838281518110610f6457610f64614440565b6020026020010151604051610f799190614456565b604051908190038120907f50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f890600090a250600101610f1a565b505050565b600080610fc2612526565b6003015460ff1692915050565b600080610fda61264a565b60009384526020525050604090206001015490565b60008051602061498b833981519152611007816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6307e473166110296112ca565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024810187905261ffff80871660448301528516606482015260840160006040518083038186803b15801561108157600080fd5b505af4158015610e51573d6000803e3d6000fd5b61109e82610fcf565b6110a7816125cd565b6110b1838361266e565b50505050565b6060610cc16106a683612557565b6001600160a01b03811633146110ee5760405163334bd91960e11b815260040160405180910390fd5b610fb282826126b0565b60008051602061492b833981519152611110816125cd565b60005b8251811015610fb257600061114084838151811061113357611133614440565b6020026020010151612557565b905061114b816126e9565b83828151811061115d5761115d614440565b60200260200101516040516111729190614456565b604051908190038120907f0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e90600090a250600101611113565b60008051602061492b8339815191526111c3816125cd565b6111d56111cf84612557565b83612727565b826040516111e39190614456565b604051908190038120838252907fdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41906020015b60405180910390a2505050565b6060600061122f612526565b9050610d9a8161254a565b60008051602061492b833981519152611252816125cd565b61126661125e866125d7565b858486612750565b846040516112749190614456565b604051908190038120907f763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae937527990600090a25050505050565b6112b36127e8565b6112bc82612878565b6112c682826129bb565b5050565b6000806112d5612a6f565b600101546001600160a01b031692915050565b60008051602061494b833981519152611300816125cd565b6001600160a01b03831661132757604051633a954ecd60e21b815260040160405180910390fd5b61133f6000805160206148ab8339815191528461266e565b5061135860008051602061498b8339815191528461266e565b506113716000805160206148cb8339815191528461266e565b506040516001600160a01b038416907fdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a87399490600090a2610fb26001600160a01b03841683612a93565b60006113c3612b2c565b5060008051602061490b83398151915290565b6113de612b75565b6000805160206148cb8339815191526113f6816125cd565b6113ff82612bab565b50611408612cd8565b50565b60606000611417612526565b90506114238382612ce9565b806002016000848152602001908152602001600020600201805480602002602001604051908101604052809291908181526020016000905b8282101561150757838290600052602060002001805461147a90614472565b80601f01602080910402602001604051908101604052809291908181526020018280546114a690614472565b80156114f35780601f106114c8576101008083540402835291602001916114f3565b820191906000526020600020905b8154815290600101906020018083116114d657829003601f168201915b50505050508152602001906001019061145b565b50505050915050919050565b6000806000611520612d13565b90508060020154816003015492509250509091565b60008051602061492b83398151915261154d816125cd565b6112c682612d37565b60008051602061494b83398151915261156e816125cd565b610fb28383612dad565b60008051602061498b833981519152611590816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63fd13a43e6115b26112ca565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018890526044810187905261ffff80871660648301528516608482015260a40160006040518083038186803b15801561161157600080fd5b505af4158015611625573d6000803e3d6000fd5b505050505050505050565b60008051602061498b833981519152611648816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63793dddac6110296112ca565b60008051602061492b833981519152611682816125cd565b61169461168e84612557565b83612e07565b826040516116a29190614456565b60405180910390207f498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf8360405161121691906141f4565b60008051602061492b8339815191526116f1816125cd565b6112c682612e4b565b6060806000611707611223565b9050600081516001600160401b0381111561172457611724613bb3565b60405190808252806020026020018201604052801561175757816020015b60608152602001906001900390816117425790505b509050600082516001600160401b0381111561177557611775613bb3565b6040519080825280602002602001820160405280156117ae57816020015b61179b613974565b8152602001906001900390816117935790505b50905060005b8351811015611836576117d2848281518110610d4757610d47614440565b8382815181106117e4576117e4614440565b602002602001018190525061181184828151811061180457611804614440565b6020026020010151611e13565b82828151811061182357611823614440565b60209081029190910101526001016117b4565b509094909350915050565b6060600061184d612526565b905061185c6006820184612ec1565b611884578260405163ba650b5f60e01b815260040161187b9190613ffa565b60405180910390fd5b6001600160a01b0383166000908152600882016020526040902080546118a990614472565b80601f01602080910402602001604051908101604052809291908181526020018280546118d590614472565b80156119225780601f106118f757610100808354040283529160200191611922565b820191906000526020600020905b81548152906001019060200180831161190557829003601f168201915b5050505050915050919050565b60008051602061492b833981519152611947816125cd565b61195961195384612557565b83612ed6565b826040516119679190614456565b60405180910390207fba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d0570232648360405161121691906141f4565b6000806119a9612526565b90506119b58382612ce9565b600092835260020160205250604090206001015460ff1690565b6000806119da61300c565b60008581526020829052604090209091506119f59084613030565b949350505050565b600080611a0861264a565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b60008051602061496b833981519152611a4b816125cd565b6001600160a01b038316611a7257604051633a954ecd60e21b815260040160405180910390fd5b6110b16001600160a01b038516848461303c565b60008051602061492b833981519152611a9e816125cd565b6112c682613094565b60606000611ab361300c565b6000848152602082905260409020909150611acd9061254a565b9392505050565b60008051602061492b833981519152611aec816125cd565b611afe611af884612557565b836130eb565b82604051611b0c9190614456565b6040519081900381208315158252907f23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab90602001611216565b6000610cc161087f83612557565b60008051602061492b833981519152611b6b816125cd565b6000611b756116fa565b50905060005b8151811015610fb257611ba1611b9c83838151811061113357611133614440565b613125565b818181518110611bb357611bb3614440565b6020026020010151604051611bc89190614456565b6040519081900381209060008051602061488b83398151915290600090a2600101611b7b565b60008051602061498b833981519152611c06816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63348e06dd610e956112ca565b600080611c33612a6f565b546001600160a01b031692915050565b60008051602061494b833981519152611c5b816125cd565b611c736000805160206148ab833981519152836126b0565b50611c8c60008051602061498b833981519152836126b0565b50611ca56000805160206148cb833981519152836126b0565b506040516001600160a01b038316907fd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc6291390600090a25050565b600080611ce961300c565b6000848152602082905260409020909150611acd9061318e565b60008051602061492b833981519152611d1b816125cd565b610fb28383613198565b611d2d612b75565b60008051602061498b833981519152611d45816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__637adf63b7611d676112ca565b6040516001600160e01b031960e084901b1681526001600160a01b0391821660048201526024810188905260448101879052908516606482015260840160006040518083038186803b158015611dbc57600080fd5b505af4158015611dd0573d6000803e3d6000fd5b5050505050610fb2612cd8565b60606000611de9612526565b9050610d9a6004820161254a565b611e0082610fcf565b611e09816125cd565b6110b183836126b0565b611e1b613974565b6000611e25612526565b9050611e318382612ce9565b6000838152600280830160209081526040808420815160608101835281548152600182015460ff161515818501529381018054835181860281018601855281815295969295938701949192909184015b82821015611f2d578382906000526020600020018054611ea090614472565b80601f0160208091040260200160405190810160405280929190818152602001828054611ecc90614472565b8015611f195780601f10611eee57610100808354040283529160200191611f19565b820191906000526020600020905b815481529060010190602001808311611efc57829003601f168201915b505050505081526020019060010190611e81565b50505091525090949350505050565b6000610cc16000805160206148ab833981519152836119fd565b60008051602061498b833981519152611f6e816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63e4c22569611f906112ca565b8a8a8a8a8a8a8a6040518963ffffffff1660e01b8152600401611fba9897969594939291906144ac565b60006040518083038186803b158015611fd257600080fd5b505af4158015611fe6573d6000803e3d6000fd5b505050505050505050505050565b60008051602061498b83398151915261200c816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63c7bffa9661202e6112ca565b846040518363ffffffff1660e01b815260040161204c92919061450a565b60006040518083038186803b15801561206457600080fd5b505af4158015612078573d6000803e3d6000fd5b505050505050565b60008051602061492b833981519152612098816125cd565b6112c682613234565b606060006120ad612526565b9050610d9a8160060161254a565b6000806120c6612526565b90506120d28382612ce9565b600092835260020160205250604090205490565b60008051602061492b8339815191526120fe816125cd565b61211061210a84612557565b836132cd565b8260405161211e9190614456565b604051908190038120907fd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c53137190600090a2505050565b6000806000612160612d13565b6001600160a01b03909416600090815260208581526040808320546001909701909152902054939492505050565b612196612b75565b60008051602061496b8339815191526121ae816125cd565b6001600160a01b0383166121d557604051633a954ecd60e21b815260040160405180910390fd5b6121e86001600160a01b03841683612a93565b826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161222391815260200190565b60405180910390a2506112c6612cd8565b60008051602061492b83398151915261224c816125cd565b612258611b9c83612557565b816040516122669190614456565b6040519081900381209060008051602061488b83398151915290600090a25050565b60008051602061498b8339815191526122a0816125cd565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63b54e72d86115b26112ca565b60006122cc61330b565b805490915060ff600160401b82041615906001600160401b03166000811580156122f35750825b90506000826001600160401b0316600114801561230f5750303b155b90508115801561231d575080155b1561233b5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b0319166001178555831561236457845460ff60401b1916600160401b1785555b61236c61332f565b61237461332f565b61237c613337565b61238760008861266e565b506123a060008051602061492b8339815191528861266e565b506123b960008051602061494b8339815191528861266e565b506123d26000805160206148eb8339815191528761266e565b5060006123dd612a6f565b80546001600160a01b03808d166001600160a01b0319928316178355600183018054918d16919092161790559050678ac7230489e80000620151806124228282613347565b505050831561162557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050505050565b60006001600160e01b03198216637965db0b60e01b1480610cc157506301ffc9a760e01b6001600160e01b0319831614610cc1565b60606124b4611c28565b6001600160a01b031663306ade09836040518263ffffffff1660e01b81526004016124e191815260200190565b600060405180830381865afa1580156124fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610cc19190810190614523565b7ff2856e5e1b7689dcde1bb551fd115c3cad8d243ea609d47a46b4d22ee58d300090565b60606000611acd8361336a565b6000612561611c28565b6001600160a01b0316631ca0e943836040518263ffffffff1660e01b815260040161258c91906141f4565b602060405180830381865afa1580156125a9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc19190614590565b61140881336133c6565b60006125e1611c28565b6001600160a01b031663352af39a836040518263ffffffff1660e01b815260040161258c91906141f4565b6000612616612526565b9050600061262760098301846133f1565b905080610fb257604051631a1e056960e01b81526004810184905260240161187b565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b60008061267961300c565b9050600061268785856133fd565b905080156119f55760008581526020839052604090206126a7908561349e565b50949350505050565b6000806126bb61300c565b905060006126c985856134b3565b905080156119f55760008581526020839052604090206126a7908561352b565b60006126f3612526565b905060006127046009830184613540565b905080610fb257604051637eae59f160e11b81526004810184905260240161187b565b6000612731612526565b905061273d8382612ce9565b6000928352600201602052604090912055565b600061275a612526565b9050600061276882876133f1565b90508061278a576040516221e3bb60e31b81526004810187905260240161187b565b60408051606081018252868152841515602080830191825282840188815260008b81526002888101845295902084518155925160018401805460ff191691151591909117905551805193949293611625938501929190910190613997565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061285857507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661284c61354c565b6001600160a01b031614155b156128765760405163703e46dd60e11b815260040160405180910390fd5b565b6000805160206148eb833981519152612890816125cd565b600061289a611c28565b6001600160a01b0316639d825bc56040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128fb91906145a9565b9050600061290761354c565b9050836001600160a01b0316816001600160a01b03160361293f5780846040516382afabc160e01b815260040161187b9291906145c6565b816001600160a01b0316846001600160a01b03161461297557818460405163699a021f60e11b815260040161187b9291906145c6565b836001600160a01b0316816001600160a01b03167fa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db460405160405180910390a350505050565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612a15575060408051601f3d908101601f19168201909252612a1291810190614590565b60015b612a345781604051634c9c8ce360e01b815260040161187b9190613ffa565b60008051602061490b8339815191528114612a6557604051632a87526960e21b81526004810182905260240161187b565b610fb28383613568565b7f0c7b73796c7cc89b9f849b9056a93200eba741881e57a1b03b9bedb2c0e0710090565b80471015612ab857478160405163cf47918160e01b815260040161187b92919061403a565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114612b05576040519150601f19603f3d011682016040523d82523d6000602084013e612b0a565b606091505b5050905080610fb25760405163d6bda27560e01b815260040160405180910390fd5b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146128765760405163703e46dd60e11b815260040160405180910390fd5b6000612b7f6135be565b805490915060011901612ba557604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000612bb5612d13565b90508060020154821115612be457806002015482604051631728bc5b60e31b815260040161187b92919061403a565b6003810154336000908152602083905260409020544291612c04916145f6565b811115612c2c5733600090815260018301602090815260408083208390559084905290208190555b6002820154336000908152600184016020526040902054612c4e9085906145f6565b1115612c755781600201548360405163d54b188760e01b815260040161187b92919061403a565b33600090815260018301602052604081208054859290612c969084906145f6565b90915550612ca690503384612a93565b60405183815233907fb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c290602001611216565b6000612ce26135be565b6001905550565b612cf381836135e2565b6112c657604051631e96f6ed60e21b81526004810183905260240161187b565b7f99a652063088b6badaeb0c7f680676baf720654b4f86f50167944489af637d0090565b6000612d41612526565b90506000612d52600483018461349e565b905080612d745782604051632872fbf960e11b815260040161187b9190613ffa565b6040516001600160a01b038416907fa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f90600090a2505050565b6000612db7612d13565b60028101849055600381018390556040519091507f8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e90612dfa908590859061403a565b60405180910390a1505050565b6000612e11612526565b9050612e1d8382612ce9565b6000838152600280830160209081526040832090910180546001810182559083529120016110b18382614666565b6000612e55612526565b90506000612e66600483018461352b565b905080612e885782604051631532e67160e21b815260040161187b9190613ffa565b6040516001600160a01b038416907f85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af290600090a2505050565b6000611acd836001600160a01b0384166135ea565b6000612ee0612526565b9050612eec8382612ce9565b600083815260028083016020526040822001905b81548110156130055783604051602001612f1a9190614456565b60405160208183030381529060405280519060200120828281548110612f4257612f42614440565b90600052602060002001604051602001612f5c919061471f565b6040516020818303038152906040528051906020012003612ffd5781548290612f8790600190614795565b81548110612f9757612f97614440565b90600052602060002001828281548110612fb357612fb3614440565b906000526020600020019081612fc991906147a8565b5081805480612fda57612fda614874565b600190038181906000526020600020016000612ff691906139ed565b9055613005565b600101612f00565b5050505050565b7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200090565b6000611acd8383613602565b610fb283846001600160a01b031663a9059cbb858560405160240161306292919061450a565b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061362c565b600061309e612526565b60038101805460ff19168415159081179091556040519081529091507fe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e39060200160405180910390a15050565b60006130f5612526565b90506131018382612ce9565b60009283526002016020526040909120600101805460ff1916911515919091179055565b600061312f612526565b9050600061313d8284613540565b90508061316057604051631e96f6ed60e21b81526004810184905260240161187b565b600083815260028084016020526040822082815560018101805460ff19169055919061300590830182613a27565b6000610cc1825490565b60006131a2612526565b905060006131b3600683018561349e565b9050806131d55783604051631a6107e360e31b815260040161187b9190613ffa565b6001600160a01b038416600090815260088301602052604090206131f98482614666565b506040516001600160a01b038516907f928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe8290600090a250505050565b600061323e612526565b9050600061324f600683018461352b565b905080613271578260405163ba650b5f60e01b815260040161187b9190613ffa565b6001600160a01b03831660009081526008830160205260408120613294916139ed565b6040516001600160a01b038416907fc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf90600090a2505050565b60006132d7612526565b90506132e38382612ce9565b600083815260028083016020908152604090922084516110b193919092019190850190613997565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b612876613694565b61333f613694565b6128766136b9565b61334f613694565b6000613359612d13565b600281019390935550600390910155565b6060816000018054806020026020016040519081016040528092919081815260200182805480156133ba57602002820191906000526020600020905b8154815260200190600101908083116133a6575b50505050509050919050565b6133d082826119fd565b6112c657808260405163e2517d3f60e01b815260040161187b92919061450a565b6000611acd83836136c1565b60008061340861264a565b905061341484846119fd565b613494576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561344a3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610cc1565b6000915050610cc1565b6000611acd836001600160a01b0384166136c1565b6000806134be61264a565b90506134ca84846119fd565b15613494576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610cc1565b6000611acd836001600160a01b03841661370b565b6000611acd838361370b565b60008051602061490b833981519152546001600160a01b031690565b613571826137f4565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156135b657610fb28282613850565b6112c66138c6565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b6000611acd83835b60009081526001919091016020526040902054151590565b600082600001828154811061361957613619614440565b9060005260206000200154905092915050565b600080602060008451602086016000885af18061364f576040513d6000823e3d81fd5b50506000513d91508115613667578060011415613674565b6001600160a01b0384163b155b156110b15783604051635274afe760e01b815260040161187b9190613ffa565b61369c6138e5565b61287657604051631afcd79f60e31b815260040160405180910390fd5b612cd8613694565b60006136cd83836135ea565b61370357508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610cc1565b506000610cc1565b6000818152600183016020526040812054801561349457600061372f600183614795565b855490915060009061374390600190614795565b90508082146137a857600086600001828154811061376357613763614440565b906000526020600020015490508087600001848154811061378657613786614440565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806137b9576137b9614874565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610cc1565b806001600160a01b03163b6000036138215780604051634c9c8ce360e01b815260040161187b9190613ffa565b60008051602061490b83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161386d9190614456565b600060405180830381855af49150503d80600081146138a8576040519150601f19603f3d011682016040523d82523d6000602084013e6138ad565b606091505b50915091506138bd8583836138ff565b95945050505050565b34156128765760405163b398979f60e01b815260040160405180910390fd5b60006138ef61330b565b54600160401b900460ff16919050565b6060826139145761390f8261394b565b611acd565b815115801561392b57506001600160a01b0384163b155b15610d795783604051639996b31560e01b815260040161187b9190613ffa565b80511561395b5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b604051806060016040528060008152602001600015158152602001606081525090565b8280548282559060005260206000209081019282156139dd579160200282015b828111156139dd57825182906139cd9082614666565b50916020019190600101906139b7565b506139e9929150613a41565b5090565b5080546139f990614472565b6000825580601f10613a09575050565b601f0160209004906000526020600020908101906114089190613a5e565b508054600082559060005260206000209081019061140891905b808211156139e9576000613a5582826139ed565b50600101613a41565b5b808211156139e95760008155600101613a5f565b600060208284031215613a8557600080fd5b81356001600160e01b031981168114611acd57600080fd5b60005b83811015613ab8578181015183820152602001613aa0565b50506000910152565b60008151808452613ad9816020860160208601613a9d565b601f01601f19169290920160200192915050565b60008282518085526020808601955060208260051b8401016020860160005b84811015613b3a57601f19868403018952613b28838351613ac1565b98840198925090830190600101613b0c565b5090979650505050505050565b602081526000611acd6020830184613aed565b6020808252825182820181905260009190848201906040850190845b81811015613b9257835183529284019291840191600101613b76565b50909695505050505050565b6001600160a01b038116811461140857600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715613bf157613bf1613bb3565b604052919050565b60006001600160401b03821115613c1257613c12613bb3565b50601f01601f191660200190565b600082601f830112613c3157600080fd5b8135613c44613c3f82613bf9565b613bc9565b818152846020838601011115613c5957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215613c8c57600080fd5b8435613c9781613b9e565b93506020850135613ca781613b9e565b92506040850135915060608501356001600160401b03811115613cc957600080fd5b613cd587828801613c20565b91505092959194509250565b600060208284031215613cf357600080fd5b81356001600160401b03811115613d0957600080fd5b6119f584828501613c20565b600080600060608486031215613d2a57600080fd5b8335613d3581613b9e565b92506020840135613d4581613b9e565b929592945050506040919091013590565b60008060408385031215613d6957600080fd5b50508035926020909101359150565b600082601f830112613d8957600080fd5b813560206001600160401b0380831115613da557613da5613bb3565b8260051b613db4838201613bc9565b9384528581018301938381019088861115613dce57600080fd5b84880192505b85831015613e0a57823584811115613dec5760008081fd5b613dfa8a87838c0101613c20565b8352509184019190840190613dd4565b98975050505050505050565b600060208284031215613e2857600080fd5b81356001600160401b03811115613e3e57600080fd5b6119f584828501613d78565b600060208284031215613e5c57600080fd5b5035919050565b803561ffff81168114613e7557600080fd5b919050565b600080600060608486031215613e8f57600080fd5b83359250613e9f60208501613e63565b9150613ead60408501613e63565b90509250925092565b60008060408385031215613ec957600080fd5b823591506020830135613edb81613b9e565b809150509250929050565b60008060408385031215613ef957600080fd5b82356001600160401b03811115613f0f57600080fd5b613f1b85828601613c20565b95602094909401359450505050565b80358015158114613e7557600080fd5b60008060008060808587031215613f5057600080fd5b84356001600160401b0380821115613f6757600080fd5b613f7388838901613c20565b955060208701359450613f8860408801613f2a565b93506060870135915080821115613f9e57600080fd5b50613cd587828801613d78565b60008060408385031215613fbe57600080fd5b8235613fc981613b9e565b915060208301356001600160401b03811115613fe457600080fd5b613ff085828601613c20565b9150509250929050565b6001600160a01b0391909116815260200190565b6000806040838503121561402157600080fd5b823561402c81613b9e565b946020939093013593505050565b918252602082015260400190565b60006020828403121561405a57600080fd5b8135611acd81613b9e565b6000806000806080858703121561407b57600080fd5b843593506020850135925061409260408601613e63565b91506140a060608601613e63565b905092959194509250565b600080604083850312156140be57600080fd5b82356001600160401b03808211156140d557600080fd5b6140e186838701613c20565b935060208501359150808211156140f757600080fd5b50613ff085828601613c20565b600060608301825184526020808401511515602086015260408401516060604087015282815180855260808801915060808160051b890101945060208301925060005b8181101561417557607f19898703018352614163868551613ac1565b95509284019291840191600101614147565b5093979650505050505050565b6040815260006141956040830185613aed565b6020838203818501528185518084528284019150828160051b85010183880160005b838110156141e557601f198784030185526141d3838351614104565b948601949250908501906001016141b7565b50909998505050505050505050565b602081526000611acd6020830184613ac1565b60006020828403121561421957600080fd5b611acd82613f2a565b6020808252825182820181905260009190848201906040850190845b81811015613b925783516001600160a01b03168352928401929184019160010161423e565b6000806040838503121561427657600080fd5b82356001600160401b0381111561428c57600080fd5b61429885828601613c20565b9250506142a760208401613f2a565b90509250929050565b6000806000606084860312156142c557600080fd5b833592506020840135915060408401356142de81613b9e565b809150509250925092565b602081526000611acd6020830184614104565b600080600080600080600060e0888a03121561431757600080fd5b873561432281613b9e565b965060208801356001600160401b0381111561433d57600080fd5b6143498a828b01613c20565b9650506040880135945060608801359350608088013561436881613b9e565b925060a0880135915061437d60c08901613f2a565b905092959891949750929550565b6000806040838503121561439e57600080fd5b82356001600160401b03808211156143b557600080fd5b6143c186838701613c20565b935060208501359150808211156143d757600080fd5b50613ff085828601613d78565b600080600080608085870312156143fa57600080fd5b843561440581613b9e565b9350602085013561441581613b9e565b9250604085013561442581613b9e565b9150606085013561443581613b9e565b939692955090935050565b634e487b7160e01b600052603260045260246000fd5b60008251614468818460208701613a9d565b9190910192915050565b600181811c9082168061448657607f821691505b6020821081036144a657634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b0389811682528881166020830152610100604083018190526000916144da8483018b613ac1565b6060850199909952608084019790975250509290931660a083015260c082015290151560e0909101529392505050565b6001600160a01b03929092168252602082015260400190565b60006020828403121561453557600080fd5b81516001600160401b0381111561454b57600080fd5b8201601f8101841361455c57600080fd5b805161456a613c3f82613bf9565b81815285602083850101111561457f57600080fd5b6138bd826020830160208601613a9d565b6000602082840312156145a257600080fd5b5051919050565b6000602082840312156145bb57600080fd5b8151611acd81613b9e565b6001600160a01b0392831681529116602082015260400190565b634e487b7160e01b600052601160045260246000fd5b80820180821115610cc157610cc16145e0565b601f821115610fb2576000816000526020600020601f850160051c810160208610156146325750805b601f850160051c820191505b818110156120785782815560010161463e565b600019600383901b1c191660019190911b1790565b81516001600160401b0381111561467f5761467f613bb3565b6146938161468d8454614472565b84614609565b602080601f8311600181146146c257600084156146b05750858301515b6146ba8582614651565b865550612078565b600085815260208120601f198616915b828110156146f1578886015182559484019460019091019084016146d2565b508582101561470f5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600080835461472d81614472565b60018281168015614745576001811461475a57614789565b60ff1984168752821515830287019450614789565b8760005260208060002060005b858110156147805781548a820152908401908201614767565b50505082870194505b50929695505050505050565b81810381811115610cc157610cc16145e0565b8181036147b3575050565b6147bd8254614472565b6001600160401b038111156147d4576147d4613bb3565b6147e28161468d8454614472565b6000601f82116001811461481057600083156147fe5750848201545b6148088482614651565b855550613005565b600085815260209020601f19841690600086815260209020845b8381101561484a578286015482556001958601959091019060200161482a565b508583101561470f5793015460001960f8600387901b161c19169092555050600190811b01905550565b634e487b7160e01b600052603160045260246000fdfe52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813e9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621ae562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c95189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9a95e87c5af084bf5db8491c3a6515da9dd6da39b24b0eb0af08d7b9cd808d91c6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e43acdf00ba9ef08b5f2c22768276611b9af078bf6c24fa36b34ec5e9f2eb061faa26469706673582212209221c4060705eee31750e5a79337ca81dd42a8753603aefd28f62fea0a84066064736f6c63430008180033",
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
// Solidity: function getService(bytes32 serviceHash) view returns((uint256,bool,string[]) service)
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
// Solidity: function getService(bytes32 serviceHash) view returns((uint256,bool,string[]) service)
func (_Cmaccount *CmaccountSession) GetService(serviceHash [32]byte) (PartnerConfigurationService, error) {
	return _Cmaccount.Contract.GetService(&_Cmaccount.CallOpts, serviceHash)
}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((uint256,bool,string[]) service)
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

// GetServiceFee is a free data retrieval call binding the contract method 0x18274da4.
//
// Solidity: function getServiceFee(string serviceName) view returns(uint256 fee)
func (_Cmaccount *CmaccountCaller) GetServiceFee(opts *bind.CallOpts, serviceName string) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceFee", serviceName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetServiceFee is a free data retrieval call binding the contract method 0x18274da4.
//
// Solidity: function getServiceFee(string serviceName) view returns(uint256 fee)
func (_Cmaccount *CmaccountSession) GetServiceFee(serviceName string) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceFee is a free data retrieval call binding the contract method 0x18274da4.
//
// Solidity: function getServiceFee(string serviceName) view returns(uint256 fee)
func (_Cmaccount *CmaccountCallerSession) GetServiceFee(serviceName string) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceFee0 is a free data retrieval call binding the contract method 0xeb5ea273.
//
// Solidity: function getServiceFee(bytes32 serviceHash) view returns(uint256 fee)
func (_Cmaccount *CmaccountCaller) GetServiceFee0(opts *bind.CallOpts, serviceHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceFee0", serviceHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetServiceFee0 is a free data retrieval call binding the contract method 0xeb5ea273.
//
// Solidity: function getServiceFee(bytes32 serviceHash) view returns(uint256 fee)
func (_Cmaccount *CmaccountSession) GetServiceFee0(serviceHash [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee0(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceFee0 is a free data retrieval call binding the contract method 0xeb5ea273.
//
// Solidity: function getServiceFee(bytes32 serviceHash) view returns(uint256 fee)
func (_Cmaccount *CmaccountCallerSession) GetServiceFee0(serviceHash [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee0(&_Cmaccount.CallOpts, serviceHash)
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
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (uint256,bool,string[])[] services)
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
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (uint256,bool,string[])[] services)
func (_Cmaccount *CmaccountSession) GetSupportedServices() (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	return _Cmaccount.Contract.GetSupportedServices(&_Cmaccount.CallOpts)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (uint256,bool,string[])[] services)
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

// AddService is a paid mutator transaction binding the contract method 0x432cf639.
//
// Solidity: function addService(string serviceName, uint256 fee, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactor) AddService(opts *bind.TransactOpts, serviceName string, fee *big.Int, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addService", serviceName, fee, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x432cf639.
//
// Solidity: function addService(string serviceName, uint256 fee, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountSession) AddService(serviceName string, fee *big.Int, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddService(&_Cmaccount.TransactOpts, serviceName, fee, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x432cf639.
//
// Solidity: function addService(string serviceName, uint256 fee, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactorSession) AddService(serviceName string, fee *big.Int, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddService(&_Cmaccount.TransactOpts, serviceName, fee, restrictedRate, capabilities)
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

// SetServiceFee is a paid mutator transaction binding the contract method 0x41bf7c69.
//
// Solidity: function setServiceFee(string serviceName, uint256 fee) returns()
func (_Cmaccount *CmaccountTransactor) SetServiceFee(opts *bind.TransactOpts, serviceName string, fee *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setServiceFee", serviceName, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x41bf7c69.
//
// Solidity: function setServiceFee(string serviceName, uint256 fee) returns()
func (_Cmaccount *CmaccountSession) SetServiceFee(serviceName string, fee *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceFee(&_Cmaccount.TransactOpts, serviceName, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x41bf7c69.
//
// Solidity: function setServiceFee(string serviceName, uint256 fee) returns()
func (_Cmaccount *CmaccountTransactorSession) SetServiceFee(serviceName string, fee *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceFee(&_Cmaccount.TransactOpts, serviceName, fee)
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

// CmaccountServiceFeeUpdatedIterator is returned from FilterServiceFeeUpdated and is used to iterate over the raw logs and unpacked data for ServiceFeeUpdated events raised by the Cmaccount contract.
type CmaccountServiceFeeUpdatedIterator struct {
	Event *CmaccountServiceFeeUpdated // Event containing the contract specifics and raw log

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
func (it *CmaccountServiceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceFeeUpdated)
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
		it.Event = new(CmaccountServiceFeeUpdated)
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
func (it *CmaccountServiceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceFeeUpdated represents a ServiceFeeUpdated event raised by the Cmaccount contract.
type CmaccountServiceFeeUpdated struct {
	ServiceName common.Hash
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeUpdated is a free log retrieval operation binding the contract event 0xdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41.
//
// Solidity: event ServiceFeeUpdated(string indexed serviceName, uint256 fee)
func (_Cmaccount *CmaccountFilterer) FilterServiceFeeUpdated(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceFeeUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceFeeUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceFeeUpdatedIterator{contract: _Cmaccount.contract, event: "ServiceFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceFeeUpdated is a free log subscription operation binding the contract event 0xdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41.
//
// Solidity: event ServiceFeeUpdated(string indexed serviceName, uint256 fee)
func (_Cmaccount *CmaccountFilterer) WatchServiceFeeUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountServiceFeeUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceFeeUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceFeeUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
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

// ParseServiceFeeUpdated is a log parse operation binding the contract event 0xdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41.
//
// Solidity: event ServiceFeeUpdated(string indexed serviceName, uint256 fee)
func (_Cmaccount *CmaccountFilterer) ParseServiceFeeUpdated(log types.Log) (*CmaccountServiceFeeUpdated, error) {
	event := new(CmaccountServiceFeeUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
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

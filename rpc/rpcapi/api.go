// Package rpcapi provides JSON RPC service.
package rpcapi

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/acentswap/aswap/v3/internal/swapapi"
	"github.com/acentswap/aswap/v3/params"
	"github.com/acentswap/aswap/v3/router"
	"github.com/acentswap/aswap/v3/tokens"
)

// RouterSwapAPI rpc api handler
type RouterSwapAPI struct{}

// RPCNullArgs null args
type RPCNullArgs struct{}

// RouterSwapKeyArgs args
type RouterSwapKeyArgs struct {
	ChainID  string `json:"chainid"`
	TxID     string `json:"txid"`
	LogIndex string `json:"logindex"`
}

// GetVersionInfo api
func (s *RouterSwapAPI) GetVersionInfo(r *http.Request, args *RPCNullArgs, result *string) error {
	version := params.VersionWithMeta
	*result = version
	return nil
}

// GetServerInfo api
func (s *RouterSwapAPI) GetServerInfo(r *http.Request, args *RPCNullArgs, result *swapapi.ServerInfo) error {
	serverInfo := swapapi.GetServerInfo()
	*result = *serverInfo
	return nil
}

// RegisterRouterSwap api
func (s *RouterSwapAPI) RegisterRouterSwap(r *http.Request, args *RouterSwapKeyArgs, result *swapapi.MapIntResult) error {
	res, err := swapapi.RegisterRouterSwap(args.ChainID, args.TxID, args.LogIndex)
	if err == nil && res != nil {
		*result = *res
	}
	return err
}

// GetRouterSwap api
func (s *RouterSwapAPI) GetRouterSwap(r *http.Request, args *RouterSwapKeyArgs, result *swapapi.SwapInfo) error {
	res, err := swapapi.GetRouterSwap(args.ChainID, args.TxID, args.LogIndex)
	if err == nil && res != nil {
		*result = *res
	}
	return err
}

// RouterGetSwapHistoryArgs args
type RouterGetSwapHistoryArgs struct {
	ChainID string `json:"chainid"`
	Address string `json:"address"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Status  string `json:"status"`
}

// GetRouterSwapHistory api
func (s *RouterSwapAPI) GetRouterSwapHistory(r *http.Request, args *RouterGetSwapHistoryArgs, result *[]*swapapi.SwapInfo) error {
	res, err := swapapi.GetRouterSwapHistory(args.ChainID, args.Address, args.Offset, args.Limit, args.Status)
	if err == nil && res != nil {
		*result = res
	}
	return err
}

// GetAllChainIDs api
func (s *RouterSwapAPI) GetAllChainIDs(r *http.Request, args *RPCNullArgs, result *[]*big.Int) error {
	*result = router.AllChainIDs
	return nil
}

// GetAllTokenIDs api
func (s *RouterSwapAPI) GetAllTokenIDs(r *http.Request, args *RPCNullArgs, result *[]string) error {
	*result = router.AllTokenIDs
	return nil
}

// GetAllMultichainTokens api
// nolint:gocritic // rpc need result of pointer type
func (s *RouterSwapAPI) GetAllMultichainTokens(r *http.Request, args *string, result *map[string]string) error {
	tokenID := *args
	*result = router.GetCachedMultichainTokens(tokenID)
	return nil
}

// GetChainConfig api
func (s *RouterSwapAPI) GetChainConfig(r *http.Request, args *string, result *swapapi.ChainConfig) error {
	chainID := *args
	bridge := router.GetBridgeByChainID(chainID)
	if bridge == nil {
		return fmt.Errorf("chainID %v not exist", chainID)
	}
	chainConfig := swapapi.ConvertChainConfig(bridge.GetChainConfig())
	if chainConfig != nil {
		*result = *chainConfig
		return nil
	}
	return fmt.Errorf("chain config not found")
}

// GetTokenConfigArgs args
type GetTokenConfigArgs struct {
	ChainID string `json:"chainid"`
	Address string `json:"address"`
}

// GetTokenConfig api
func (s *RouterSwapAPI) GetTokenConfig(r *http.Request, args *GetTokenConfigArgs, result *swapapi.TokenConfig) error {
	chainID := args.ChainID
	address := args.Address
	bridge := router.GetBridgeByChainID(chainID)
	if bridge == nil {
		return fmt.Errorf("chainID %v not exist", chainID)
	}
	tokenConfig := swapapi.ConvertTokenConfig(bridge.GetTokenConfig(address))
	if tokenConfig != nil {
		*result = *tokenConfig
		return nil
	}
	return fmt.Errorf("token config not found")
}

// GetSwapConfigArgs args
type GetSwapConfigArgs struct {
	TokenID string `json:"tokenid"`
	ChainID string `json:"chainid"`
}

// GetSwapConfig api
func (s *RouterSwapAPI) GetSwapConfig(r *http.Request, args *GetSwapConfigArgs, result *swapapi.SwapConfig) error {
	tokenID := args.TokenID
	chainID := args.ChainID
	swapConfig := swapapi.ConvertSwapConfig(tokens.GetSwapConfig(tokenID, chainID))
	if swapConfig != nil {
		*result = *swapConfig
		return nil
	}
	return fmt.Errorf("swap config not found")
}

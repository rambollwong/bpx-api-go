package bpx

import "github.com/syp25815/bpx-api-go/bpx/types"

func (c *Client) FuturesPositions(symbol string) (resp []*types.Positon) {
	params := map[string]any{}
	params["symbol"] = symbol
	url := API_BASE + "api/v1/position"
	c.wrapAgent(newAgent().
		Get(url), params).
		EndStruct(&resp)
	return
}

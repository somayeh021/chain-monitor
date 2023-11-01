package controller

import (
	"context"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"github.com/scroll-tech/chain-monitor/internal/config"
	"github.com/scroll-tech/chain-monitor/internal/logic/contracts"
	"github.com/scroll-tech/chain-monitor/internal/logic/events"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type ContractController struct {
	conf             config.Config
	eventGatherLogic *events.EventGather
	contractsLogic   *contracts.Contracts

	l1EventCategoryList []types.TxEventCategory
	l2EventCategoryList []types.TxEventCategory
}

func NewContractController(conf config.Config, client *ethclient.Client) *ContractController {
	c := &ContractController{
		eventGatherLogic: events.NewEventGather(),
		contractsLogic:   contracts.NewContracts(client),
	}

	if err := c.contractsLogic.Register(c.conf); err != nil {
		log.Crit("contract register failure", "error", err)
		return nil
	}

	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC20EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC721EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ERC1155EventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.ETHEventCategory)
	c.l1EventCategoryList = append(c.l1EventCategoryList, types.MessengerEventCategory)

	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC20EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC721EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ERC1155EventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.ETHEventCategory)
	c.l2EventCategoryList = append(c.l2EventCategoryList, types.MessengerEventCategory)

	return c
}

// Watch the l1/l2 events, contains gateways events, transfer events, messenger events
func (c *ContractController) Watch() error {

}

// @param Start of the queried range
// @param End of the range (nil = latest)
func (c *ContractController) l1Watch(ctx context.Context, start uint64, end *uint64) {
	for _, eventCategory := range c.l1EventCategoryList {
		opt := bind.FilterOpts{
			Start:   start,
			End:     end,
			Context: ctx,
		}
		wrapIterList, err := c.contractsLogic.Iterator(ctx, &opt, types.Layer1, eventCategory)
		if err != nil {
			log.Error("get contract iterator failed", "layer", types.Layer1, "eventCategory", eventCategory, "error", err)
			continue
		}

		// parse the event data
		eventDataList := c.eventGatherLogic.Dispatch(ctx, types.Layer1, eventCategory, wrapIterList)
		if eventDataList == nil {
			log.Error("event gather deal event return empty data")
			continue
		}

		// match transfer event
		// insert to db, wait the l1/l2 event match
	}
}

func (c *ContractController) l2Watch() {

}

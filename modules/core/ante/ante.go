package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/keeper"
)

type AnteDecorator struct {
	k *keeper.Keeper
}

func NewAnteDecorator(k *keeper.Keeper) AnteDecorator {
	return AnteDecorator{k: k}
}

func (ad AnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// do not run redundancy check on DeliverTx or simulate
	if (ctx.IsCheckTx() || ctx.IsReCheckTx()) && !simulate {
		// keep track of total packet messages and number of redundancies across `RecvPacket`, `AcknowledgePacket`, and `TimeoutPacket/OnClose`
		redundancies := 0
		packetMsgs := 0
		for _, m := range tx.GetMsgs() {
			switch msg := m.(type) {
			case *channeltypes.MsgRecvPacket:
				response, err := ad.k.RecvPacket(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgAcknowledgement:
				response, err := ad.k.Acknowledgement(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgTimeout:
				response, err := ad.k.Timeout(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgTimeoutOnClose:
				response, err := ad.k.TimeoutOnClose(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *clienttypes.MsgUpdateClient:
				_, err := ad.k.UpdateClient(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}

			default:
				// if the multiMsg tx has a msg that is not a packet msg or update msg, then we will not return error
				// regardless of if all packet messages are redundant. This ensures that non-packet messages get processed
				// even if they get batched with redundant packet messages.
				return next(ctx, tx, simulate)
			}
		}

		// only return error if all packet messages are redundant
		if redundancies == packetMsgs && packetMsgs > 0 {
			return ctx, channeltypes.ErrRedundantTx
		}
	}
	return next(ctx, tx, simulate)
}

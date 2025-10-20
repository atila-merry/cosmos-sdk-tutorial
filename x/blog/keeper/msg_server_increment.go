package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/atila-merry/cosmos-sdk-tutorial/x/blog/types"
)

func (k msgServer) Increment(ctx context.Context, msg *types.MsgIncrement) (*types.MsgIncrementResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgIncrementResponse{}, nil
}

package blog

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/atila-merry/cosmos-sdk-tutorial/x/blog/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "Increment",
					Use:            "increment [count]",
					Short:          "Send a increment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "count"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

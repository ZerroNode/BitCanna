package bcna

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/BitCannaGlobal/bcna/api/bcna/bcna"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "BitcannaidAll",
					Use:       "list-bitcannaid",
					Short:     "List all bitcannaid",
				},
				{
					RpcMethod:      "Bitcannaid",
					Use:            "show-bitcannaid [id]",
					Short:          "Shows a bitcannaid by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "BitcannaidByBcnaid",
					Use:            "search-bitcannaid [bcnaid]",
					Short:          "Search a bitcannaid by bcnaid",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "bcnaid"}},
				},
				{
					RpcMethod: "SupplychainAll",
					Use:       "list-supplychain",
					Short:     "List all supplychain",
				},
				{
					RpcMethod:      "Supplychain",
					Use:            "show-supplychain [id]",
					Short:          "Shows a supplychain by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateBitcannaid",
					Use:            "create-bitcannaid [bcnaid] [address]",
					Short:          "Create bitcannaid",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "bcnaid"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "UpdateBitcannaid",
					Use:            "update-bitcannaid [id] [bcnaid] [address]",
					Short:          "Update bitcannaid",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "bcnaid"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "DeleteBitcannaid",
					Use:            "delete-bitcannaid [id]",
					Short:          "Delete bitcannaid",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateSupplychain",
					Use:            "create-supplychain [product] [info] [supplyinfo] [supplyextra]",
					Short:          "Create supplychain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "product"}, {ProtoField: "info"}, {ProtoField: "supplyinfo"}, {ProtoField: "supplyextra"}},
				},
				{
					RpcMethod:      "UpdateSupplychain",
					Use:            "update-supplychain [id] [product] [info] [supplyinfo] [supplyextra]",
					Short:          "Update supplychain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "product"}, {ProtoField: "info"}, {ProtoField: "supplyinfo"}, {ProtoField: "supplyextra"}},
				},
				{
					RpcMethod:      "DeleteSupplychain",
					Use:            "delete-supplychain [id]",
					Short:          "Delete supplychain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

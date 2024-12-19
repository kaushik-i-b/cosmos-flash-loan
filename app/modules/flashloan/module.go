package flashloan

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return "flashloan"
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	// Register the types for the flashloan module
}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return json.RawMessage{}
}

func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
}

func (AppModuleBasic) RegisterRESTRoutes(ctx client.Context, rtr *mux.Router) {
	// Register REST routes for the flashloan module
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

func NewAppModule(k Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string {
	return "flashloan"
}

func (am AppModule) NewHandler() sdk.Handler {
	return nil
}

func (am AppModule) QuerierRoute() string {
	return "flashloan"
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return nil
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return json.RawMessage{}
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
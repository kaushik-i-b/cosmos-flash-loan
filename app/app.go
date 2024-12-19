package app

import (
    "github.com/cosmos/cosmos-sdk/baseapp"
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/version"
    "github.com/cosmos/cosmos-sdk/x/auth"
    "github.com/cosmos/cosmos-sdk/x/bank"
    "github.com/cosmos/cosmos-sdk/x/staking"
    "github.com/yourusername/cosmos-flash-loan/app/modules/flashloan"
)

type App struct {
    *baseapp.BaseApp
    cdc      *codec.Codec
    keyMain  *types.KVStoreKey
    keyAuth  *types.KVStoreKey
    keyBank  *types.KVStoreKey
    keyStaking *types.KVStoreKey
    mm       *module.Manager
}

func NewApp() *App {
    cdc := codec.New()

    keyMain := types.NewKVStoreKey("main")
    keyAuth := types.NewKVStoreKey(auth.StoreKey)
    keyBank := types.NewKVStoreKey(bank.StoreKey)
    keyStaking := types.NewKVStoreKey(staking.StoreKey)

    app := &App{
        BaseApp: baseapp.NewBaseApp("cosmos-flash-loan", nil),
        cdc:     cdc,
        keyMain: keyMain,
        keyAuth: keyAuth,
        keyBank: keyBank,
        keyStaking: keyStaking,
        mm: module.NewManager(
            flashloan.NewModule(),
        ),
    }

    app.mm.SetOrderBeginBlockers(staking.ModuleName)
    app.mm.SetOrderEndBlockers(auth.ModuleName)

    return app
}

func (app *App) InitChainer(ctx types.Context, req types.RequestInitChain) types.ResponseInitChain {
    return app.mm.InitGenesis(ctx, req)
}

func (app *App) BeginBlocker(ctx types.Context, req types.RequestBeginBlock) {
    app.mm.BeginBlock(ctx, req)
}

func (app *App) EndBlocker(ctx types.Context, req types.RequestEndBlock) types.ResponseEndBlock {
    return app.mm.EndBlock(ctx, req)
}

func (app *App) LoadHeight(height int64) error {
    return app.LoadVersion(height)
}

func (app *App) Codec() *codec.Codec {
    return app.cdc
}

func (app *App) Name() string {
    return version.Version
}
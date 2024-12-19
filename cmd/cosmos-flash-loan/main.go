package main

import (
    "github.com/cosmos/cosmos-sdk/server"
    "github.com/cosmos/cosmos-sdk/server/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/x/auth"
    "github.com/cosmos/cosmos-sdk/x/bank"
    "github.com/cosmos/cosmos-sdk/x/flashloan"
    "os"
)

func main() {
    // Initialize the application
    app := NewApp()

    // Set up the server
    serverCtx := types.NewContext(app)

    // Start the application
    if err := server.Start(serverCtx); err != nil {
        os.Exit(1)
    }
}

func NewApp() *module.App {
    // Create and configure the application
    app := module.NewApp()

    // Register necessary modules
    app.RegisterModule(auth.NewAppModule())
    app.RegisterModule(bank.NewAppModule())
    app.RegisterModule(flashloan.NewAppModule())

    return app
}
package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"cosmos-flash-loan/app/modules/flashloan/types"
)

func NewTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flashloan",
		Short: "Flash loan transactions",
	}

	cmd.AddCommand(
		NewCreateFlashLoanCmd(cdc),
	)

	return cmd
}

func NewCreateFlashLoanCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create [amount] [recipient]",
		Short: "Create a new flash loan",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			amount := args[0]
			recipient := args[1]

			msg := types.NewMsgCreateFlashLoan(amount, recipient)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
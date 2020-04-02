package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/videocoin/go-bridge/cmd/nativetotoken"
	"github.com/videocoin/go-bridge/cmd/tokentonative"
)

func main() {
	root := &cobra.Command{
		Use:   "bridge [sub]",
		Short: "Bridge from/to ERC20 token to native blockchain coin.",
	}
	root.AddCommand(tokentonative.Command())
	root.AddCommand(nativetotoken.Command())
	if err := root.Execute(); err != nil {
		fmt.Printf("bridge execution failed with %v", err)
		os.Exit(1)
	}
}

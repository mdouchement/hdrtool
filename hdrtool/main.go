package main

import (
	"fmt"

	"github.com/mdouchement/hdrtool/hdrtool/cmd"
	"github.com/spf13/cobra"
)

func main() {
	c := &cobra.Command{
		Use: "hdrtools",
	}
	c.AddCommand(cmd.QualityCommand)
	c.AddCommand(cmd.ConvertCommand)
	c.AddCommand(cmd.HistogramCommand)

	if err := c.Execute(); err != nil {
		fmt.Println(err)
	}
}

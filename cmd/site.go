package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/guoweikuang/v2exgo/http"
)

var siteCmd = &cobra.Command{
	Use: "site",
	Short: "v2ex site",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
   Run: func(cmd *cobra.Command, args []string) {
	   fmt.Println(args)
	   subCmd := "stat"
	   if len(args) > 0 {
	   		if args[0] != "stat" && args[0] != "info" {
				fmt.Println("Invalid params")
				return
			} else {
				subCmd = args[0]
			}
	   }

	   siteService := http.NewRawV2exService("https://www.v2ex.com/api")

	   if subCmd == "stat" {
	   		stats, err := siteService.GetStats()
	   		if err != nil {
	   			panic(err)
			}

	   		fmt.Printf("Max Topic is %d\n", *stats.TopicMax)
	   		fmt.Printf("Max User is %d\n", *stats.MemberMax)
	   } else {
			info, err := siteService.GetInfo()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Domain: %s\n", info.Domain)
			fmt.Printf("Title: %s\n", info.Title)
			fmt.Printf("Description: %s\n", info.Description)
			fmt.Printf("Slogan: %s\n", info.Slogan)
	   }
   },
}

func init()  {
	rootCmd.AddCommand(siteCmd)
}
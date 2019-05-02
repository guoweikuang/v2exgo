package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/guoweikuang/v2exgo"
	"github.com/guoweikuang/v2exgo/http"
	"strconv"
)

var topicCmd = &cobra.Command{
	Use: "topic",
	Short: "v2ex topic",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		subCmd := "hot"

		if len(args) > 0 {
			subCmd = args[0]
		}

		topicService := http.NewTopicService("https://www.v2ex.com/api")
		switch subCmd {
		case "hot":
			showHotTopic(topicService)
		case "last":
			showLatestTopic(topicService)
		case "reply":
			showReplies(topicService, args[1])
		default:
			showTopic(topicService, subCmd)
		}

		//if subCmd == "hot" {
		//	showHotTopic(topicService)
		//}

	},
}

func init()  {
	rootCmd.AddCommand(topicCmd)
}

func showHotTopic(s v2ex.TopicService) {
	topics, err := s.GetHotTopic()
	if err != nil {
		panic(err)
	}

	for idx, topic := range topics {
		fmt.Printf("%d: %s\n", idx, s.TopicSummary(&topic))
	}
}

func showLatestTopic(s v2ex.TopicService) {
	topics, err := s.GetLatestTopic()
	if err != nil {
		panic(err)
	}

	for idx, topic := range topics {
		fmt.Printf("%2d: %s", idx, s.TopicSummary(&topic))
	}
}

func showTopic(s v2ex.TopicService, id string) {
	topicId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	topic, err := s.GetTopicByID(topicId)
	fmt.Printf("%s\n", s.TopicDetail(topic))
}

func showReplies(s v2ex.TopicService, id string) {
	topicId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	replies, err := s.GetRepliesByTopicID(topicId)

	for idx, reply := range replies {
		fmt.Printf("Id: %d", idx)
		fmt.Printf("%s", s.ReplyDetail(&reply))
		fmt.Println("------------------------\n")
	}
}
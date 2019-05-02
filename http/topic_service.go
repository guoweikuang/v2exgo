package http

import (
	"fmt"
	"github.com/guoweikuang/v2exgo"
	"strconv"
)


type TopicService struct {
	ApiUrl string
}

func NewTopicService(apiUrl string) *TopicService {
	return &TopicService{
		ApiUrl: apiUrl,
	}
}

func (t *TopicService) GetHotTopic() ([]v2ex.Topic, error) {
	url := t.ApiUrl + "/topics/hot.json"
	fmt.Println(url)
	var topics []v2ex.Topic
	err := v2ex.GetAPIData(url, &topics)
	return topics, err
}

func (t *TopicService) TopicSummary(v *v2ex.Topic) string {
	return fmt.Sprintf("%6d %3d  %s\n", v.ID, v.Replies, v.Title)
}

func (t *TopicService) GetLatestTopic() ([]v2ex.Topic, error) {
	url := t.ApiUrl + "/topics/latest.json"
	fmt.Println(url)
	var topics []v2ex.Topic
	err := v2ex.GetAPIData(url, &topics)
	return topics, err
}

func (t *TopicService) GetTopicByID(id int) (*v2ex.Topic, error) {
	var topics []*v2ex.Topic
	url := t.ApiUrl + "/topics/show.json?id=" + strconv.Itoa(int(id))
	err := v2ex.GetAPIData(url, &topics)
	return topics[0], err
}

func (t *TopicService) GetRepliesByTopicID(id int) ([]v2ex.Reply, error) {
	url := t.ApiUrl + "/replies/show.json?topic_id=" + strconv.Itoa(int(id))
	var relies []v2ex.Reply

	err := v2ex.GetAPIData(url, &relies)
	return relies, err
}

func (t *TopicService) ReplyDetail(r *v2ex.Reply) string {
	str := fmt.Sprintf("\tUser: %s ", r.Member.Username)
	str += fmt.Sprintf("\tTks: %d\n", r.Thanks)
	//str += fmt.Sprintln("\t---------------")
	return str
}

func (t *TopicService) TopicDetail(topic *v2ex.Topic) string {
	str := fmt.Sprintf("Title: %s\n", topic.Title)
	str += fmt.Sprintf("Url: %s\n", topic.URL)
	str += fmt.Sprintf("Author: %s\n", topic.Member.Username)
	str += fmt.Sprintf("Content: %s\n", topic.Content)
	str += fmt.Sprintf("Replicas: %d\n", topic.Replies)
	str += fmt.Sprintf("-------------\n")

	relies, err := t.GetRepliesByTopicID(topic.ID)
	if err != nil {
		return "error occur"
	}
	for idx, reply := range relies {
		str += fmt.Sprintf("\t ID: %d ", idx)
		str += fmt.Sprintf("%s\n", t.ReplyDetail(&reply))
		str += fmt.Sprintln("\t-----------------------")
	}

	return str
}
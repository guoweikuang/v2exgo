package v2ex

type Topic struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	Replies         int    `json:"replies"`
	Member          Member `json:"member"`
	Node            Node   `json:"node"`
	Created         int    `json:"created"`
	LastModified    int    `json:"last_modified"`
	LastTouched     int    `json:"last_touched"`
}


type Reply struct {
	ID              int    `json:"id"`
	Thanks          int    `json:"thanks"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	Member          Member `json:"member"`
	Created         int    `json:"created"`
	LastModified    int    `json:"last_modified"`
}


type TopicService interface {
	GetHotTopic() ([]Topic, error)
	TopicSummary(t *Topic) string
	GetLatestTopic() ([]Topic, error)
	GetTopicByID(id int) (*Topic, error)

	TopicDetail(t *Topic) string
}
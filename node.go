package v2ex

type Node struct {
	ID               int    `json:"id"`                //节点 ID
	Name             string `json:"name"`              //节点缩略名
	URL              string `json:"url"`               //节点地址
	Title            string `json:"title"`             //节点名称
	TitleAlternative string `json:"title_alternative"` //备选节点名称
	Topics           int    `json:"topics"`            //节点主题总数
	Header           string `json:"header"`            //节点头部信息
	Footer           string `json:"footer"`            //节点脚部信息
	Created          int    `json:"created"`           //节点创建时间
	Avatar
}

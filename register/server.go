package register

/*
  保存服务名称、服务部署的结点信息
*/
type Server struct {
	Name string `json:"name"`
	Node []*ServerNode `json:"node"`
}

/*
  保存结点id 端口号
*/
type ServerNode struct {
	Ip string `json:"ip"`
	Port string `json:"port"`
	Weight int `json:"weight"`
}

package types

type DomainState int

const (
	DomainState404              DomainState = iota // 0 接口不存在(极其严重) 15分钟 >= 1
	DomainStateContent                             // 1 内容异常(极其严重)   15分钟 >= 10
	DomainStateNetworkException                    // 2 网络错误(严重)       15分钟 >= 50
	DomainStateTimeout                             // 3 网络超时             15分钟 >= 200
	DomainStateNotAll                              // 4 全部不通(大概率是玩家网络问题)  15分钟 >= 500
	DomainStateOk                                  // 5 ok, 仅统计用
	DomainStateMax                                 // 6 only for count
)

func (v DomainState) String() string {

	switch v {
	case DomainState404:
		return "404"
	case DomainStateContent:
		return "ContentInvalid"
	case DomainStateNetworkException:
		return "NetworkException"
	case DomainStateTimeout:
		return "Timeout"
	case DomainStateNotAll:
		return "NotAll"
	case DomainStateOk:
		return "ok"
	}

	return "NoImpl"
}

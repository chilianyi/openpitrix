package constants

import "fmt"

const (
	en            = "en"
	zhCN          = "zh_cn"
	defaultLocale = zhCN
)

type NotifyMessage struct {
	en   string
	zhCN string
}

func (n *NotifyMessage) GetMessage(locale string, params ...interface{}) string {
	switch locale {
	case en:
		return fmt.Sprintf(n.en, params...)
	case zhCN:
		return fmt.Sprintf(n.zhCN, params)
	default:
		return fmt.Sprintf(n.zhCN, params)
	}
}

func (n *NotifyMessage) GetDefaultMessage(params ...interface{}) string {
	return n.GetMessage(defaultLocale, params...)
}

var (
	InviteIsvNotifyTitle = NotifyMessage{
		zhCN: "【OpenPitrix】邀请您成为平台服务商",
	}
	InviteIsvNotifyContent = NotifyMessage{
		zhCN: `
%s 您好：
	【OpenPitrix】邀请您入驻应用市场，成为优质服务商，为平台用户提供企业解决方案、产品和集成服务，共享快速收益。
`,
	}
)

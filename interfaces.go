package notify

const (
	RESULT_STATUS_ERROR   = 1
	RESULT_STATUS_SUCCESS = 0
)

type IClient interface {
	I() IClient
	Name() string
	Send(message IMessage) IResult
}

// IMessage 消息
type IMessage interface {
	// I 初始化并返回IMessage
	I() IMessage
	// AtMobiles 消息需要@人员手机号
	AtMobiles(mobiles []string)
	// AtUserIds 消息需要@人员userid
	AtUserIds(userIds []string)
	// AtAll @所有人
	AtAll()
}

// IResult 发送结果处理
type IResult interface {
	error
	// WithStatus 携带错误状态 RESULT_STATUS_ERROR 或 RESULT_STATUS_SUCCESS
	WithStatus(status int)
	// WithException  携带错误信息
	WithException(err error)
	// WithResult 携带请求接口返回值
	WithResult(response any)
	Status() int
	// Result 第三发返回数据
	Result() any
}

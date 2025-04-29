// fixclient/const.go
package fixclient

// 标准FIX4.4字段
type FixTag int

const (
	// 基础字段
	BeginString     FixTag = 8
	BodyLength      FixTag = 9
	MsgType         FixTag = 35
	MsgSeqNum       FixTag = 34
	SenderCompID    FixTag = 49
	TargetCompID    FixTag = 56
	SendingTime     FixTag = 52
	CheckSum        FixTag = 10
	PossDupFlag     FixTag = 43
	PossResend      FixTag = 97
	OrigSendingTime FixTag = 122

	// 登录相关
	EncryptMethod        FixTag = 98
	HeartBtInt           FixTag = 108
	RawDataLength        FixTag = 95
	RawData              FixTag = 96
	ResetSeqNumFlag      FixTag = 141
	Username             FixTag = 553
	Password             FixTag = 554
	TestMessageIndicator FixTag = 464
	Text                 FixTag = 58

	// 订单相关
	ClOrdID       FixTag = 11
	OrderID       FixTag = 37
	OrigClOrdID   FixTag = 41
	Symbol        FixTag = 55
	Side          FixTag = 54 // 1-Buy, 2-Sell
	OrdType       FixTag = 40 // 1-Market, 2-Limit
	Price         FixTag = 44
	StopPx        FixTag = 99
	OrderQty      FixTag = 38
	TimeInForce   FixTag = 59 // 1-GTC, 3-IOC, 4-FOK
	ExecInst      FixTag = 18 // 0-Add Liquidity, 1-ReduceOnly
	HandlInst     FixTag = 21
	MinQty        FixTag = 110
	MaxFloor      FixTag = 111
	ExpireDate    FixTag = 432
	CashOrderQty  FixTag = 152
	Account       FixTag = 1
	NoPartyIDs    FixTag = 453
	PartyID       FixTag = 448
	PartyIDSource FixTag = 447
	PartyRole     FixTag = 452

	// 执行报告
	ExecID       FixTag = 17
	ExecType     FixTag = 150 // 0-New, 1-Partial fill, 2-Fill, 4-Canceled, etc.
	OrdStatus    FixTag = 39  // 0-New, 1-Partially filled, 2-Filled, 4-Canceled, etc.
	OrdRejReason FixTag = 103 // 订单拒绝原因
	LeavesQty    FixTag = 151 // 剩余数量
	CumQty       FixTag = 14  // 累计成交数量
	AvgPx        FixTag = 6   // 平均成交价格
	LastPx       FixTag = 31  // 最新成交价格
	LastQty      FixTag = 32  // 最新成交数量
	TradeDate    FixTag = 75
	TransactTime FixTag = 60

	// Binance扩展字段
	RecvWindow              FixTag = 25000
	MessageHandling         FixTag = 25035 // 1-无序,2-顺序
	ResponseMode            FixTag = 25036 // 1-全部响应,2-仅ACK
	DropCopyFlag            FixTag = 9406  // Y/N
	OrderCategory           FixTag = 25010 // 1-NORMAL,2-MARGIN
	SelfTradePreventionMode FixTag = 25016 // 0-NONE, 1-EXPIRE_TAKER, 2-EXPIRE_MAKER, etc.
	TrailingDelta           FixTag = 25015
	WorkingFloor            FixTag = 25012
	PreventedMatchID        FixTag = 25020
	PreventedQuantity       FixTag = 25021
	PostOnly                FixTag = 25006 // Y-是, N-否
	ClientOrderMode         FixTag = 25003 // 1-标准, 2-SOR, 3-OCO
	AppliedWorkingFloor     FixTag = 25019 // 1-标准, 2-特殊

	ExpireTime FixTag = 126

	// 安全控制
	Signature  FixTag = 89
	SecurityID FixTag = 48

	// 撤单相关
	CxlRejResponseTo FixTag = 434 // 1=Order Cancel Request
	CxlRejReason     FixTag = 102 // 0=Too late to cancel

	// 标识和编号
	RefMsgType          FixTag = 372 // 引用消息类型
	RefSeqNum           FixTag = 45  // 引用序列号
	RefTagID            FixTag = 371 // 引用标签ID
	BusinessRejectRefID FixTag = 379 // 业务拒绝引用ID

	// 批量订单相关
	NoOrders FixTag = 73 // 订单数量
	ListID   FixTag = 66 // 列表ID

	// 限价订单本相关
	MDReqID     FixTag = 262 // 市场数据请求ID
	MDEntryType FixTag = 269 // 0=买入, 1=卖出
	NoMDEntries FixTag = 268 // 市场数据条目数量
	MDEntryPx   FixTag = 270 // 市场数据条目价格
	MDEntrySize FixTag = 271 // 市场数据条目数量

	// 查询相关
	MassStatusReqID   FixTag = 584 // 批量状态请求ID
	MassStatusReqType FixTag = 585 // 批量状态请求类型
)

// 消息类型常量
const (
	Heartbeat              = "0"
	TestRequest            = "1"
	ResendRequest          = "2"
	Reject                 = "3"
	SequenceReset          = "4"
	Logout                 = "5"
	ExecutionReport        = "8"
	OrderCancelReject      = "9"
	Logon                  = "A"
	NewOrderSingle         = "D"
	OrderCancelRequest     = "F"
	OrderStatusRequest     = "H"
	OrderMassStatusRequest = "AF"
	OrderMassCancelRequest = "q"
	BusinessMessageReject  = "j"
)

// 枚举值常量
const (
	// Side - 交易方向
	SIDE_BUY  = "1"
	SIDE_SELL = "2"

	// OrderType - 订单类型
	ORDER_TYPE_MARKET     = "1"
	ORDER_TYPE_LIMIT      = "2"
	ORDER_TYPE_STOP       = "3"
	ORDER_TYPE_STOP_LIMIT = "4"

	// TimeInForce - 有效期类型
	TIF_GTC = "1" // Good Till Cancel
	TIF_IOC = "3" // Immediate or Cancel
	TIF_FOK = "4" // Fill or Kill
	TIF_GTX = "5" // Good Till Crossing
	TIF_GTD = "6" // Good Till Date

	// ExecInst - 执行指令
	EXEC_INST_NORMAL      = "0" // 常规订单
	EXEC_INST_REDUCE_ONLY = "1" // 只减仓

	// OrdStatus - 订单状态
	ORD_STATUS_NEW              = "0" // 新订单
	ORD_STATUS_PARTIALLY_FILLED = "1" // 部分成交
	ORD_STATUS_FILLED           = "2" // 全部成交
	ORD_STATUS_CANCELED         = "4" // 已撤销
	ORD_STATUS_PENDING_CANCEL   = "6" // 撤销中
	ORD_STATUS_REJECTED         = "8" // 已拒绝
	ORD_STATUS_PENDING_NEW      = "A" // 待处理新订单
	ORD_STATUS_EXPIRED          = "C" // 已过期

	// ExecType - 执行类型
	EXEC_TYPE_NEW            = "0" // 新订单
	EXEC_TYPE_PARTIAL_FILL   = "1" // 部分成交
	EXEC_TYPE_FILL           = "2" // 全部成交
	EXEC_TYPE_CANCELED       = "4" // 已撤销
	EXEC_TYPE_REPLACED       = "5" // 已替换
	EXEC_TYPE_PENDING_CANCEL = "6" // 撤销中
	EXEC_TYPE_STOPPED        = "7" // 已停止
	EXEC_TYPE_REJECTED       = "8" // 已拒绝
	EXEC_TYPE_SUSPENDED      = "9" // 已挂起
	EXEC_TYPE_PENDING_NEW    = "A" // 待处理新订单
	EXEC_TYPE_EXPIRED        = "C" // 已过期
	EXEC_TYPE_RESTATED       = "D" // 重新声明
	EXEC_TYPE_TRADE          = "F" // 成交
	EXEC_TYPE_TRADE_CORRECT  = "G" // 成交更正
	EXEC_TYPE_TRADE_CANCEL   = "H" // 成交取消
	EXEC_TYPE_ORDER_STATUS   = "I" // 订单状态

	// EncryptMethod - 加密方法
	ENCRYPT_METHOD_NONE = "0" // 无加密

	// SelfTradePreventionMode - 自成交防护模式
	STP_MODE_NONE         = "0" // 不防护
	STP_MODE_EXPIRE_TAKER = "1" // 取消挂单
	STP_MODE_EXPIRE_MAKER = "2" // 取消maker订单
	STP_MODE_EXPIRE_BOTH  = "3" // 全部取消

	// MessageHandling - 消息处理模式
	MSG_HANDLING_UNORDERED = "1" // 无序处理
	MSG_HANDLING_ORDERED   = "2" // 顺序处理

	// ResponseMode - 响应模式
	RESPONSE_MODE_ALL = "1" // 全部响应
	RESPONSE_MODE_ACK = "2" // 仅ACK

	// OrderCategory - 订单类别
	ORDER_CATEGORY_NORMAL = "1" // 普通订单
	ORDER_CATEGORY_MARGIN = "2" // 杠杆订单

	// ClientOrderMode - 客户端订单模式
	CLIENT_ORDER_MODE_STANDARD = "1" // 标准模式
	CLIENT_ORDER_MODE_SOR      = "2" // SOR模式
	CLIENT_ORDER_MODE_OCO      = "3" // OCO模式
)

// 错误码常量
const (
	ERROR_UNKNOWN           = "-1000" // 未知错误
	ERROR_DISCONNECTED      = "-1001" // 断开连接
	ERROR_UNAUTHORIZED      = "-1002" // 未授权
	ERROR_TOO_MANY_REQUESTS = "-1003" // 请求过多
	ERROR_SERVER_BUSY       = "-1004" // 服务器忙
	ERROR_INVALID_FORMAT    = "-1100" // 无效格式
	ERROR_INVALID_TIMESTAMP = "-1021" // 无效时间戳
	ERROR_INVALID_SIGNATURE = "-1022" // 无效签名
	// ... 其他错误码
)

// 验证标签有效性的方法
func IsValidFixTag(t FixTag) bool {
	switch t {
	case BeginString, BodyLength, MsgType, MsgSeqNum, SenderCompID, TargetCompID,
		SendingTime, CheckSum, ClOrdID, OrderID, Symbol, Side, OrdType, Price,
		OrderQty, TimeInForce, ExecInst, ExecID, ExecType, OrdStatus, LeavesQty,
		CumQty, AvgPx, RecvWindow, MessageHandling, ResponseMode, DropCopyFlag,
		OrderCategory, Signature:
		return true
	default:
		return false
	}
}

package ctp

//RspInfoField 响应信息
type RspInfoField struct {
	///错误代码
	ErrorID  int
	///错误信息
	ErrorMsg string
}

//SettlementInfoField 投资者结算结果
type SettlementInfoField struct {
	///交易日
	TradingDay   string
	///结算编号
	SettlementID int
	///经纪公司代码
	BrokerID     string
	///投资者代码
	InvestorID   string
	///序号
	SequenceNo   int
	///消息正文
	Content      string
}

//SettlementInfoConfirmField 投资者结算结果确认信息
type SettlementInfoConfirmField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///确认日期
	ConfirmDate string
	///确认时间
	ConfirmTime string
}

//TradingAccountField 交易账号
type TradingAccountField struct {
	///经纪公司代码
	BrokerID string
	///投资者帐号
	AccountID string
	///上次质押金额
	PreMortgage float64
	///上次信用额度
	PreCredit float64
	///上次存款额
	PreDeposit float64
	///上次结算准备金
	PreBalance float64
	///上次占用的保证金
	PreMargin float64
	///利息基数
	InterestBase float64
	///利息收入
	Interest float64
	///入金金额
	Deposit float64
	///出金金额
	Withdraw float64
	///冻结的保证金
	FrozenMargin float64
	///冻结的资金
	FrozenCash float64
	///冻结的手续费
	FrozenCommission float64
	///当前保证金总额
	CurrMargin float64
	///资金差额
	CashIn float64
	///手续费
	Commission float64
	///平仓盈亏
	CloseProfit float64
	///持仓盈亏
	PositionProfit float64
	///期货结算准备金
	Balance float64
	///可用资金
	Available float64
	///可取资金
	WithdrawQuota float64
	///基本准备金
	Reserve float64
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///信用额度
	Credit float64
	///质押金额
	Mortgage float64
	///交易所保证金
	ExchangeMargin float64
	///投资者交割保证金
	DeliveryMargin float64
	///交易所交割保证金
	ExchangeDeliveryMargin float64
	///保底期货结算准备金
	ReserveBalance float64
	///币种代码
	CurrencyID string
	///上次货币质入金额
	PreFundMortgageIn float64
	///上次货币质出金额
	PreFundMortgageOut float64
	///货币质入金额
	FundMortgageIn float64
	///货币质出金额
	FundMortgageOut float64
	///货币质押余额
	FundMortgageAvailable float64
	///可质押货币金额
	MortgageableFund float64
	///特殊产品占用保证金
	SpecProductMargin float64
	///特殊产品冻结保证金
	SpecProductFrozenMargin float64
	///特殊产品手续费
	SpecProductCommission float64
	///特殊产品冻结手续费
	SpecProductFrozenCommission float64
	///特殊产品持仓盈亏
	SpecProductPositionProfit float64
	///特殊产品平仓盈亏
	SpecProductCloseProfit float64
	///根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg float64
	///特殊产品交易所保证金
	SpecProductExchangeMargin float64
}

//ReqUserLoginField 用户登录请求
type ReqUserLoginField struct {
	///交易日
	TradingDay string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///密码
	Password string
	///用户端产品信息
	UserProductInfo string
	///接口端产品信息
	InterfaceProductInfo string
	///协议信息
	ProtocolInfo string
	///Mac地址
	MacAddress string
	///动态密码
	OneTimePassword string
	///终端IP地址
	ClientIPAddress string
	///登录备注
	LoginRemark string
} 

//RspUserLoginInfoField 用户登录响应信息
type RspUserLoginInfoField struct {
	///交易日
	TradingDay string
	///登录成功时间
	LoginTime string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///交易系统名称
	SystemName string
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///最大报单引用
	MaxOrderRef string
	///上期所时间
	SHFETime string
	///大商所时间
	DCETime string
	///郑商所时间
	CZCETime string
	///中金所时间
	FFEXTime string
	///能源所时间
	INETime string
}

//ReqUserLogoutInfoField 用户登出请求
type ReqUserLogoutInfoField struct {
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
}

//SpecificInstrumentField 指定的合约
type SpecificInstrumentField struct {
	///合约代码
	InstrumentID string
}

//InstrumentField 合约
type InstrumentField struct {
	///合约代码
	InstrumentID string
	///交易所代码
	ExchangeID string
	///合约名称
	InstrumentName string
	///合约在交易所的代码
	ExchangeInstIDType string
	///产品代码
	ProductID string
	///产品类型
	ProductClass string
	///交割年份
	DeliveryYear int
	///交割月
	DeliveryMonth int
	///市价单最大下单量
	MaxMarketOrderVolume int
	///市价单最小下单量
	MinMarketOrderVolume int
	///限价单最大下单量
	MaxLimitOrderVolume int
	///限价单最小下单量
	MinLimitOrderVolume int
	///合约数量乘数
	VolumeMultiple int
	///最小变动价位
	PriceTick float64
	///创建日
	CreateDate string
	///上市日
	OpenDate string
	///到期日
	ExpireDate string
	///开始交割日
	StartDelivDate string
	///结束交割日
	EndDelivDate string
	///合约生命周期状态
	InstLifePhase string
	///当前是否交易
	IsTrading bool
	///持仓类型
	PositionType string
	///持仓日期类型
	PositionDateType string
	///多头保证金率
	LongMarginRatio float64
	///空头保证金率
	ShortMarginRatio float64
}

//InvestorPositionField 投资者持仓
type InvestorPositionField struct {
	///合约代码
	InstrumentID string
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///持仓多空方向
	PosiDirection string
	///投机套保标志
	HedgeFlag string
	///持仓日期类型
	PositionDate string
	///上日持仓
	YdPosition int
	///今日持仓
	Position int
	///多头冻结
	LongFrozen int
	///空头冻结
	ShortFrozen int
	///开仓冻结金额
	LongFrozenAmount float64
	///开仓冻结金额
	ShortFrozenAmount float64
	///开仓量
	OpenVolume int
	///平仓量
	CloseVolume int
	///开仓金额
	OpenAmount float64
	///平仓金额
	CloseAmount float64
	///持仓成本
	PositionCost float64
	///上次占用的保证金
	PreMargin float64
	///占用的保证金
	UseMargin float64
	///冻结的保证金
	FrozenMargin float64
	///冻结的资金
	FrozenCash float64
	///冻结的手续费
	FrozenCommission float64
	///资金差额
	CashIn float64
	///手续费
	Commission float64
	///平仓盈亏
	CloseProfit float64
	///持仓盈亏
	PositionProfit float64
	///上次结算价
	PreSettlementPrice float64
	///本次结算价
	SettlementPrice float64
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///开仓成本
	OpenCost float64
	///交易所保证金
	ExchangeMargin float64
	///组合成交形成的持仓
	CombPosition int
	///组合多头冻结
	CombLongFrozen int
	///组合空头冻结
	CombShortFrozen int
	///逐日盯市平仓盈亏
	CloseProfitByDate float64
	///逐笔对冲平仓盈亏
	CloseProfitByTrade float64
	///今日持仓
	TodayPosition int
	///保证金率
	MarginRateByMoney float64
	///保证金率(按手数)
	MarginRateByVolume float64
	///执行冻结
	StrikeFrozen int
	///执行冻结金额
	StrikeFrozenAmount float64
	///放弃执行冻结
	AbandonFrozen int
}

//DepthMarketDataField 深度行情
type DepthMarketDataField struct {
	///交易日
	TradingDay string
	///合约代码
	InstrumentID string
	///交易所代码
	ExchangeID string
	///合约在交易所的代码
	ExchangeInstID string
	///最新价
	LastPrice float64
	///上次结算价
	PreSettlementPrice float64
	///昨收盘
	PreClosePrice float64
	///昨持仓量
	PreOpenInterest float64
	///今开盘
	OpenPrice float64
	///最高价
	HighestPrice float64
	///最低价
	LowestPrice float64
	///数量
	Volume int
	///成交金额
	Turnover float64
	///持仓量
	OpenInterest float64
	///今收盘
	ClosePrice float64
	///本次结算价
	SettlementPrice float64
	///涨停板价
	UpperLimitPrice float64
	///跌停板价
	LowerLimitPrice float64
	///昨虚实度
	PreDelta float64
	///今虚实度
	CurrDelta float64
	///最后修改时间
	UpdateTime string
	///最后修改毫秒
	UpdateMillisec int
	///申买价一
	BidPrice1 float64
	///申买量一
	BidVolume1 float64
	///申卖价一
	AskPrice1 float64
	///申卖量一
	AskVolume1 float64
	///申买价二
	BidPrice2 float64
	///申买量二
	BidVolume2 float64
	///申卖价二
	AskPrice2 float64
	///申卖量二
	AskVolume2 float64
	///申买价三
	BidPrice3 float64
	///申买量三
	BidVolume3 float64
	///申卖价三
	AskPrice3 float64
	///申卖量三
	AskVolume3 float64
	///申买价四
	BidPrice4 float64
	///申买量四
	BidVolume4 float64
	///申卖价四
	AskPrice4 float64
	///申卖量四
	AskVolume4 float64
	///申买价五
	BidPrice5 float64
	///申买量五
	BidVolume5 float64
	///申卖价五
	AskPrice5 float64
	///申卖量五
	AskVolume5 float64
	///当日均价
	AveragePrice float64
}

/*//SessionInfo 会话信息
type SessionInfo struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	RequestID  int    ///请求编号
	FrontID    int    ///前置编号
	SessionID  int    ///会话编号
	ExchangeID string ///交易所代码
	UserID     string ///用户代码
}*/

//InputOrderField 报单
type InputOrderField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///合约代码
	InstrumentID string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///报单价格条件
	OrderPriceType byte
	///买卖方向
	Direction byte
	///组合开平标志
	CombOffsetFlag string
	///组合投机套保标志
	CombHedgeFlag string
	///价格
	LimitPrice float64
	///数量
	VolumeTotalOriginal int
	///有效期类型
	TimeCondition byte
	///GTD日期
	GTDDate string
	///成交量类型
	VolumeCondition byte
	///最小成交量
	MinVolume int
	///触发条件
	ContingentCondition byte
	///止损价
	StopPrice float64
	///强平原因
	ForceCloseReason byte
	///自动挂起标志
	IsAutoSuspend bool
	///业务单元
	BusinessUnit int
	///请求编号
	RequestID int
	///用户强评标志
	UserForceClose bool
	///互换单标志
	IsSwapOrder bool
}

//InputOrderActionField 报单引用
type InputOrderActionField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///报单操作引用
	OrderActionRef string
	///报单引用
	OrderRef string
	///请求编号
	RequestID int
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///交易所代码
	ExchangeID string
	///报单编号
	OrderSysID string
	///操作标志
	ActionFlag string
	///价格
	LimitPrice float64
	///数量变化
	VolumeChange int
	///用户代码
	UserID string
	///合约代码
	InstrumentID string
}

///OrderField报单
type OrderField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///合约代码
	InstrumentID string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///报单价格条件
	OrderPriceType byte
	///买卖方向
	Direction byte
	///组合开平标志
	CombOffsetFlag string
	///组合投机套保标志
	CombHedgeFlag string
	///价格
	LimitPrice float64
	///数量
	VolumeTotalOriginal int
	///有效期类型
	TimeCondition byte
	///GTD日期
	GTDDate string
	///成交量类型
	VolumeCondition byte
	///最小成交量
	MinVolume int
	///触发条件
	ContingentCondition byte
	///止损价
	StopPrice float64
	///强平原因
	ForceCloseReason byte
	///自动挂起标志
	IsAutoSuspend bool
	///业务单元
	BusinessUnit string
	///请求编号
	RequestID int
	///本地报单编号
	OrderLocalID string
	///交易所代码
	ExchangeID string
	///会员代码
	ParticipantID string
	///客户代码
	ClientID string
	///合约在交易所的代码
	ExchangeInstID string
	///交易所交易员代码
	TraderID string
	///安装编号
	InstallID int
	///报单提交状态
	OrderSubmitStatus byte
	///报单提示序号
	NotifySequence int
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///报单编号
	OrderSysID string
	///报单来源
	OrderSource byte
	///报单状态
	OrderStatus byte
	///报单类型
	OrderType byte
	///今成交数量
	VolumeTraded int
	///剩余数量
	VolumeTotal int
	///报单日期
	InsertDate string
	///委托时间
	InsertTime string
	///激活时间
	ActiveTime string
	///挂起时间
	SuspendTime string
	///最后修改时间
	UpdateTime string
	///撤销时间
	CancelTime string
	///最后修改交易所交易员代码
	ActiveTraderID string
	///结算会员编号
	ClearingPartID string
	///序号
	SequenceNo int
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///用户端产品信息
	UserProductInfo string
	///状态信息
	StatusMsg string
	///用户强评标志
	UserForceClose bool
	///操作用户代码
	ActiveUserID string
	///经纪公司报单编号
	BrokerOrderSeq int
	///相关报单
	RelativeOrderSysID string
	///郑商所成交数量
	ZCETotalTradedVolume int
	///互换单标志
	IsSwapOrder bool
}

///TradeField成交
type TradeField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///合约代码
	InstrumentID string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///交易所代码
	ExchangeID string
	///成交编号
	TradeID string
	///买卖方向
	Direction byte
	///报单编号
	OrderSysID string
	///会员代码
	ParticipantID string
	///客户代码
	ClientID string
	///交易角色
	TradingRole byte
	///合约在交易所的代码
	ExchangeInstID string
	///开平标志
	OffsetFlag byte
	///投机套保标志
	HedgeFlag byte
	///价格
	Price float64
	///数量
	Volume int
	///成交时期
	TradeDate string
	///成交时间
	TradeTime string
	///成交类型
	TradeType byte
	///成交价来源
	PriceSource byte
	///交易所交易员代码
	TraderID string
	///本地报单编号
	OrderLocalID string
	///结算会员编号
	ClearingPartID string
	///业务单元
	BusinessUnit string
	///序号
	SequenceNo int
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///经纪公司报单编号
	BrokerOrderSeq int
	///成交来源
	TradeSource byte
}
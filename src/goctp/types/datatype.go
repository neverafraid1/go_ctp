package ctp

const IntMax = 0xffff

/////////////////////////////////////////////////////////////////////////
///报单状态类型
/////////////////////////////////////////////////////////////////////////
///全部成交
const OrderStatus_AllTraded = '0'
///部分成交还在队列中
const OrderStatus_PartTradedQueueing = '1'
///部分成交不在队列中
const OrderStatus_PartTradedNotQueueing = '2'
///未成交还在队列中
const OrderStatus_NoTradeQueueing = '3'
///未成交不在队列中
const OrderStatus_NoTradeNotQueueing = '4'
///撤单
const OrderStatus_Canceled = '5'
///未知
const OrderStatus_Unknown = 'a'
///尚未触发
const OrderStatus_NotTouched = 'b'
///已触发
const OrderStatus_Touched = 'c'

/////////////////////////////////////////////////////////////////////////
///报单提交状态类型
/////////////////////////////////////////////////////////////////////////
///已经提交
const OrderSubmitStatus_InsertSubmitted = '0'
///撤单已经提交
const OrderSubmitStatus_CancelSubmitted = '1'
///修改已经提交
const OrderSubmitStatus_ModifySubmitted = '2'
///已经接受
const OrderSubmitStatus_Accepted = '3'
///报单已经被拒绝
const OrderSubmitStatus_InsertRejected = '4'
///撤单已经被拒绝
const OrderSubmitStatus_CancelRejected = '5'
///改单已经被拒绝
const OrderSubmitStatus_ModifyRejected = '6'

/////////////////////////////////////////////////////////////////////////
///持仓日期类型
/////////////////////////////////////////////////////////////////////////
///今日持仓
const PositionDate_Today = '1'
///历史持仓
const PositionDate_History = '2'

/////////////////////////////////////////////////////////////////////////
///产品类型类型
/////////////////////////////////////////////////////////////////////////
///期货
const ProductClass_Futures = '1'
///期货期权
const ProductClass_Options = '2'
///组合
const ProductClass_Combination = '3'
///即期
const ProductClass_Spot = '4'
///期转现
const ProductClass_EFP = '5'
///现货期权
const ProductClass_SpotOption = '6'


/////////////////////////////////////////////////////////////////////////
///合约生命周期状态类型
/////////////////////////////////////////////////////////////////////////
///未上市
const InstLifePhase_NotStart = '0'
///上市
const InstLifePhase_Started = '1'
///停牌
const InstLifePhase_Pause = '2'
///到期
const InstLifePhase_Expired = '3'

/////////////////////////////////////////////////////////////////////////
///买卖方向类型
/////////////////////////////////////////////////////////////////////////
///买
const Direction_Buy = '0'
///卖
const Direction_Sell = '1'

/////////////////////////////////////////////////////////////////////////
///持仓多空方向类型
/////////////////////////////////////////////////////////////////////////
///净
const PosiDirection_Net = '1'
///多头
const PosiDirection_Long = '2'
///空头
const PosiDirection_Short = '3'

/////////////////////////////////////////////////////////////////////////
///投机套保标志类型
/////////////////////////////////////////////////////////////////////////
///投机
const HedgeFlag_Speculation = '1'
///套利
const HedgeFlag_Arbitrage = '2'
///套保
const HedgeFlag_Hedge = '3'

/////////////////////////////////////////////////////////////////////////
///报单价格条件类型
/////////////////////////////////////////////////////////////////////////
///任意价
const OrderPriceType_AnyPrice = '1'
///限价
const OrderPriceType_LimitPrice = '2'
///最优价
const OrderPriceType_BestPrice = '3'
///最新价
const OrderPriceType_LastPrice = '4'
///最新价浮动上浮1个ticks
const OrderPriceType_LastPricePlusOneTicks = '5'
///最新价浮动上浮2个ticks
const OrderPriceType_LastPricePlusTwoTicks = '6'
///最新价浮动上浮3个ticks
const OrderPriceType_LastPricePlusThreeTicks = '7'
///卖一价
const OrderPriceType_AskPrice1 = '8'
///卖一价浮动上浮1个ticks
const OrderPriceType_AskPrice1PlusOneTicks = '9'
///卖一价浮动上浮2个ticks
const OrderPriceType_AskPrice1PlusTwoTicks = 'A'
///卖一价浮动上浮3个ticks
const OrderPriceType_AskPrice1PlusThreeTicks = 'B'
///买一价
const OrderPriceType_BidPrice1 = 'C'
///买一价浮动上浮1个ticks
const OrderPriceType_BidPrice1PlusOneTicks = 'D'
///买一价浮动上浮2个ticks
const OrderPriceType_BidPrice1PlusTwoTicks = 'E'
///买一价浮动上浮3个ticks
const OrderPriceType_BidPrice1PlusThreeTicks = 'F'
///五档价
const OrderPriceType_FiveLevelPrice = 'G'

/////////////////////////////////////////////////////////////////////////
///开平标志类型
/////////////////////////////////////////////////////////////////////////
///开仓
const OffsetFlag_Open = '0'
///平仓
const OffsetFlag_Close = '1'
///强平
const OffsetFlag_ForceClose = '2'
///平今
const OffsetFlag_CloseToday = '3'
///平昨
const OffsetFlag_CloseYesterday = '4'
///强减
const OffsetFlag_ForceOff = '5'
///本地强平
const OffsetFlag_LocalForceClose = '6'

/////////////////////////////////////////////////////////////////////////
///强平原因类型
/////////////////////////////////////////////////////////////////////////
///非强平
const ForceCloseReason_NotForceClose = '0'
///资金不足
const ForceCloseReason_LackDeposit = '1'
///客户超仓
const ForceCloseReason_ClientOverPositionLimit = '2'
///会员超仓
const ForceCloseReason_MemberOverPositionLimit = '3'
///持仓非整数倍
const ForceCloseReason_NotMultiple = '4'
///违规
const ForceCloseReason_Violation = '5'
///其它
const ForceCloseReason_Other = '6'
///自然人临近交割
const ForceCloseReason_PersonDeliv = '7'

/////////////////////////////////////////////////////////////////////////
///报单类型类型
/////////////////////////////////////////////////////////////////////////
///正常
const FtdcOrderType_Normal = '0'
///报价衍生
const FtdcOrderType_DeriveFromQuote = '1'
///组合衍生
const FtdcOrderType_DeriveFromCombination = '2'
///组合报单
const FtdcOrderType_Combination = '3'
///条件单
const FtdcOrderType_ConditionalOrder = '4'
///互换单
const FtdcOrderType_Swap = '5'

/////////////////////////////////////////////////////////////////////////
///有效期类型类型
/////////////////////////////////////////////////////////////////////////
///立即完成，否则撤销
const TimeCondition_IOC = '1'
///本节有效
const TimeCondition_GFS = '2'
///当日有效
const TimeCondition_GFD = '3'
///指定日期前有效
const TimeCondition_GTD = '4'
///撤销前有效
const TimeCondition_GTC = '5'
///集合竞价有效
const TimeCondition_GFA = '6'

/////////////////////////////////////////////////////////////////////////
///成交量类型类型
/////////////////////////////////////////////////////////////////////////
///任何数量
const VolumeCondition_AV = '1'
///最小数量
const VolumeCondition_MV = '2'
///全部数量
const VolumeCondition_CV = '3'

/////////////////////////////////////////////////////////////////////////
///触发条件类型
/////////////////////////////////////////////////////////////////////////
///立即
const ContingentCondition_Immediately = '1'
///止损
const ContingentCondition_Touch = '2'
///止赢
const ContingentCondition_TouchProfit = '3'
///预埋单
const ContingentCondition_ParkedOrder = '4'
///最新价大于条件价
const ContingentCondition_LastPriceGreaterThanStopPrice = '5'
///最新价大于等于条件价
const ContingentCondition_LastPriceGreaterEqualStopPrice = '6'
///最新价小于条件价
const ContingentCondition_LastPriceLesserThanStopPrice = '7'
///最新价小于等于条件价
const ContingentCondition_LastPriceLesserEqualStopPrice = '8'
///卖一价大于条件价
const ContingentCondition_AskPriceGreaterThanStopPrice = '9'
///卖一价大于等于条件价
const ContingentCondition_AskPriceGreaterEqualStopPrice = 'A'
///卖一价小于条件价
const ContingentCondition_AskPriceLesserThanStopPrice = 'B'
///卖一价小于等于条件价
const ContingentCondition_AskPriceLesserEqualStopPrice = 'C'
///买一价大于条件价
const ContingentCondition_BidPriceGreaterThanStopPrice = 'D'
///买一价大于等于条件价
const ContingentCondition_BidPriceGreaterEqualStopPrice = 'E'
///买一价小于条件价
const ContingentCondition_BidPriceLesserThanStopPrice = 'F'
///买一价小于等于条件价
const ContingentCondition_BidPriceLesserEqualStopPrice = 'H'

/////////////////////////////////////////////////////////////////////////
///操作标志类型
/////////////////////////////////////////////////////////////////////////
///删除
const ActionFlag_Delete = '0'
///修改
const ActionFlag_Modify = '3'

/////////////////////////////////////////////////////////////////////////
///交易角色类型
/////////////////////////////////////////////////////////////////////////
///代理
const TradingRole_Broker = '1'
///自营
const TradingRole_Host = '2'
///做市商
const TradingRole_Maker = '3'

/////////////////////////////////////////////////////////////////////////
///成交类型类型
/////////////////////////////////////////////////////////////////////////
///组合持仓拆分为单一持仓,初始化不应包含该类型的持仓
const TradeType_SplitCombination = '#'
///普通成交
const TradeType_Common = '0'
///期权执行
const TradeType_OptionsExecution = '1'
///OTC成交
const TradeType_OTC = '2'
///期转现衍生成交
const TradeType_EFPDerived = '3'
///组合衍生成交
const TradeType_CombinationDerived = '4'

/////////////////////////////////////////////////////////////////////////
///成交价来源类型
/////////////////////////////////////////////////////////////////////////
///前成交价
const PriceSource_LastPrice = '0'
///买委托价
const PriceSource_Buy = '1'
///卖委托价
const PriceSource_Sell = '2'

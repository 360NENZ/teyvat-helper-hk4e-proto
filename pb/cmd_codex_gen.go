// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(CodexDataFullNotify) },
		func() ProtoMessage { return new(CodexDataUpdateNotify) },
		func() ProtoMessage { return new(QueryCodexMonsterBeKilledNumReq) },
		func() ProtoMessage { return new(QueryCodexMonsterBeKilledNumRsp) },
		func() ProtoMessage { return new(ViewCodexReq) },
		func() ProtoMessage { return new(ViewCodexRsp) },
		func() ProtoMessage { return new(SetCodexPushtipsReadReq) },
		func() ProtoMessage { return new(SetCodexPushtipsReadRsp) },
	)
}

const (
	ProtoCommandCodexDataFullNotify             ProtoCommand = 4205
	ProtoCommandCodexDataUpdateNotify           ProtoCommand = 4207
	ProtoCommandQueryCodexMonsterBeKilledNumReq ProtoCommand = 4203
	ProtoCommandQueryCodexMonsterBeKilledNumRsp ProtoCommand = 4209
	ProtoCommandViewCodexReq                    ProtoCommand = 4202
	ProtoCommandViewCodexRsp                    ProtoCommand = 4201
	ProtoCommandSetCodexPushtipsReadReq         ProtoCommand = 4208
	ProtoCommandSetCodexPushtipsReadRsp         ProtoCommand = 4206
)

func (*CodexDataFullNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCodexDataFullNotify }
func (*CodexDataFullNotify) ProtoMessageType() ProtoMessageType { return "CodexDataFullNotify" }

func (*CodexDataUpdateNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCodexDataUpdateNotify }
func (*CodexDataUpdateNotify) ProtoMessageType() ProtoMessageType { return "CodexDataUpdateNotify" }

func (*QueryCodexMonsterBeKilledNumReq) ProtoCommand() ProtoCommand {
	return ProtoCommandQueryCodexMonsterBeKilledNumReq
}
func (*QueryCodexMonsterBeKilledNumReq) ProtoMessageType() ProtoMessageType {
	return "QueryCodexMonsterBeKilledNumReq"
}

func (*QueryCodexMonsterBeKilledNumRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandQueryCodexMonsterBeKilledNumRsp
}
func (*QueryCodexMonsterBeKilledNumRsp) ProtoMessageType() ProtoMessageType {
	return "QueryCodexMonsterBeKilledNumRsp"
}

func (*ViewCodexReq) ProtoCommand() ProtoCommand         { return ProtoCommandViewCodexReq }
func (*ViewCodexReq) ProtoMessageType() ProtoMessageType { return "ViewCodexReq" }

func (*ViewCodexRsp) ProtoCommand() ProtoCommand         { return ProtoCommandViewCodexRsp }
func (*ViewCodexRsp) ProtoMessageType() ProtoMessageType { return "ViewCodexRsp" }

func (*SetCodexPushtipsReadReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSetCodexPushtipsReadReq
}
func (*SetCodexPushtipsReadReq) ProtoMessageType() ProtoMessageType { return "SetCodexPushtipsReadReq" }

func (*SetCodexPushtipsReadRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSetCodexPushtipsReadRsp
}
func (*SetCodexPushtipsReadRsp) ProtoMessageType() ProtoMessageType { return "SetCodexPushtipsReadRsp" }

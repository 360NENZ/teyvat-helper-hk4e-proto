// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(ScenePlayBattleInfoNotify) },
		func() ProtoMessage { return new(ScenePlayOwnerCheckReq) },
		func() ProtoMessage { return new(ScenePlayOwnerCheckRsp) },
		func() ProtoMessage { return new(ScenePlayOwnerStartInviteReq) },
		func() ProtoMessage { return new(ScenePlayOwnerStartInviteRsp) },
		func() ProtoMessage { return new(ScenePlayOwnerInviteNotify) },
		func() ProtoMessage { return new(ScenePlayGuestReplyInviteReq) },
		func() ProtoMessage { return new(ScenePlayGuestReplyInviteRsp) },
		func() ProtoMessage { return new(ScenePlayGuestReplyNotify) },
		func() ProtoMessage { return new(ScenePlayInviteResultNotify) },
		func() ProtoMessage { return new(ScenePlayInfoListNotify) },
		func() ProtoMessage { return new(ScenePlayBattleInterruptNotify) },
		func() ProtoMessage { return new(ScenePlayBattleResultNotify) },
		func() ProtoMessage { return new(ScenePlayBattleUidOpNotify) },
		func() ProtoMessage { return new(ScenePlayBattleInfoListNotify) },
		func() ProtoMessage { return new(ScenePlayOutofRegionNotify) },
	)
}

const (
	ProtoCommandScenePlayBattleInfoNotify      ProtoCommand = 4422
	ProtoCommandScenePlayOwnerCheckReq         ProtoCommand = 4448
	ProtoCommandScenePlayOwnerCheckRsp         ProtoCommand = 4362
	ProtoCommandScenePlayOwnerStartInviteReq   ProtoCommand = 4385
	ProtoCommandScenePlayOwnerStartInviteRsp   ProtoCommand = 4357
	ProtoCommandScenePlayOwnerInviteNotify     ProtoCommand = 4371
	ProtoCommandScenePlayGuestReplyInviteReq   ProtoCommand = 4353
	ProtoCommandScenePlayGuestReplyInviteRsp   ProtoCommand = 4440
	ProtoCommandScenePlayGuestReplyNotify      ProtoCommand = 4423
	ProtoCommandScenePlayInviteResultNotify    ProtoCommand = 4449
	ProtoCommandScenePlayInfoListNotify        ProtoCommand = 4381
	ProtoCommandScenePlayBattleInterruptNotify ProtoCommand = 4425
	ProtoCommandScenePlayBattleResultNotify    ProtoCommand = 4398
	ProtoCommandScenePlayBattleUidOpNotify     ProtoCommand = 4447
	ProtoCommandScenePlayBattleInfoListNotify  ProtoCommand = 4431
	ProtoCommandScenePlayOutofRegionNotify     ProtoCommand = 4355
)

func (*ScenePlayBattleInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayBattleInfoNotify
}
func (*ScenePlayBattleInfoNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayBattleInfoNotify"
}

func (*ScenePlayOwnerCheckReq) ProtoCommand() ProtoCommand         { return ProtoCommandScenePlayOwnerCheckReq }
func (*ScenePlayOwnerCheckReq) ProtoMessageType() ProtoMessageType { return "ScenePlayOwnerCheckReq" }

func (*ScenePlayOwnerCheckRsp) ProtoCommand() ProtoCommand         { return ProtoCommandScenePlayOwnerCheckRsp }
func (*ScenePlayOwnerCheckRsp) ProtoMessageType() ProtoMessageType { return "ScenePlayOwnerCheckRsp" }

func (*ScenePlayOwnerStartInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayOwnerStartInviteReq
}
func (*ScenePlayOwnerStartInviteReq) ProtoMessageType() ProtoMessageType {
	return "ScenePlayOwnerStartInviteReq"
}

func (*ScenePlayOwnerStartInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayOwnerStartInviteRsp
}
func (*ScenePlayOwnerStartInviteRsp) ProtoMessageType() ProtoMessageType {
	return "ScenePlayOwnerStartInviteRsp"
}

func (*ScenePlayOwnerInviteNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayOwnerInviteNotify
}
func (*ScenePlayOwnerInviteNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayOwnerInviteNotify"
}

func (*ScenePlayGuestReplyInviteReq) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayGuestReplyInviteReq
}
func (*ScenePlayGuestReplyInviteReq) ProtoMessageType() ProtoMessageType {
	return "ScenePlayGuestReplyInviteReq"
}

func (*ScenePlayGuestReplyInviteRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayGuestReplyInviteRsp
}
func (*ScenePlayGuestReplyInviteRsp) ProtoMessageType() ProtoMessageType {
	return "ScenePlayGuestReplyInviteRsp"
}

func (*ScenePlayGuestReplyNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayGuestReplyNotify
}
func (*ScenePlayGuestReplyNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayGuestReplyNotify"
}

func (*ScenePlayInviteResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayInviteResultNotify
}
func (*ScenePlayInviteResultNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayInviteResultNotify"
}

func (*ScenePlayInfoListNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayInfoListNotify
}
func (*ScenePlayInfoListNotify) ProtoMessageType() ProtoMessageType { return "ScenePlayInfoListNotify" }

func (*ScenePlayBattleInterruptNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayBattleInterruptNotify
}
func (*ScenePlayBattleInterruptNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayBattleInterruptNotify"
}

func (*ScenePlayBattleResultNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayBattleResultNotify
}
func (*ScenePlayBattleResultNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayBattleResultNotify"
}

func (*ScenePlayBattleUidOpNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayBattleUidOpNotify
}
func (*ScenePlayBattleUidOpNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayBattleUidOpNotify"
}

func (*ScenePlayBattleInfoListNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayBattleInfoListNotify
}
func (*ScenePlayBattleInfoListNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayBattleInfoListNotify"
}

func (*ScenePlayOutofRegionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandScenePlayOutofRegionNotify
}
func (*ScenePlayOutofRegionNotify) ProtoMessageType() ProtoMessageType {
	return "ScenePlayOutofRegionNotify"
}

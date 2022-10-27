// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(EntityPropNotify) },
		func() ProtoMessage { return new(LifeStateChangeNotify) },
		func() ProtoMessage { return new(EntityFightPropNotify) },
		func() ProtoMessage { return new(EntityFightPropUpdateNotify) },
		func() ProtoMessage { return new(AvatarFightPropNotify) },
		func() ProtoMessage { return new(AvatarFightPropUpdateNotify) },
		func() ProtoMessage { return new(EntityFightPropChangeReasonNotify) },
		func() ProtoMessage { return new(AvatarLifeStateChangeNotify) },
		func() ProtoMessage { return new(AvatarPropChangeReasonNotify) },
		func() ProtoMessage { return new(PlayerPropChangeReasonNotify) },
		func() ProtoMessage { return new(AvatarPropNotify) },
		func() ProtoMessage { return new(MarkNewNotify) },
	)
}

const (
	ProtoCommandEntityPropNotify                  ProtoCommand = 1272
	ProtoCommandLifeStateChangeNotify             ProtoCommand = 1298
	ProtoCommandEntityFightPropNotify             ProtoCommand = 1212
	ProtoCommandEntityFightPropUpdateNotify       ProtoCommand = 1235
	ProtoCommandAvatarFightPropNotify             ProtoCommand = 1207
	ProtoCommandAvatarFightPropUpdateNotify       ProtoCommand = 1221
	ProtoCommandEntityFightPropChangeReasonNotify ProtoCommand = 1203
	ProtoCommandAvatarLifeStateChangeNotify       ProtoCommand = 1290
	ProtoCommandAvatarPropChangeReasonNotify      ProtoCommand = 1273
	ProtoCommandPlayerPropChangeReasonNotify      ProtoCommand = 1299
	ProtoCommandAvatarPropNotify                  ProtoCommand = 1231
	ProtoCommandMarkNewNotify                     ProtoCommand = 1275
)

func (*EntityPropNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityPropNotify }
func (*EntityPropNotify) ProtoMessageType() ProtoMessageType { return "EntityPropNotify" }

func (*LifeStateChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLifeStateChangeNotify }
func (*LifeStateChangeNotify) ProtoMessageType() ProtoMessageType { return "LifeStateChangeNotify" }

func (*EntityFightPropNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityFightPropNotify }
func (*EntityFightPropNotify) ProtoMessageType() ProtoMessageType { return "EntityFightPropNotify" }

func (*EntityFightPropUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEntityFightPropUpdateNotify
}
func (*EntityFightPropUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "EntityFightPropUpdateNotify"
}

func (*AvatarFightPropNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarFightPropNotify }
func (*AvatarFightPropNotify) ProtoMessageType() ProtoMessageType { return "AvatarFightPropNotify" }

func (*AvatarFightPropUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarFightPropUpdateNotify
}
func (*AvatarFightPropUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarFightPropUpdateNotify"
}

func (*EntityFightPropChangeReasonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEntityFightPropChangeReasonNotify
}
func (*EntityFightPropChangeReasonNotify) ProtoMessageType() ProtoMessageType {
	return "EntityFightPropChangeReasonNotify"
}

func (*AvatarLifeStateChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarLifeStateChangeNotify
}
func (*AvatarLifeStateChangeNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarLifeStateChangeNotify"
}

func (*AvatarPropChangeReasonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarPropChangeReasonNotify
}
func (*AvatarPropChangeReasonNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarPropChangeReasonNotify"
}

func (*PlayerPropChangeReasonNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandPlayerPropChangeReasonNotify
}
func (*PlayerPropChangeReasonNotify) ProtoMessageType() ProtoMessageType {
	return "PlayerPropChangeReasonNotify"
}

func (*AvatarPropNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarPropNotify }
func (*AvatarPropNotify) ProtoMessageType() ProtoMessageType { return "AvatarPropNotify" }

func (*MarkNewNotify) ProtoCommand() ProtoCommand         { return ProtoCommandMarkNewNotify }
func (*MarkNewNotify) ProtoMessageType() ProtoMessageType { return "MarkNewNotify" }
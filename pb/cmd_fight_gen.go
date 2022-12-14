// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(EvtBeingHitNotify) },
		func() ProtoMessage { return new(EvtAnimatorParameterNotify) },
		func() ProtoMessage { return new(HostPlayerNotify) },
		func() ProtoMessage { return new(EvtDoSkillSuccNotify) },
		func() ProtoMessage { return new(EvtCreateGadgetNotify) },
		func() ProtoMessage { return new(EvtDestroyGadgetNotify) },
		func() ProtoMessage { return new(EvtFaceToEntityNotify) },
		func() ProtoMessage { return new(EvtFaceToDirNotify) },
		func() ProtoMessage { return new(EvtCostStaminaNotify) },
		func() ProtoMessage { return new(EvtSetAttackTargetNotify) },
		func() ProtoMessage { return new(EvtAnimatorStateChangedNotify) },
		func() ProtoMessage { return new(EvtRushMoveNotify) },
		func() ProtoMessage { return new(EvtBulletHitNotify) },
		func() ProtoMessage { return new(EvtBulletDeactiveNotify) },
		func() ProtoMessage { return new(EvtEntityStartDieEndNotify) },
		func() ProtoMessage { return new(EvtBulletMoveNotify) },
		func() ProtoMessage { return new(EvtAvatarEnterFocusNotify) },
		func() ProtoMessage { return new(EvtAvatarExitFocusNotify) },
		func() ProtoMessage { return new(EvtAvatarUpdateFocusNotify) },
		func() ProtoMessage { return new(EntityAuthorityChangeNotify) },
		func() ProtoMessage { return new(AvatarBuffAddNotify) },
		func() ProtoMessage { return new(AvatarBuffDelNotify) },
		func() ProtoMessage { return new(MonsterAlertChangeNotify) },
		func() ProtoMessage { return new(MonsterForceAlertNotify) },
		func() ProtoMessage { return new(AvatarEnterElementViewNotify) },
		func() ProtoMessage { return new(TriggerCreateGadgetToEquipPartNotify) },
		func() ProtoMessage { return new(EvtEntityRenderersChangedNotify) },
		func() ProtoMessage { return new(AnimatorForceSetAirMoveNotify) },
		func() ProtoMessage { return new(EvtAiSyncSkillCdNotify) },
		func() ProtoMessage { return new(EvtBeingHitsCombineNotify) },
		func() ProtoMessage { return new(EvtAvatarSitDownNotify) },
		func() ProtoMessage { return new(EvtAvatarStandUpNotify) },
		func() ProtoMessage { return new(CreateMassiveEntityReq) },
		func() ProtoMessage { return new(CreateMassiveEntityRsp) },
		func() ProtoMessage { return new(CreateMassiveEntityNotify) },
		func() ProtoMessage { return new(DestroyMassiveEntityNotify) },
		func() ProtoMessage { return new(MassiveEntityStateChangedNotify) },
		func() ProtoMessage { return new(SyncTeamEntityNotify) },
		func() ProtoMessage { return new(DelTeamEntityNotify) },
		func() ProtoMessage { return new(CombatInvocationsNotify) },
		func() ProtoMessage { return new(ServerBuffChangeNotify) },
		func() ProtoMessage { return new(EvtAiSyncCombatThreatInfoNotify) },
		func() ProtoMessage { return new(MassiveEntityElementOpBatchNotify) },
		func() ProtoMessage { return new(EntityAiSyncNotify) },
		func() ProtoMessage { return new(LuaSetOptionNotify) },
		func() ProtoMessage { return new(EvtDestroyServerGadgetNotify) },
		func() ProtoMessage { return new(EntityAiKillSelfNotify) },
		func() ProtoMessage { return new(EvtAvatarLockChairReq) },
		func() ProtoMessage { return new(EvtAvatarLockChairRsp) },
		func() ProtoMessage { return new(ReportFightAntiCheatNotify) },
		func() ProtoMessage { return new(EvtBeingHealedNotify) },
		func() ProtoMessage { return new(EvtLocalGadgetOwnerLeaveSceneNotify) },
	)
}

const (
	ProtoCommandEvtBeingHitNotify                    ProtoCommand = 372
	ProtoCommandEvtAnimatorParameterNotify           ProtoCommand = 398
	ProtoCommandHostPlayerNotify                     ProtoCommand = 312
	ProtoCommandEvtDoSkillSuccNotify                 ProtoCommand = 335
	ProtoCommandEvtCreateGadgetNotify                ProtoCommand = 307
	ProtoCommandEvtDestroyGadgetNotify               ProtoCommand = 321
	ProtoCommandEvtFaceToEntityNotify                ProtoCommand = 303
	ProtoCommandEvtFaceToDirNotify                   ProtoCommand = 390
	ProtoCommandEvtCostStaminaNotify                 ProtoCommand = 373
	ProtoCommandEvtSetAttackTargetNotify             ProtoCommand = 399
	ProtoCommandEvtAnimatorStateChangedNotify        ProtoCommand = 331
	ProtoCommandEvtRushMoveNotify                    ProtoCommand = 375
	ProtoCommandEvtBulletHitNotify                   ProtoCommand = 348
	ProtoCommandEvtBulletDeactiveNotify              ProtoCommand = 397
	ProtoCommandEvtEntityStartDieEndNotify           ProtoCommand = 381
	ProtoCommandEvtBulletMoveNotify                  ProtoCommand = 365
	ProtoCommandEvtAvatarEnterFocusNotify            ProtoCommand = 304
	ProtoCommandEvtAvatarExitFocusNotify             ProtoCommand = 393
	ProtoCommandEvtAvatarUpdateFocusNotify           ProtoCommand = 327
	ProtoCommandEntityAuthorityChangeNotify          ProtoCommand = 394
	ProtoCommandAvatarBuffAddNotify                  ProtoCommand = 388
	ProtoCommandAvatarBuffDelNotify                  ProtoCommand = 326
	ProtoCommandMonsterAlertChangeNotify             ProtoCommand = 363
	ProtoCommandMonsterForceAlertNotify              ProtoCommand = 395
	ProtoCommandAvatarEnterElementViewNotify         ProtoCommand = 334
	ProtoCommandTriggerCreateGadgetToEquipPartNotify ProtoCommand = 350
	ProtoCommandEvtEntityRenderersChangedNotify      ProtoCommand = 343
	ProtoCommandAnimatorForceSetAirMoveNotify        ProtoCommand = 374
	ProtoCommandEvtAiSyncSkillCdNotify               ProtoCommand = 376
	ProtoCommandEvtBeingHitsCombineNotify            ProtoCommand = 346
	ProtoCommandEvtAvatarSitDownNotify               ProtoCommand = 324
	ProtoCommandEvtAvatarStandUpNotify               ProtoCommand = 356
	ProtoCommandCreateMassiveEntityReq               ProtoCommand = 342
	ProtoCommandCreateMassiveEntityRsp               ProtoCommand = 330
	ProtoCommandCreateMassiveEntityNotify            ProtoCommand = 367
	ProtoCommandDestroyMassiveEntityNotify           ProtoCommand = 358
	ProtoCommandMassiveEntityStateChangedNotify      ProtoCommand = 370
	ProtoCommandSyncTeamEntityNotify                 ProtoCommand = 317
	ProtoCommandDelTeamEntityNotify                  ProtoCommand = 302
	ProtoCommandCombatInvocationsNotify              ProtoCommand = 319
	ProtoCommandServerBuffChangeNotify               ProtoCommand = 361
	ProtoCommandEvtAiSyncCombatThreatInfoNotify      ProtoCommand = 329
	ProtoCommandMassiveEntityElementOpBatchNotify    ProtoCommand = 357
	ProtoCommandEntityAiSyncNotify                   ProtoCommand = 400
	ProtoCommandLuaSetOptionNotify                   ProtoCommand = 316
	ProtoCommandEvtDestroyServerGadgetNotify         ProtoCommand = 387
	ProtoCommandEntityAiKillSelfNotify               ProtoCommand = 340
	ProtoCommandEvtAvatarLockChairReq                ProtoCommand = 318
	ProtoCommandEvtAvatarLockChairRsp                ProtoCommand = 366
	ProtoCommandReportFightAntiCheatNotify           ProtoCommand = 368
	ProtoCommandEvtBeingHealedNotify                 ProtoCommand = 333
	ProtoCommandEvtLocalGadgetOwnerLeaveSceneNotify  ProtoCommand = 384
)

func (*EvtBeingHitNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtBeingHitNotify }
func (*EvtBeingHitNotify) ProtoMessageType() ProtoMessageType { return "EvtBeingHitNotify" }

func (*EvtAnimatorParameterNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtAnimatorParameterNotify
}
func (*EvtAnimatorParameterNotify) ProtoMessageType() ProtoMessageType {
	return "EvtAnimatorParameterNotify"
}

func (*HostPlayerNotify) ProtoCommand() ProtoCommand         { return ProtoCommandHostPlayerNotify }
func (*HostPlayerNotify) ProtoMessageType() ProtoMessageType { return "HostPlayerNotify" }

func (*EvtDoSkillSuccNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtDoSkillSuccNotify }
func (*EvtDoSkillSuccNotify) ProtoMessageType() ProtoMessageType { return "EvtDoSkillSuccNotify" }

func (*EvtCreateGadgetNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtCreateGadgetNotify }
func (*EvtCreateGadgetNotify) ProtoMessageType() ProtoMessageType { return "EvtCreateGadgetNotify" }

func (*EvtDestroyGadgetNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtDestroyGadgetNotify }
func (*EvtDestroyGadgetNotify) ProtoMessageType() ProtoMessageType { return "EvtDestroyGadgetNotify" }

func (*EvtFaceToEntityNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtFaceToEntityNotify }
func (*EvtFaceToEntityNotify) ProtoMessageType() ProtoMessageType { return "EvtFaceToEntityNotify" }

func (*EvtFaceToDirNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtFaceToDirNotify }
func (*EvtFaceToDirNotify) ProtoMessageType() ProtoMessageType { return "EvtFaceToDirNotify" }

func (*EvtCostStaminaNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtCostStaminaNotify }
func (*EvtCostStaminaNotify) ProtoMessageType() ProtoMessageType { return "EvtCostStaminaNotify" }

func (*EvtSetAttackTargetNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtSetAttackTargetNotify
}
func (*EvtSetAttackTargetNotify) ProtoMessageType() ProtoMessageType {
	return "EvtSetAttackTargetNotify"
}

func (*EvtAnimatorStateChangedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtAnimatorStateChangedNotify
}
func (*EvtAnimatorStateChangedNotify) ProtoMessageType() ProtoMessageType {
	return "EvtAnimatorStateChangedNotify"
}

func (*EvtRushMoveNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtRushMoveNotify }
func (*EvtRushMoveNotify) ProtoMessageType() ProtoMessageType { return "EvtRushMoveNotify" }

func (*EvtBulletHitNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtBulletHitNotify }
func (*EvtBulletHitNotify) ProtoMessageType() ProtoMessageType { return "EvtBulletHitNotify" }

func (*EvtBulletDeactiveNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtBulletDeactiveNotify
}
func (*EvtBulletDeactiveNotify) ProtoMessageType() ProtoMessageType { return "EvtBulletDeactiveNotify" }

func (*EvtEntityStartDieEndNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtEntityStartDieEndNotify
}
func (*EvtEntityStartDieEndNotify) ProtoMessageType() ProtoMessageType {
	return "EvtEntityStartDieEndNotify"
}

func (*EvtBulletMoveNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtBulletMoveNotify }
func (*EvtBulletMoveNotify) ProtoMessageType() ProtoMessageType { return "EvtBulletMoveNotify" }

func (*EvtAvatarEnterFocusNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtAvatarEnterFocusNotify
}
func (*EvtAvatarEnterFocusNotify) ProtoMessageType() ProtoMessageType {
	return "EvtAvatarEnterFocusNotify"
}

func (*EvtAvatarExitFocusNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtAvatarExitFocusNotify
}
func (*EvtAvatarExitFocusNotify) ProtoMessageType() ProtoMessageType {
	return "EvtAvatarExitFocusNotify"
}

func (*EvtAvatarUpdateFocusNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtAvatarUpdateFocusNotify
}
func (*EvtAvatarUpdateFocusNotify) ProtoMessageType() ProtoMessageType {
	return "EvtAvatarUpdateFocusNotify"
}

func (*EntityAuthorityChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEntityAuthorityChangeNotify
}
func (*EntityAuthorityChangeNotify) ProtoMessageType() ProtoMessageType {
	return "EntityAuthorityChangeNotify"
}

func (*AvatarBuffAddNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarBuffAddNotify }
func (*AvatarBuffAddNotify) ProtoMessageType() ProtoMessageType { return "AvatarBuffAddNotify" }

func (*AvatarBuffDelNotify) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarBuffDelNotify }
func (*AvatarBuffDelNotify) ProtoMessageType() ProtoMessageType { return "AvatarBuffDelNotify" }

func (*MonsterAlertChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMonsterAlertChangeNotify
}
func (*MonsterAlertChangeNotify) ProtoMessageType() ProtoMessageType {
	return "MonsterAlertChangeNotify"
}

func (*MonsterForceAlertNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMonsterForceAlertNotify
}
func (*MonsterForceAlertNotify) ProtoMessageType() ProtoMessageType { return "MonsterForceAlertNotify" }

func (*AvatarEnterElementViewNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarEnterElementViewNotify
}
func (*AvatarEnterElementViewNotify) ProtoMessageType() ProtoMessageType {
	return "AvatarEnterElementViewNotify"
}

func (*TriggerCreateGadgetToEquipPartNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandTriggerCreateGadgetToEquipPartNotify
}
func (*TriggerCreateGadgetToEquipPartNotify) ProtoMessageType() ProtoMessageType {
	return "TriggerCreateGadgetToEquipPartNotify"
}

func (*EvtEntityRenderersChangedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtEntityRenderersChangedNotify
}
func (*EvtEntityRenderersChangedNotify) ProtoMessageType() ProtoMessageType {
	return "EvtEntityRenderersChangedNotify"
}

func (*AnimatorForceSetAirMoveNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAnimatorForceSetAirMoveNotify
}
func (*AnimatorForceSetAirMoveNotify) ProtoMessageType() ProtoMessageType {
	return "AnimatorForceSetAirMoveNotify"
}

func (*EvtAiSyncSkillCdNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtAiSyncSkillCdNotify }
func (*EvtAiSyncSkillCdNotify) ProtoMessageType() ProtoMessageType { return "EvtAiSyncSkillCdNotify" }

func (*EvtBeingHitsCombineNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtBeingHitsCombineNotify
}
func (*EvtBeingHitsCombineNotify) ProtoMessageType() ProtoMessageType {
	return "EvtBeingHitsCombineNotify"
}

func (*EvtAvatarSitDownNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtAvatarSitDownNotify }
func (*EvtAvatarSitDownNotify) ProtoMessageType() ProtoMessageType { return "EvtAvatarSitDownNotify" }

func (*EvtAvatarStandUpNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtAvatarStandUpNotify }
func (*EvtAvatarStandUpNotify) ProtoMessageType() ProtoMessageType { return "EvtAvatarStandUpNotify" }

func (*CreateMassiveEntityReq) ProtoCommand() ProtoCommand         { return ProtoCommandCreateMassiveEntityReq }
func (*CreateMassiveEntityReq) ProtoMessageType() ProtoMessageType { return "CreateMassiveEntityReq" }

func (*CreateMassiveEntityRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCreateMassiveEntityRsp }
func (*CreateMassiveEntityRsp) ProtoMessageType() ProtoMessageType { return "CreateMassiveEntityRsp" }

func (*CreateMassiveEntityNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCreateMassiveEntityNotify
}
func (*CreateMassiveEntityNotify) ProtoMessageType() ProtoMessageType {
	return "CreateMassiveEntityNotify"
}

func (*DestroyMassiveEntityNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandDestroyMassiveEntityNotify
}
func (*DestroyMassiveEntityNotify) ProtoMessageType() ProtoMessageType {
	return "DestroyMassiveEntityNotify"
}

func (*MassiveEntityStateChangedNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMassiveEntityStateChangedNotify
}
func (*MassiveEntityStateChangedNotify) ProtoMessageType() ProtoMessageType {
	return "MassiveEntityStateChangedNotify"
}

func (*SyncTeamEntityNotify) ProtoCommand() ProtoCommand         { return ProtoCommandSyncTeamEntityNotify }
func (*SyncTeamEntityNotify) ProtoMessageType() ProtoMessageType { return "SyncTeamEntityNotify" }

func (*DelTeamEntityNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDelTeamEntityNotify }
func (*DelTeamEntityNotify) ProtoMessageType() ProtoMessageType { return "DelTeamEntityNotify" }

func (*CombatInvocationsNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCombatInvocationsNotify
}
func (*CombatInvocationsNotify) ProtoMessageType() ProtoMessageType { return "CombatInvocationsNotify" }

func (*ServerBuffChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandServerBuffChangeNotify }
func (*ServerBuffChangeNotify) ProtoMessageType() ProtoMessageType { return "ServerBuffChangeNotify" }

func (*EvtAiSyncCombatThreatInfoNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtAiSyncCombatThreatInfoNotify
}
func (*EvtAiSyncCombatThreatInfoNotify) ProtoMessageType() ProtoMessageType {
	return "EvtAiSyncCombatThreatInfoNotify"
}

func (*MassiveEntityElementOpBatchNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMassiveEntityElementOpBatchNotify
}
func (*MassiveEntityElementOpBatchNotify) ProtoMessageType() ProtoMessageType {
	return "MassiveEntityElementOpBatchNotify"
}

func (*EntityAiSyncNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityAiSyncNotify }
func (*EntityAiSyncNotify) ProtoMessageType() ProtoMessageType { return "EntityAiSyncNotify" }

func (*LuaSetOptionNotify) ProtoCommand() ProtoCommand         { return ProtoCommandLuaSetOptionNotify }
func (*LuaSetOptionNotify) ProtoMessageType() ProtoMessageType { return "LuaSetOptionNotify" }

func (*EvtDestroyServerGadgetNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtDestroyServerGadgetNotify
}
func (*EvtDestroyServerGadgetNotify) ProtoMessageType() ProtoMessageType {
	return "EvtDestroyServerGadgetNotify"
}

func (*EntityAiKillSelfNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEntityAiKillSelfNotify }
func (*EntityAiKillSelfNotify) ProtoMessageType() ProtoMessageType { return "EntityAiKillSelfNotify" }

func (*EvtAvatarLockChairReq) ProtoCommand() ProtoCommand         { return ProtoCommandEvtAvatarLockChairReq }
func (*EvtAvatarLockChairReq) ProtoMessageType() ProtoMessageType { return "EvtAvatarLockChairReq" }

func (*EvtAvatarLockChairRsp) ProtoCommand() ProtoCommand         { return ProtoCommandEvtAvatarLockChairRsp }
func (*EvtAvatarLockChairRsp) ProtoMessageType() ProtoMessageType { return "EvtAvatarLockChairRsp" }

func (*ReportFightAntiCheatNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandReportFightAntiCheatNotify
}
func (*ReportFightAntiCheatNotify) ProtoMessageType() ProtoMessageType {
	return "ReportFightAntiCheatNotify"
}

func (*EvtBeingHealedNotify) ProtoCommand() ProtoCommand         { return ProtoCommandEvtBeingHealedNotify }
func (*EvtBeingHealedNotify) ProtoMessageType() ProtoMessageType { return "EvtBeingHealedNotify" }

func (*EvtLocalGadgetOwnerLeaveSceneNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandEvtLocalGadgetOwnerLeaveSceneNotify
}
func (*EvtLocalGadgetOwnerLeaveSceneNotify) ProtoMessageType() ProtoMessageType {
	return "EvtLocalGadgetOwnerLeaveSceneNotify"
}

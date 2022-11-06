#!/usr/bin/env bash

set -eux

protoc --go_out=$GOPATH/src --proto_path=proto define.proto
protoc --go_out=$GOPATH/src --proto_path=proto editor.proto
protoc --go_out=$GOPATH/src --proto_path=proto retcode.proto

protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_ability.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_achievement.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_activity.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_aranara_collection.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_avatar.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_battle_pass.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_blossom.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_chat.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_codex.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_coop.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_custom_dungeon.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_draft.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_dungeon.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_fight.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_fish.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_gacha.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_gadget.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_gallery.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_gcg.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_gcg_common.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_group_link.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_h5_activity.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_home.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_hunting.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_investigation.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_item.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_mail.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_match.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_mechanicus.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_miracle_ring.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_misc.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_monster.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_mp.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_multistage_play.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_npc.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_offering.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_op_activity.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_pathfinding.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_player.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_property.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_quest.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_recharge.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_region_search.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_regional_play.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_reminder.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_reputation.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_reunion.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_routine.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_scene.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_scene_play.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_share_cd.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_shop.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_sign_in.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_skill.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_social.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_stat.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_tothemoon.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_tower.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_ugc.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_watcher.proto
protoc --go_out=$GOPATH/src --proto_path=proto cmd/cmd_widget.proto

protoc --go_out=$GOPATH/src --proto_path=proto pb/packet_head.proto

protoc --go_out=$GOPATH/src --proto_path=proto server_only/bin.block.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/bin.home.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/bin.multiserver.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/bin.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/bin_common.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_activity.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_gcg.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_home.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_id_config.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_mail.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_match.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_misc.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_mp.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_muip.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_offline_op.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_player.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_recharge.server.proto
# protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_security.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/cmd_social.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/config.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/enum.server.proto
protoc --go_out=$GOPATH/src --proto_path=proto server_only/redis_data.proto
# protoc --go_out=$GOPATH/src --proto_path=proto server_only/SVOProto.proto

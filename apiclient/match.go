package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Tilo-K/riot/constants/champion"
	"github.com/Tilo-K/riot/constants/event"
	"github.com/Tilo-K/riot/constants/lane"
	"github.com/Tilo-K/riot/constants/queue"
	"github.com/Tilo-K/riot/constants/region"
	"github.com/Tilo-K/riot/constants/season"
	"github.com/Tilo-K/riot/constants/v5region"
	"github.com/Tilo-K/riot/types"
)

type Match struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		GameCreation       int64  `json:"gameCreation"`
		GameDuration       int    `json:"gameDuration"`
		GameEndTimestamp   int64  `json:"gameEndTimestamp"`
		GameID             int64  `json:"gameId"`
		GameMode           string `json:"gameMode"`
		GameName           string `json:"gameName"`
		GameStartTimestamp int64  `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapID              int    `json:"mapId"`
		Participants       []struct {
			AllInPings    int `json:"allInPings"`
			AssistMePings int `json:"assistMePings"`
			Assists       int `json:"assists"`
			BaitPings     int `json:"baitPings"`
			BaronKills    int `json:"baronKills"`
			BasicPings    int `json:"basicPings"`
			BountyLevel   int `json:"bountyLevel"`
			Challenges    struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronBuffGoldAdvantageOverThreshold      int     `json:"baronBuffGoldAdvantageOverThreshold"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarliestBaron                            float64 `json:"earliestBaron"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  float64 `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                            int     `json:"killingSprees"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                           int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			ChampExperience             int    `json:"champExperience"`
			ChampLevel                  int    `json:"champLevel"`
			ChampionID                  int    `json:"championId"`
			ChampionName                string `json:"championName"`
			ChampionTransform           int    `json:"championTransform"`
			CommandPings                int    `json:"commandPings"`
			ConsumablesPurchased        int    `json:"consumablesPurchased"`
			DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
			DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
			DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
			DamageSelfMitigated         int    `json:"damageSelfMitigated"`
			DangerPings                 int    `json:"dangerPings"`
			Deaths                      int    `json:"deaths"`
			DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
			DoubleKills                 int    `json:"doubleKills"`
			DragonKills                 int    `json:"dragonKills"`
			EligibleForProgression      bool   `json:"eligibleForProgression"`
			EnemyMissingPings           int    `json:"enemyMissingPings"`
			EnemyVisionPings            int    `json:"enemyVisionPings"`
			FirstBloodAssist            bool   `json:"firstBloodAssist"`
			FirstBloodKill              bool   `json:"firstBloodKill"`
			FirstTowerAssist            bool   `json:"firstTowerAssist"`
			FirstTowerKill              bool   `json:"firstTowerKill"`
			GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
			GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
			GetBackPings                int    `json:"getBackPings"`
			GoldEarned                  int    `json:"goldEarned"`
			GoldSpent                   int    `json:"goldSpent"`
			HoldPings                   int    `json:"holdPings"`
			IndividualPosition          string `json:"individualPosition"`
			InhibitorKills              int    `json:"inhibitorKills"`
			InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
			InhibitorsLost              int    `json:"inhibitorsLost"`
			Item0                       int    `json:"item0"`
			Item1                       int    `json:"item1"`
			Item2                       int    `json:"item2"`
			Item3                       int    `json:"item3"`
			Item4                       int    `json:"item4"`
			Item5                       int    `json:"item5"`
			Item6                       int    `json:"item6"`
			ItemsPurchased              int    `json:"itemsPurchased"`
			KillingSprees               int    `json:"killingSprees"`
			Kills                       int    `json:"kills"`
			Lane                        string `json:"lane"`
			LargestCriticalStrike       int    `json:"largestCriticalStrike"`
			LargestKillingSpree         int    `json:"largestKillingSpree"`
			LargestMultiKill            int    `json:"largestMultiKill"`
			LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
			MagicDamageDealt            int    `json:"magicDamageDealt"`
			MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
			MagicDamageTaken            int    `json:"magicDamageTaken"`
			NeedVisionPings             int    `json:"needVisionPings"`
			NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
			NexusKills                  int    `json:"nexusKills"`
			NexusLost                   int    `json:"nexusLost"`
			NexusTakedowns              int    `json:"nexusTakedowns"`
			ObjectivesStolen            int    `json:"objectivesStolen"`
			ObjectivesStolenAssists     int    `json:"objectivesStolenAssists"`
			OnMyWayPings                int    `json:"onMyWayPings"`
			ParticipantID               int    `json:"participantId"`
			PentaKills                  int    `json:"pentaKills"`
			Perks                       struct {
				StatPerks struct {
					Defense int `json:"defense"`
					Flex    int `json:"flex"`
					Offense int `json:"offense"`
				} `json:"statPerks"`
				Styles []struct {
					Description string `json:"description"`
					Selections  []struct {
						Perk int `json:"perk"`
						Var1 int `json:"var1"`
						Var2 int `json:"var2"`
						Var3 int `json:"var3"`
					} `json:"selections"`
					Style int `json:"style"`
				} `json:"styles"`
			} `json:"perks"`
			PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
			ProfileIcon                    int    `json:"profileIcon"`
			PushPings                      int    `json:"pushPings"`
			Puuid                          string `json:"puuid"`
			QuadraKills                    int    `json:"quadraKills"`
			RiotIDName                     string `json:"riotIdName"`
			RiotIDTagline                  string `json:"riotIdTagline"`
			Role                           string `json:"role"`
			SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
			Spell1Casts                    int    `json:"spell1Casts"`
			Spell2Casts                    int    `json:"spell2Casts"`
			Spell3Casts                    int    `json:"spell3Casts"`
			Spell4Casts                    int    `json:"spell4Casts"`
			Summoner1Casts                 int    `json:"summoner1Casts"`
			Summoner1ID                    int    `json:"summoner1Id"`
			Summoner2Casts                 int    `json:"summoner2Casts"`
			Summoner2ID                    int    `json:"summoner2Id"`
			SummonerID                     string `json:"summonerId"`
			SummonerLevel                  int    `json:"summonerLevel"`
			SummonerName                   string `json:"summonerName"`
			TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
			TeamID                         int    `json:"teamId"`
			TeamPosition                   string `json:"teamPosition"`
			TimeCCingOthers                int    `json:"timeCCingOthers"`
			TimePlayed                     int    `json:"timePlayed"`
			TotalDamageDealt               int    `json:"totalDamageDealt"`
			TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
			TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
			TotalDamageTaken               int    `json:"totalDamageTaken"`
			TotalHeal                      int    `json:"totalHeal"`
			TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
			TotalMinionsKilled             int    `json:"totalMinionsKilled"`
			TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
			TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
			TotalUnitsHealed               int    `json:"totalUnitsHealed"`
			TripleKills                    int    `json:"tripleKills"`
			TrueDamageDealt                int    `json:"trueDamageDealt"`
			TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                int    `json:"trueDamageTaken"`
			TurretKills                    int    `json:"turretKills"`
			TurretTakedowns                int    `json:"turretTakedowns"`
			TurretsLost                    int    `json:"turretsLost"`
			UnrealKills                    int    `json:"unrealKills"`
			VisionClearedPings             int    `json:"visionClearedPings"`
			VisionScore                    int    `json:"visionScore"`
			VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
			WardsKilled                    int    `json:"wardsKilled"`
			WardsPlaced                    int    `json:"wardsPlaced"`
			Win                            bool   `json:"win"`
			Challenges0                    struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  float64 `json:"alliedJungleMonsterKills"`
				BaronBuffGoldAdvantageOverThreshold       int     `json:"baronBuffGoldAdvantageOverThreshold"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestBaron                             float64 `json:"earliestBaron"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   float64 `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes"`
				JunglerKillsEarlyJungle                   int     `json:"junglerKillsEarlyJungle"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       int     `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                             int     `json:"killingSprees"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnLanersEarlyJungleAsJungler         int     `json:"killsOnLanersEarlyJungleAsJungler"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				ShortestTimeToAceFromFirstTakedown        float64 `json:"shortestTimeToAceFromFirstTakedown"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges1 struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 float64 `json:"alliedJungleMonsterKills"`
				BaronBuffGoldAdvantageOverThreshold      int     `json:"baronBuffGoldAdvantageOverThreshold"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarliestBaron                            float64 `json:"earliestBaron"`
				EarliestDragonTakedown                   float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                            int     `json:"killingSprees"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                           int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges2 struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronBuffGoldAdvantageOverThreshold      int     `json:"baronBuffGoldAdvantageOverThreshold"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarliestBaron                            float64 `json:"earliestBaron"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                float64 `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  float64 `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				HighestCrowdControlScore                 int     `json:"highestCrowdControlScore"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  float64 `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      int     `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                           int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				ShortestTimeToAceFromFirstTakedown       float64 `json:"shortestTimeToAceFromFirstTakedown"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges3 struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronBuffGoldAdvantageOverThreshold      int     `json:"baronBuffGoldAdvantageOverThreshold"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarliestBaron                            float64 `json:"earliestBaron"`
				EarliestDragonTakedown                   float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FasterSupportQuestCompletion             int     `json:"fasterSupportQuestCompletion"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                            int     `json:"killingSprees"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				ShortestTimeToAceFromFirstTakedown       float64 `json:"shortestTimeToAceFromFirstTakedown"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges4 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  float64 `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				HighestChampionDamage                     int     `json:"highestChampionDamage"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       int     `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                             int     `json:"killingSprees"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				ShortestTimeToAceFromFirstTakedown        float64 `json:"shortestTimeToAceFromFirstTakedown"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				SoloTurretsLategame                       int     `json:"soloTurretsLategame"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges5 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  float64 `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   float64 `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				HighestWardKills                          int     `json:"highestWardKills"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes"`
				JunglerKillsEarlyJungle                   int     `json:"junglerKillsEarlyJungle"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       float64 `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnLanersEarlyJungleAsJungler         int     `json:"killsOnLanersEarlyJungleAsJungler"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges6 struct {
				One2AssistStreakCount                    int     `json:"12AssistStreakCount"`
				AbilityUses                              int     `json:"abilityUses"`
				AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                 int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                           int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                               int     `json:"bountyGold"`
				BuffsStolen                              int     `json:"buffsStolen"`
				CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
				DamagePerMinute                          float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                               int     `json:"doubleAces"`
				DragonTakedowns                          int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                   float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage         int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                int     `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                  int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilledTime                    float64 `json:"firstTurretKilledTime"`
				FlawlessAces                             int     `json:"flawlessAces"`
				FullTeamTakedown                         int     `json:"fullTeamTakedown"`
				GameLength                               float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                            float64 `json:"goldPerMinute"`
				HadOpenNexus                             int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                         int     `json:"initialBuffCount"`
				InitialCrabCount                         int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                  int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                      float64 `json:"kda"`
				KillAfterHiddenWithAlly                  int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                        float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived    int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                            int     `json:"killingSprees"`
				KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage              int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                           int     `json:"legendaryCount"`
				LostAnInhibitor                          int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent             float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                           int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent              int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
				Multikills                               int     `json:"multikills"`
				MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                           int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                         int     `json:"outnumberedKills"`
				OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                              int     `json:"perfectGame"`
				PickKillWithAlly                         int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                int     `json:"playedChampSelectPosition"`
				PoroExplosions                           int     `json:"poroExplosions"`
				QuickCleanse                             int     `json:"quickCleanse"`
				QuickFirstTurret                         int     `json:"quickFirstTurret"`
				QuickSoloKills                           int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                        int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
				SkillshotsDodged                         int     `json:"skillshotsDodged"`
				SkillshotsHit                            int     `json:"skillshotsHit"`
				SnowballsHit                             int     `json:"snowballsHit"`
				SoloBaronKills                           int     `json:"soloBaronKills"`
				SoloKills                                int     `json:"soloKills"`
				StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight          int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
				Takedowns                                int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                   int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                           int     `json:"teamBaronKills"`
				TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                  int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
				TurretTakedowns                          int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                            int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
				WardTakedowns                            int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                             int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges7 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       float64 `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                             int     `json:"killingSprees"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
			Challenges8 struct {
				One2AssistStreakCount                     int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   int     `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       int     `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              int     `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MoreEnemyJungleThanOpponent               int     `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				MythicItemUsed                            int     `json:"mythicItemUsed"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
			} `json:"challenges,omitempty"`
		} `json:"participants"`
		PlatformID string `json:"platformId"`
		QueueID    int    `json:"queueId"`
		Teams      []struct {
			Bans []struct {
				ChampionID int `json:"championId"`
				PickTurn   int `json:"pickTurn"`
			} `json:"bans"`
			Objectives struct {
				Baron struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"baron"`
				Champion struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"champion"`
				Dragon struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"dragon"`
				Inhibitor struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"inhibitor"`
				RiftHerald struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"riftHerald"`
				Tower struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"tower"`
			} `json:"objectives"`
			TeamID int  `json:"teamId"`
			Win    bool `json:"win"`
		} `json:"teams"`
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}

type Rune struct {
	RuneID int `json:"runeID"`
	Rank   int `json:"rank"`
}

// Interval represents a range of game time, measured in minutes.
//
// The value 999 is used to represent an endpoint that is coded as the literal
// string "end"
type Interval struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}

// IntervalValues represents a mapping from intervals to values.
type IntervalValues []IntervalValue

func (i *IntervalValues) UnmarshalJSON(b []byte) error {
	var obj map[string]float64
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}
	var vals []IntervalValue
	for k, v := range obj {
		intervals := strings.Split(k, "-")
		if len(intervals) != 2 {
			return fmt.Errorf("unable to parse intervals: %v", intervals)
		}
		intervals[0] = strings.TrimSpace(intervals[0])
		intervals[1] = strings.TrimSpace(intervals[1])

		begin, err := strconv.ParseInt(intervals[0], 10, 64)
		if err != nil {
			return err
		}
		var end int64
		if intervals[1] == "end" {
			end = 999
		} else {
			end, err = strconv.ParseInt(intervals[1], 10, 64)
			if err != nil {
				return err
			}
		}
		vals = append(vals, IntervalValue{
			Interval: Interval{
				Begin: int(begin),
				End:   int(end),
			},
			Value: v,
		})
	}
	*i = IntervalValues(vals)
	return nil
}

type IntervalValue struct {
	Interval Interval `json:"interval"`
	Value    float64  `json:"value"`
}

type ParticipantTimeline struct {
	Lane                        lane.Lane      `json:"lane"`
	ParticipantID               int            `json:"participantID"`
	CSDiffPerMinDeltas          IntervalValues `json:"cSDiffPerMinDeltas"`
	GoldPerMinDeltas            IntervalValues `json:"goldPerMinDeltas"`
	XPDiffPerMinDeltas          IntervalValues `json:"xPDiffPerMinDeltas"`
	CreepsPerMinDeltas          IntervalValues `json:"creepsPerMinDeltas"`
	XPPerMinDeltas              IntervalValues `json:"xPPerMinDeltas"`
	Role                        string         `json:"role"`
	DamageTakenDiffPerMinDeltas IntervalValues `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     IntervalValues `json:"damageTakenPerMinDeltas"`
}

type Mastery struct {
	MasteryID int `json:"masteryID"`
	Rank      int `json:"rank"`
}

func (c *client) GetMatch(ctx context.Context, r v5region.V5Region, matchID string) (*Match, error) {
	var res Match
	_, err := c.dispatchAndUnmarshalV5(ctx, r, "/lol/match/v5/matches", fmt.Sprintf("/%s", matchID), nil, &res)
	return &res, err
}

type Matchlist struct {
	Matches    []MatchReference `json:"matches",datastore:",noindex"`
	TotalGames int              `json:"totalGames",datastore:",noindex"`
	StartIndex int              `json:"startIndex",datastore:",noindex"`
	EndIndex   int              `json:"endIndex",datastore:",noindex"`
}

type MatchReference struct {
	Lane       lane.Lane          `json:"lane"`
	GameID     int64              `json:"gameID"`
	Champion   champion.Champion  `json:"champion"`
	PlatformID string             `json:"platformID"`
	Season     season.Season      `json:"season"`
	Queue      queue.Queue        `json:"queue"`
	Role       string             `json:"role"`
	Timestamp  types.Milliseconds `json:"timestamp"`
}

// GetMatchlistOptions provides filtering options for GetMatchlist. The zero
// value means that the option will not be used in filtering.
type GetMatchlistOptions struct {
	Queue      []queue.Queue       `json:"queue"`
	Season     []season.Season     `json:"season"`
	Champion   []champion.Champion `json:"champion"`
	BeginTime  *time.Time          `json:"beginTime"`
	EndTime    *time.Time          `json:"endTime"`
	BeginIndex *int                `json:"beginIndex"`
	EndIndex   *int                `json:"endIndex"`
}

type GetMatchIdsOptions struct {
	StartTime int64  `json:"beginTime"`
	EndTime   int64  `json:"endTime"`
	Queue     int    `json:"queue"`
	Type      string `json:"type"`
	Start     int    `json:"start"`
	Count     int    `json:"count"`
}

func timeToUnixMilliseconds(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond/time.Nanosecond)
}

func (c *client) GetMatchIds(ctx context.Context, r v5region.V5Region, PUUID string, opts *GetMatchIdsOptions) ([]string, error) {
	var (
		res  []string
		vals url.Values
	)

	if opts != nil {
		vals = url.Values(make(map[string][]string))
		if opts.Queue != 0 {
			vals.Add("queue", fmt.Sprintf("%d", opts.Queue))
		}
		if opts.StartTime != 0 {
			vals.Add("startTime", fmt.Sprintf("%d", opts.StartTime))
		}
		if opts.EndTime != 0 {
			vals.Add("endTime", fmt.Sprintf("%d", opts.EndTime))
		}
		if opts.Start != 0 {
			vals.Add("beginIndex", fmt.Sprintf("%d", opts.Start))
		}
		if opts.Count != 0 {
			vals.Add("count", fmt.Sprintf("%d", opts.Count))
		}
	}
	_, err := c.dispatchAndUnmarshalV5(ctx, r, "/lol/match/v5/matches/by-puuid", fmt.Sprintf("/%s/ids", PUUID), vals, &res)
	return res, err
}

func (c *client) GetMatchlist(ctx context.Context, r region.Region, accountID string, opts *GetMatchlistOptions) (*Matchlist, error) {
	var (
		res  Matchlist
		vals url.Values
	)

	if opts != nil {
		vals = url.Values(make(map[string][]string))
		if len(opts.Queue) != 0 {
			for _, v := range opts.Queue {
				vals.Add("queue", fmt.Sprintf("%d", v))
			}
		}
		if len(opts.Season) != 0 {
			for _, v := range opts.Season {
				vals.Add("season", fmt.Sprintf("%d", v))
			}
		}
		if len(opts.Champion) != 0 {
			for _, v := range opts.Champion {
				vals.Add("champion", fmt.Sprintf("%d", v))
			}
		}
		if opts.BeginTime != nil {
			vals.Add("beginTime", fmt.Sprintf("%d", timeToUnixMilliseconds(*opts.BeginTime)))
		}
		if opts.EndTime != nil {
			vals.Add("endTime", fmt.Sprintf("%d", timeToUnixMilliseconds(*opts.EndTime)))
		}
		if opts.BeginIndex != nil {
			vals.Add("beginIndex", fmt.Sprintf("%d", *opts.BeginIndex))
		}
		if opts.EndIndex != nil {
			vals.Add("endIndex", fmt.Sprintf("%d", *opts.EndIndex))
		}
	}
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/match/v4/matchlists/by-account", fmt.Sprintf("/%s", accountID), vals, &res)
	return &res, err
}

func (c *client) GetRecentMatchlist(ctx context.Context, r region.Region, accountID string) (*Matchlist, error) {
	var res Matchlist
	// Recent matchlists are a separate API call from matchlists, even though
	// both have the same Method. Add "recent" as a uniquifier for this method.
	_, err := c.dispatchAndUnmarshalWithUniquifier(ctx, r, "/lol/match/v4/matchlists/by-account", fmt.Sprintf("/%s/recent", accountID), nil, "recent", &res)
	return &res, err
}

type MatchTimeline struct {
	Frames        []MatchFrame       `json:"frames",datastore:",noindex"`
	FrameInterval types.Milliseconds `json:"frameInterval",datastore:",noindex"`
}

// ParticipantFrames stores frames corresponding to each participant. The order
// is not defined (i.e. do not assume the order is ascending by participant ID).
type ParticipantFrames struct {
	Frames []MatchParticipantFrame `json:"frames"`
}

func (p *ParticipantFrames) UnmarshalJSON(b []byte) error {
	var obj map[int]MatchParticipantFrame
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}
	var vals []MatchParticipantFrame

	for _, v := range obj {
		vals = append(vals, v)
	}
	p.Frames = vals
	return nil
}

type MatchFrame struct {
	Timestamp         types.Milliseconds `json:"timestamp"`
	ParticipantFrames ParticipantFrames  `json:"participantFrames"`
	Events            []MatchEvent       `json:"events"`
}

type MatchParticipantFrame struct {
	TotalGold           int           `json:"totalGold"`
	TeamScore           int           `json:"teamScore"`
	ParticipantID       int           `json:"participantID"`
	Level               int           `json:"level"`
	CurrentGold         int           `json:"currentGold"`
	MinionsKilled       int           `json:"minionsKilled"`
	DominionScore       int           `json:"dominionScore"`
	Position            MatchPosition `json:"position"`
	XP                  int           `json:"xP"`
	JungleMinionsKilled int           `json:"jungleMinionsKilled"`
}

type MatchPosition struct {
	Y int `json:"y"`
	X int `json:"x"`
}

type MatchEvent struct {
	EventType               string             `json:"eventType"`
	TowerType               string             `json:"towerType"`
	TeamID                  int                `json:"teamID"`
	AscendedType            string             `json:"ascendedType"`
	KillerID                int                `json:"killerID"`
	LevelUpType             string             `json:"levelUpType"`
	PointCaptured           string             `json:"pointCaptured"`
	AssistingParticipantIDs []int              `json:"assistingParticipantIDs"`
	WardType                string             `json:"wardType"`
	MonsterType             string             `json:"monsterType"`
	Type                    event.Event        `json:"type"`
	SkillSlot               int                `json:"skillSlot"`
	VictimID                int                `json:"victimID"`
	Timestamp               types.Milliseconds `json:"timestamp"`
	AfterID                 int                `json:"afterID"`
	MonsterSubType          string             `json:"monsterSubType"`
	LaneType                lane.Type          `json:"laneType"`
	ItemID                  int                `json:"itemID"`
	ParticipantID           int                `json:"participantID"`
	BuildingType            string             `json:"buildingType"`
	CreatorID               int                `json:"creatorID"`
	Position                MatchPosition      `json:"position"`
	BeforeID                int                `json:"beforeID"`
}

func (c *client) GetMatchTimeline(ctx context.Context, r region.Region, matchID int64) (*MatchTimeline, error) {
	var res MatchTimeline
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/match/v4/timelines/by-match", fmt.Sprintf("/%d", matchID), nil, &res)
	return &res, err
}

package levels

import (
	"encoding/json"
	"fmt"
	h "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"strconv"
)

var Home, _ = h.Dir()

type CharLevel struct {
	Level               int `json:"level"`
	AbilityScoreBonuses int `json:"ability_score_bonuses,omitempty"`
	Index               int `json:"index"`

	ProfBonus      int           `json:"prof_bonus,omitempty"`
	FeatureChoices []interface{} `json:"feature_choices"`
	Features       []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"features"`
	ClassSpecific struct {
		BardicInspirationDie   int `json:"bardic_inspiration_die"`
		SongOfRestDie          int `json:"song_of_rest_die"`
		MagicalSecretsMax5     int `json:"magical_secrets_max_5"`
		MagicalSecretsMax7     int `json:"magical_secrets_max_7"`
		MagicalSecretsMax9     int `json:"magical_secrets_max_9"`
		ChannelDivinityCharges int `json:"channel_divinity_charges"`
		DestroyUndeadCr        int `json:"destroy_undead_cr"`
		ActionSurges           int `json:"action_surges"`
		IndomitableUses        int `json:"indomitable_uses"`
		ExtraAttacks           int `json:"extra_attacks"`

		ArcaneRecoveryLevels int `json:"arcane_recovery_levels"`

		InvocationsKnown    int `json:"invocations_known"`
		MysticArcanumLevel6 int `json:"mystic_arcanum_level_6"`
		MysticArcanumLevel7 int `json:"mystic_arcanum_level_7"`
		MysticArcanumLevel8 int `json:"mystic_arcanum_level_8"`
		MysticArcanumLevel9 int `json:"mystic_arcanum_level_9"`

		MartialArts struct {
			DiceCount int `json:"dice_count"`
			DiceValue int `json:"dice_value"`
		} `json:"martial_arts"`
		KiPoints          int `json:"ki_points"`
		UnarmoredMovement int `json:"unarmored_movement"`

		SneakAttack struct {
			DiceCount int `json:"dice_count"`
			DiceValue int `json:"dice_value"`
		} `json:"sneak_attack"`

		SorceryPoints      int           `json:"sorcery_points"`
		MetamagicKnown     int           `json:"metamagic_known"`
		CreatingSpellSlots []interface{} `json:"creating_spell_slots"`

		FavoredEnemies int `json:"favored_enemies"`
		FavoredTerrain int `json:"favored_terrain"`

		AuraRange int `json:"aura_range"`

		WildShapeMaxCr int  `json:"wild_shape_max_cr"`
		WildShapeSwim  bool `json:"wild_shape_swim"`
		WildShapeFly   bool `json:"wild_shape_fly"`

		RageCount          int `json:"rage_count"`
		RageDamageBonus    int `json:"rage_damage_bonus"`
		BrutalCriticalDice int `json:"brutal_critical_dice"`
	}
	ClassSpecificBarbarian struct {
		RageCount          int `json:"rage_count"`
		RageDamageBonus    int `json:"rage_damage_bonus"`
		BrutalCriticalDice int `json:"brutal_critical_dice"`
	} `json:"class_specific,omitempty"`
	Class struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"class"`
	URL string `json:"url"`
	//Spellcasting struct {
	//} `json:"spellcasting,omitempty"`
	Spellcasting struct {
		CantripsKnown    int `json:"cantrips_known"`
		SpellsKnown      int `json:"spells_known"`
		SpellSlotsLevel1 int `json:"spell_slots_level_1"`
		SpellSlotsLevel2 int `json:"spell_slots_level_2"`
		SpellSlotsLevel3 int `json:"spell_slots_level_3"`
		SpellSlotsLevel4 int `json:"spell_slots_level_4"`
		SpellSlotsLevel5 int `json:"spell_slots_level_5"`
		SpellSlotsLevel6 int `json:"spell_slots_level_6"`
		SpellSlotsLevel7 int `json:"spell_slots_level_7"`
		SpellSlotsLevel8 int `json:"spell_slots_level_8"`
		SpellSlotsLevel9 int `json:"spell_slots_level_9"`
	} `json:"spellcasting,omitempty"`
	ClassSpecificBard struct {
		BardicInspirationDie int `json:"bardic_inspiration_die"`
		SongOfRestDie        int `json:"song_of_rest_die"`
		MagicalSecretsMax5   int `json:"magical_secrets_max_5"`
		MagicalSecretsMax7   int `json:"magical_secrets_max_7"`
		MagicalSecretsMax9   int `json:"magical_secrets_max_9"`
	} `json:"class_specific,omitempty"`
	ClassSpecificCleric struct {
		ChannelDivinityCharges int `json:"channel_divinity_charges"`
		DestroyUndeadCr        int `json:"destroy_undead_cr"`
	} `json:"class_specific,omitempty"`
	ClassSpecificDruid struct {
		WildShapeMaxCr int  `json:"wild_shape_max_cr"`
		WildShapeSwim  bool `json:"wild_shape_swim"`
		WildShapeFly   bool `json:"wild_shape_fly"`
	} `json:"class_specific,omitempty"`
	ClassSpecificFighter struct {
		ActionSurges    int `json:"action_surges"`
		IndomitableUses int `json:"indomitable_uses"`
		ExtraAttacks    int `json:"extra_attacks"`
	} `json:"class_specific,omitempty"`
	ClassSpecificMonk struct {
		MartialArts struct {
			DiceCount int `json:"dice_count"`
			DiceValue int `json:"dice_value"`
		} `json:"martial_arts"`
		KiPoints          int `json:"ki_points"`
		UnarmoredMovement int `json:"unarmored_movement"`
	} `json:"class_specific,omitempty"`
	ClassSpecificPaladin struct {
		AuraRange int `json:"aura_range"`
	} `json:"class_specific,omitempty"`
	ClassSpecificRanger struct {
		FavoredEnemies int `json:"favored_enemies"`
		FavoredTerrain int `json:"favored_terrain"`
	} `json:"class_specific,omitempty"`
	ClassSpecificRogue struct {
		SneakAttack struct {
			DiceCount int `json:"dice_count"`
			DiceValue int `json:"dice_value"`
		} `json:"sneak_attack"`
	} `json:"class_specific,omitempty"`
	ClassSpecificSorcerer struct {
		SorceryPoints      int           `json:"sorcery_points"`
		MetamagicKnown     int           `json:"metamagic_known"`
		CreatingSpellSlots []interface{} `json:"creating_spell_slots"`
	} `json:"class_specific,omitempty"`
	ClassSpecificWarlock struct {
		InvocationsKnown    int `json:"invocations_known"`
		MysticArcanumLevel6 int `json:"mystic_arcanum_level_6"`
		MysticArcanumLevel7 int `json:"mystic_arcanum_level_7"`
		MysticArcanumLevel8 int `json:"mystic_arcanum_level_8"`
		MysticArcanumLevel9 int `json:"mystic_arcanum_level_9"`
	} `json:"class_specific,omitempty"`
	ClassSpecificWizard struct {
		ArcaneRecoveryLevels int `json:"arcane_recovery_levels"`
	} `json:"class_specific,omitempty"`
	Subclass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"subclass,omitempty"`
	SubclassSpecific struct {
		AdditionalMagicalSecretsMaxLvl int `json:"additional_magical_secrets_max_lvl"`
	} `json:"subclass_specific,omitempty"`
}
type Levels struct {
	Levels []CharLevel
}

func (rc *Levels) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/levels.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &rc.Levels); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}

func (rc *Levels) ById(id string) {

	n, _ := strconv.Atoi(id)
	fmt.Println(rc.Levels[n]) //.Strength)
	return
}

func (rc *Levels) List() {
	for i := range rc.Levels {
		fmt.Print((rc.Levels[i].Index - 1))
		fmt.Println(" ", rc.Levels[i].Class.Name, rc.Levels[i].Level)
	}
	//fmt.Println(cl.Levels)
	return

}

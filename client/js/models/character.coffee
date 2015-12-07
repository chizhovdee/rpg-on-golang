
# Аналоги серверной реализции
# energyPoints()
# healthPoints()
# restorable()
# restoreSeconds()
# restoresSinceLastUpdate()
# restoreBonus()


Spine = require("spine")
CharacterRestorableAttributes = require("./character_restorable_attributes.coffee")

class Character extends Spine.Model
  @configure "Character", "level", "experience", "ep", "energy", "hp", "health",
    "basic_money", "vip_money", "full_refill_duration", "hp_restore_duration",
    "ep_restore_duration", "hp_updated_at", "ep_updated_at",

    # for test
    "restorable_hp", "restorable_ep"


  @include CharacterRestorableAttributes

  energyPoints: ->
    @energy

  healthPoints: ->
    @health


module.exports = Character


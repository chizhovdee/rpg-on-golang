utils = require("../utils/utils.coffee")
CharacterRestorableAttributes = require("./character_restorable_attributes.coffee")

class Character extends Spine.Model
  @configure "Character", "level", "experience", "ep", "energy", "hp", "health",
    "basic_money", "vip_money", "full_refill_duration", "hp_restore_duration",
    "ep_restore_duration", "hp_updated_at", "ep_updated_at", "oldAttributes"

    # for test
    "restorable_hp", "restorable_ep"

  @include CharacterRestorableAttributes


  setOldAttributes: (attributes)->
    @oldAttributes = utils.deepClone(attributes, "oldAttributes")

  changes: ->
    changes = {}

    for attribute, value of @.attributes()
      continue if attribute == "oldAttributes"

      if _.isObject(value)
        unless _.isEqual(value, @oldAttributes[attribute])
          changes[attribute] = [@oldAttributes[attribute], value]

      else if value != @oldAttributes[attribute]
        changes[attribute] = [@oldAttributes[attribute], value]

    changes

  energyPoints: ->
    # server duplicate
    @energy

  healthPoints: ->
    # server duplicate
    @health


module.exports = Character


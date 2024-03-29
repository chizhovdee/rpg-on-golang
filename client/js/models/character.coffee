utils = require("../utils/utils.coffee")

class Character extends Spine.Model
  @configure "Character", "level", "experience", "restorable_ep", "energy_points",
    "restorable_hp", "health_points", "basic_money", "vip_money", "hp_restore_in",
    "ep_restore_in", "oldAttributes"

  update: ->
    @.setOldAttributes(@constructor.irecords[@id].attributes())

    super

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

  epPercentage: ->
    @restorable_ep / @energy_points * 100

  hpPercentage: ->
    @restorable_hp / @health_points * 100

module.exports = Character


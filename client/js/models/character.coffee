Spine = require("spine")

class Character extends Spine.Model
  @configure "Character", "level", "experience", "ep", "energy", "hp", "health",
    "basic_money", "vip_money", "full_refill_duration", "hp_restore_duration",
    "ep_restore_duration", "hp_restore_bonus", "ep_restore_bonus", "hp_updated_at",
    "ep_updated_at",

    # for test
    "restorable_hp", "restorable_ep"


  restoreSeconds: (attribute)->
    throw new Error("Requires one argument") unless attribute?

    Math.floor(@["#{attribute}_restore_duration"] * (100 - @["#{attribute}_restore_bonus"]) / 100)


  restoresSinceLastUpdate: (attribute)->
    Math.floor(
      (
        Math.floor((Date.now() / 1000)) - @["#{attribute}_updated_at"]
      ) / @.restoreSeconds(attribute)
    )

  restorable: (attribute)->
    total = (
      switch attribute
        when "hp"
          @health
        when "ep"
          @energy
        else
          throw new Error("No correct or not encluded argument")
    )

    if @["#{attribute}_updated_at"] < Math.floor(Date.now() / 1000) - @full_refill_duration
      total
    else
      calculatedValue = @[attribute] + @.restoresSinceLastUpdate(attribute)

      if calculatedValue >= total
        total
      else if calculatedValue < 0
        0
      else
        calculatedValue

module.exports = Character


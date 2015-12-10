CharacterRestorableAttributes =
  restorable: (attribute)->
    # server duplicate

    total = (
      switch attribute
        when "hp"
          @.healthPoints()
        when "ep"
          @.energyPoints()
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

  restoreSeconds: (attribute)->
    # server duplicate

    throw new Error("Requires one argument") unless attribute?

    Math.floor(@["#{attribute}_restore_duration"] * (100 - @.restoreBonus(attribute)) / 100)

  restoresSinceLastUpdate: (attribute)->
    # server duplicate

    Math.floor(
      (
        Math.floor((Date.now() / 1000)) - @["#{attribute}_updated_at"]
      ) / @.restoreSeconds(attribute)
    )

  restoreBonus: (attribute)->
    # server duplicate

    switch attribute
      when "hp"
        0
      when "ep"
        0
      else
        throw new Error("No correct or not encluded argument")

  secondsSinceLastUpdate: (attribute)->
    restores = (
      Math.floor((Date.now() / 1000)) - @["#{attribute}_updated_at"]
    ) / @.restoreSeconds(attribute)

    decimal = restores - Math.floor(restores)

    if decimal > 0
      @.restoreSeconds(attribute) * (1 - decimal)
    else
      @.restoreSeconds(attribute)

  epPercentage: ->
    @.restorable("ep") / @.energyPoints() * 100

  hpPercentage: ->
    @.restorable("hp") / @.healthPoints() * 100

module.exports = CharacterRestorableAttributes

Layer = require("./base.coffee")
Character = require("../models/character.coffee")
VisualTimer = require("../lib/visual_timer.coffee")

class MainHeaderLayer extends Layer
  elements:
    '.energy': 'energyEl'
    '.health': 'healthEl'

  epTimer: null
  hpTimer: null

  hide: ->
    @epTimer?.stop()
    @epTimer = null

    @hpTimer?.stop()
    @hpTimer = null

    super

  show: ->
    @character = Character.first() # !!! должен быть ининциализован первым

    super

    console.log @character
    console.log "Hp", "serv:", @character.restorable_hp, "client:", @character.restorable("hp")
    console.log "Ep", "serv:", @character.restorable_ep, "client:", @character.restorable("ep")

    console.log "last hp secs", @character.secondsSinceLastUpdate("hp")
    console.log "last ep secs", @character.secondsSinceLastUpdate("ep")

    @.render()

    @.setupTimers()

  render: ->
    @html(@.renderTemplate("main_header"))

  bindEventListeners: ->
    super

    @character.bind("beforeUpdate", @.onCharacterUpdate)
    @character.bind("beforeUpdate", @.onCharacterUpdate1)

  unbindEventListeners: ->
    super

    @character.unbind("beforeUpdate", @.onCharacterUpdate)
    @character.unbind("beforeUpdate", @.onCharacterUpdate1)

  setupTimers: ->
    @epTimer = new VisualTimer(@energyEl.find(".timer"), @.updateEp)
    @hpTimer = new VisualTimer(@healthEl.find(".timer"), @.updateHp)

    @epTimer.start(@character.secondsSinceLastUpdate("ep"))
    @hpTimer.start(@character.secondsSinceLastUpdate("hp"))

  updateEp: =>
    console.log "Updated ep"

    @energyEl.find(".value").text("#{ @character.restorable("ep") } / #{ @character.energyPoints() }")
    @energyEl.find(".progress").css(width: "#{ @character.epPercentage() }%")

    if @character.restorable("ep") < @character.energyPoints()
      @epTimer.start(@character.secondsSinceLastUpdate("ep"))

    console.log "last hp secs", @character.secondsSinceLastUpdate("ep")

  updateHp: =>
    console.log "Updated hp"

    @healthEl.find(".value").text("#{ @character.restorable("hp") } / #{ @character.healthPoints() }")
    @healthEl.find(".progress").css(width: "#{ @character.hpPercentage() }%")

    if @character.restorable("hp") < @character.healthPoints()
      @hpTimer.start(@character.secondsSinceLastUpdate("hp"))

    console.log "last hp secs", @character.secondsSinceLastUpdate("hp")

  onCharacterUpdate: (args...)=>
    console.log "Character Update"
    console.log args[0].clone(args[0].attributes())
    console.log @character.attributes().hp
    console.log args[0].attributes().hp
    console.log args[1]

  onCharacterUpdate1: =>
    console.log "onCharacterUpdate1"

module.exports = MainHeaderLayer
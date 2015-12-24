Layer = require("../lib/layer.coffee")
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

    @.render()

    @.setupTimers()

  render: ->
    @html(@.renderTemplate("main_header"))

  bindEventListeners: ->
    super

    @character.bind("update", (args...)=> @.onCharacterUpdate(args...))

  unbindEventListeners: ->
    super

    @character.unbind("update", (args...)=> @.onCharacterUpdate(args...))

  setupTimers: ->
    @epTimer = new VisualTimer(@energyEl.find(".timer"), => @.updateEp())
    @hpTimer = new VisualTimer(@healthEl.find(".timer"), => @.updateHp())

    @.startEpTimer()
    @.startHpTimer()

  updateEp: ->
    console.log "Updated ep"

    @energyEl.find(".value").text("#{ @character.restorable("ep") } / #{ @character.energyPoints() }")
    @energyEl.find(".progress").css(width: "#{ @character.epPercentage() }%")

    @.startEpTimer()

  startEpTimer: ->
    if @character.restorable("ep") < @character.energyPoints()
      @epTimer.start(@character.secondsSinceLastUpdate("ep"))

    console.log "last hp secs", @character.secondsSinceLastUpdate("ep")

  updateHp: ->
    console.log "Updated hp"

    @healthEl.find(".value").text("#{ @character.restorable("hp") } / #{ @character.healthPoints() }")
    @healthEl.find(".progress").css(width: "#{ @character.hpPercentage() }%")

    @.startHpTimer()

  startHpTimer: ->
    if @character.restorable("hp") < @character.healthPoints()
      @hpTimer.start(@character.secondsSinceLastUpdate("hp"))

    console.log "last hp secs", @character.secondsSinceLastUpdate("hp")

  onCharacterUpdate: (character)->
    console.log "Character Update"
    console.log changes = character.changes()

    @.updateHp() if changes.hp?
    @.updateEp() if changes.ep?



module.exports = MainHeaderLayer
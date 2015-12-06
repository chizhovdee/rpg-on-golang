Layer = require("./base.coffee")
Character = require("../../models/character.coffee")

class MainHeaderLayer extends Layer
  show: ->
    super

    console.log @character = Character.first()
    console.log "Hp", "serv:", @character.restorable_hp, "client:", @character.restorable("hp")
    console.log "Ep", "serv:", @character.restorable_ep, "client:", @character.restorable("ep")

    @.render()

  render: ->
    @html(@.renderTemplate("main_header"))

  bindEventListeners:->


  unbindEventListeners:->

module.exports = MainHeaderLayer
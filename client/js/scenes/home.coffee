Scene = require("./base.coffee")
MainHeaderLayer = require("../layers/main_header.coffee")
Character = require("../models/character.coffee")

class HomeScene extends Scene
  className: "home scene"

  hide: ->
    MainHeaderLayer.hide()

    super

  show: ->
    super

    @.render()

    MainHeaderLayer.show(el: @el.find("#main_header"))

    @character = Character.first()

  render: ->
    @html(@.renderTemplate("home/index"))

  bindEventListeners: ->
    super

    @el.on("click", ".menu.missions", @.onMenuClick)


  unbindEventListeners: ->
    super

    @el.off("click", ".menu.missions", @.onMenuClick)

  onMenuClick: (e)=>
    console.log "Click missions"

module.exports = HomeScene

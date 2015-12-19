Scene = require("./base.coffee")
sceneManager = require("../lib/scene_manager.coffee")
MainHeaderLayer = require("../layers/main_header.coffee")
Character = require("../models/character.coffee")
HomeScene = require("./home.coffee")

class MissionsScene extends Scene
  className: "missions scene"

  hide: ->
    MainHeaderLayer.hide()

    super

  show: ->
    super

    @.render()

    MainHeaderLayer.show(el: @el.find("#main_header"))

  render: ->
    @html(@.renderTemplate("missions/index"))


  bindEventListeners: ->
    super

    @el.on("click", "button.back", => @.onBackButtonClick())

  unbindEventListeners: ->
    super

    @el.off("click", "button.back", => @.onBackButtonClick())

  onBackButtonClick: ->
    sceneManager.run("home")

module.exports = MissionsScene

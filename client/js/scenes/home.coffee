Scene = require("./base.coffee")
MainHeaderLayer = require("../layers/main_header.coffee")
Character = require("../models/character.coffee")
sceneManager = require("../lib/scene_manager.coffee")

class HomeScene extends Scene
  className: "home scene"

  hide: ->
    MainHeaderLayer.hide()

    super

  show: ->
    super

    @.render()

    MainHeaderLayer.show(el: @el.find("#main_header"))

  render: ->
    @html(@.renderTemplate("home/index"))

  bindEventListeners: ->
    super

    @el.on("click", ".menu.missions", (e)=> @.onMenuClick(e))

  unbindEventListeners: ->
    super

    @el.off("click", ".menu.missions", (e)=> @.onMenuClick(e))

  onMenuClick: (e)->
    sceneManager.run("missions")

module.exports = HomeScene

Scene = require("../lib/scene.coffee")
Character = require("../models/character.coffee")
sceneManager = require("../lib/scene_manager.coffee")

class HomeScene extends Scene
  className: "home scene"

  hide: ->
    super

  show: ->
    super

    @.render()

  render: ->
    @html(@.renderTemplate("home/index"))

  bindEventListeners: ->
    super

  unbindEventListeners: ->
    super

module.exports = HomeScene

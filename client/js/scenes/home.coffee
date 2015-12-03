Scene = require("./base.coffee")
Character = require("../models/character.coffee")

class HomeScene extends Scene
  className: "home scene"

  show: ->
    super

    console.log @character = Character.first()

    @.render()

  render: ->
    @html(@.renderTemplate("home/index"))

module.exports = HomeScene

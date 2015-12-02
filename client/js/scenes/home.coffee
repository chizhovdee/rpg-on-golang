Scene = require("./base.coffee")

class HomeScene extends Scene
  className: "home scene"

  constructor: ->
    super

    console.log "HOME SCENE"

  show: ->
    super

    @first_data = 1111111111

    @.render()

  render: ->
    @el.html(@.renderTemplate("home/index"))

module.exports = HomeScene

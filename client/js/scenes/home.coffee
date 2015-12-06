Scene = require("./base.coffee")
MainHeaderLayer = require("./layers/main_header.coffee")

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

module.exports = HomeScene

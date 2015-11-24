Spine = require("Spine")

class App extends Spine.Controller
  className: "app"

  constructor: ->
    super

    console.log 11111111

    @el.html(JST["test"]())

    @el.appendTo("body")

$(->
  new App()

)
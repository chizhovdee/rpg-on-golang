HomeScene = require("./scenes/home.coffee")

class Application
  constructor: ->
    console.log "Initialize application"

    new HomeScene()

    console.log "1111"

module.exports = Application

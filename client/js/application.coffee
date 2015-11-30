HomeScene = require("./scenes/home.coffee")
transport = require("./lib/transport.coffee")

class Application
  constructor: ->
    console.log "Initialize application"

    transport.one("character_game_data_loaded", (response)-> console.log "EHUI", response)

    transport.send("loadCharacterGameData")

    HomeScene.show()

module.exports = Application

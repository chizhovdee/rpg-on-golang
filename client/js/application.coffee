HomeScene = require("./scenes/home.coffee")
transport = require("./lib/transport.coffee")

class Application
  constructor: ->
    console.log "Initialize application"

    @.bindEventListeners()

    transport.send("loadCharacterGameData")

  bindEventListeners: ->
    transport.one("character_game_data_loaded", @.onCharacterGameDataLoaded)

  onCharacterGameDataLoaded: (response)->
    console.log "onCharacterGameDataLoaded"
    console.log response

    HomeScene.show()

module.exports = Application

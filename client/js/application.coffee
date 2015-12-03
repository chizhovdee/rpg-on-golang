HomeScene = require("./scenes/home.coffee")
transport = require("./lib/transport.coffee")

Character = require("./models/character.coffee")

class Application
  constructor: ->
    console.log "Initialize application"

    @.bindEventListeners()

    transport.send("loadCharacterGameData")

  bindEventListeners: ->
    transport.one("character_game_data_loaded", @.onCharacterGameDataLoaded)

  onCharacterGameDataLoaded: (response)->
    Character.create(response.character)

    HomeScene.show()

module.exports = Application

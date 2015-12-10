HomeScene = require("./scenes/home.coffee")
transport = require("./lib/transport.coffee")
Character = require("./models/character.coffee")

class Application
  constructor: ->
    console.log "Initialize application"

    transport.one("character_game_data_loaded", @.onCharacterGameDataLoaded)

    transport.send("loadCharacterGameData")

  onCharacterGameDataLoaded: (response)->
    character = Character.create(response.character)

    Character.bind("beforeUpdate", (record)-> record.setOldAttributes(character.attributes()))

    HomeScene.show()

module.exports = Application

transport = require("./lib/transport.coffee")
Character = require("./models/character.coffee")
gameData = require("./game_data.coffee")
sceneManager = require("./lib/scene_manager.coffee")
scenes = require("./scenes/scenes.coffee")

class Application
  constructor: ->
    console.log "Initialize application"

    transport.one("character_game_data_loaded", @.onCharacterGameDataLoaded)

    transport.send("loadCharacterGameData")

  onCharacterGameDataLoaded: (response)->
    character = Character.create(response.character)

    Character.bind("beforeUpdate", (record)-> record.setOldAttributes(character.attributes()))

    gameData.define()

    sceneManager.setup(scenes)

    sceneManager.run("home")

module.exports = Application

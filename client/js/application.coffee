transport = require("./lib/transport.coffee")
Character = require("./models/character.coffee")
gameData = require("./game_data.coffee")
sceneManager = require("./lib/scene_manager.coffee")
scenes = require("./scenes/scenes.coffee")
preloader = require("./preloader.coffee")

class Application
  constructor: ->
    console.log "Initialize application"

    transport.one("character_game_data_loaded", @.onCharacterGameDataLoaded)

    $(document).on("click", "button.back", -> sceneManager.run("home"))

    preloader.on("complete", @.onManifestLoadComplete, this)

    preloader.on("progress", @.onManifestLoadProgress, this)

    preloader.loadManifest([
      {id: "locale", src: "locales/ru.json"}
    ])

  onManifestLoadProgress: (e)->
    console.log "Total:", e.total, ", loaded:", e.loaded

  onManifestLoadComplete: ->
    console.log "onManifestLoadComplete"

    console.log preloader.getResult("locale")

    transport.send("loadCharacterGameData")

  onCharacterGameDataLoaded: (response)->
    character = Character.create(response.character)

    Character.bind("beforeUpdate", (record)-> record.setOldAttributes(character.attributes()))

    gameData.define()

    sceneManager.setup(scenes)

    # язык должен грузиться из сервера в самом начале
    lng = "ru"
    locales = {}
    locales[lng] = { translation: preloader.getResult("locale") }

    i18next.init(
      lng: "ru"
      resources: locales
      (err, t)->
        console.log i18next.t("hello", name: "Dmitry")
        sceneManager.run("home")
    )



module.exports = Application

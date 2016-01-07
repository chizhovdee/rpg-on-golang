transport = require("./lib/transport.coffee")
Character = require("./models/character.coffee")
gameData = require("./game_data.coffee")
sceneManager = require("./lib/scene_manager.coffee")
scenes = require("./scenes/scenes.coffee")
preloader = require("./preloader.coffee")
HeaderLayer = require("./layers/header.coffee")

# сначала грузиться манифест с помощью прелоадера
# затем загружается персонаж
# затем запускается главная сцена

class Application
  character: null

  constructor: ->
    @.setupEventListeners()

    gameData.define()

    sceneManager.setup(scenes)

    preloader.loadManifest([
      {id: "locale", src: "locales/#{ window.lng }.json"}
    ])

  # все общие события для игры
  setupEventListeners: ->
    # события прелоадера
    preloader.on("complete", @.onManifestLoadComplete, this)
    preloader.on("progress", @.onManifestLoadProgress, this)

    # события транспорта
    transport.one("character_game_data_loaded", (response)=> @.onCharacterGameDataLoaded(response))
    transport.bind("character_status_loaded", (response)=> @.onCharacterStatusLoaded(response))

    # события DOM

  onManifestLoadProgress: (e)->
    console.log "Total:", e.total, ", loaded:", e.loaded

  onManifestLoadComplete: ->
    console.log "onManifestLoadComplete"

    @.setTranslations()

    transport.send("loadCharacterGameData")

  onCharacterGameDataLoaded: (response)->
    Character.create(response.character)

    HeaderLayer.show(el: $("#application .header"))

    sceneManager.run("home")

  onCharacterStatusLoaded: (response)->
    console.log "onCharacterStatusLoaded"

    @character ?= Character.first()

    @character.updateAttributes(response.character)

  setTranslations: ->
    I18n.defaultLocale = window.lng
    I18n.locale = window.lng
    I18n.translations ?= {}
    I18n.translations[window.lng] = preloader.getResult("locale")

module.exports = Application

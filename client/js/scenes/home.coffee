Scene = require("../lib/scene.coffee")
Character = require("../models/character.coffee")
sceneManager = require("../lib/scene_manager.coffee")

class HomeScene extends Scene
  className: "home scene"

  hide: ->
    super

  show: ->
    super

    @.render()

  render: ->
    @html(@.renderTemplate("home/index"))

  bindEventListeners: ->
    super

    @el.on("click", ".menu.quests", (e)=> @.onMenuClick(e))
    @el.on("click", ".menu.shop", (e)=> @.onShopClick(e))

  unbindEventListeners: ->
    super

    @el.off("click", ".menu.quests", (e)=> @.onMenuClick(e))
    @el.off("click", ".menu.shop", (e)=> @.onShopClick(e))

  onMenuClick: (e)->
    sceneManager.run("quests")

  onShopClick: (e)->
    console.log "Shop click"

module.exports = HomeScene

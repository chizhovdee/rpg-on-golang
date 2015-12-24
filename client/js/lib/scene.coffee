BaseController = require("./base_controller.coffee")

class Scene extends BaseController
  @show: ->
    super

    @scene ?= new @()
    @scene.show()

  @hide: ->
    super

    @scene?.hide()
    @scene = null

  show: ->
    super

    console.log "Show Scene"

    $("#application").append(@el)

  hide: ->
    console.log "Hide Scene"

    super

module.exports = Scene

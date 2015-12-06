Spine = require("spine")
RenderUtils = require("../utils/render.coffee")
TimeUtils = require("../utils/time.coffee")

$ = Spine.$

class Scene extends Spine.Controller
  @include RenderUtils
  @include TimeUtils

  @show: ->
    @scene ?= new @()
    @scene.show()

  @hide: ->
    # утилизация
    @scene?.hide()
    @scene = null

  show: ->
    console.log "Show Scene"

    $("#application").append(@el)

  hide: ->
    console.log "Hide Scene"

    @el.remove()


module.exports = Scene

Spine = require("spine")
RenderUtils = require("../../utils/render.coffee")
TimeUtils = require("../../utils/time.coffee")
DesignUtils = require("../../utils/design.coffee")

$ = Spine.$

class Layer extends Spine.Controller
  @include RenderUtils
  @include TimeUtils
  @include DesignUtils

  @show: (data = {})->
    throw new Error("Requires 'el' data options for layer rendering") unless data.el?

    @layer ?= new @(data)
    @layer.show()

  @hide: ->
# утилизация
    @layer?.hide()
    @layer = null

  show: ->
    @.bindEventListeners()

  hide: ->
    console.log "Hide Layer"

    @.unbindEventListeners()

    @el.remove()

  bindEventListeners: ->

  unbindEventListeners: ->


module.exports = Layer
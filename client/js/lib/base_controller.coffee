RenderUtils = require("../utils/render.coffee")
TimeUtils = require("../utils/time.coffee")
DesignUtils = require("../utils/design.coffee")

class BaseController extends Spine.Controller
  @include RenderUtils
  @include TimeUtils
  @include DesignUtils

  @show: ->

  @hide: ->

  show: ->
    @.bindEventListeners()

  hide: ->
    @.unbindEventListeners()

    @el.remove()

  bindEventListeners: ->

  unbindEventListeners: ->

module.exports = BaseController
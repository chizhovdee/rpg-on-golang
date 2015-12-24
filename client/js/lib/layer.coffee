BaseController = require("./base_controller.coffee")

class Layer extends BaseController
  @show: (data = {})->
    super

    throw new Error("Requires 'el' data options for layer rendering") unless data.el?

    @layer ?= new @(data)
    @layer.show()

  @hide: ->
    @layer?.hide()
    @layer = null

    super

  show: ->
    super

    console.log "Show Layer"

  hide: ->
    console.log "Hide Layer"

    super


module.exports = Layer
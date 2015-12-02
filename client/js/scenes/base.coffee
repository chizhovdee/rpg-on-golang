Spine = require("spine")
config = require("../config.coffee")
$ = Spine.$

class Scene extends Spine.Controller
  @show: ->
    @scene ?= new @()
    @scene.show()

  @hide: ->
    # утилизация
    @scene.hide()
    @scene = null

  show: ->
    console.log "Show Scene"

    $("#application").append(@el)

  hide: ->
    console.log "Hide Scene"

  renderTemplate: (name)->
    JST[name](@)

module.exports = Scene

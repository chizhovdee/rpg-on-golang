MainScene = require('./scenes/main.coffee')

window.onload = ->
  cc.game.onStart = ->
    cc.LoaderScene.preload([], ->
      # манипуляции с экраном

      cc.director.runScene(new MainScene())
    )

  cc.game.run("gameCanvas")
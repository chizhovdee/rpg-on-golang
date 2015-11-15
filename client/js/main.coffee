GameScene = require('./game.coffee')

window.onload = ->
  cc.game.onStart = ->
    cc.LoaderScene.preload([], ->
      # манипуляции с экраном

      cc.director.runScene(new GameScene())
    )

  cc.game.run("gameCanvas")
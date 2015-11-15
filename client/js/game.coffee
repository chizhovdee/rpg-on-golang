GameScene = cc.Scene.extend(
  onEnter: ->
    @._super()

    console.log "Loaded GameScene"

    size = cc.director.getWinSize()

    label = cc.LabelTTF.create("RPG Game", "Arial", 40)
    label.setPosition(size.width / 2, size.height / 2)
    @.addChild(label, 1)
)

module.exports = GameScene
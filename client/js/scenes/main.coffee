MainHeader = require("../layers/main_header.coffee")

MainLayer = cc.Layer.extend(
  init: ->
    @._super()

    size = cc.winSize

    main_header = new MainHeader()
    main_header.init()
    @.addChild(main_header)

    menuItem1 = new cc.MenuItemFont("Shop", @.onMenuClick)
    menuItem1.setUserData("shop")
    menuItem2 = new cc.MenuItemFont("Missions", @.onMenuClick)
    menuItem2.setUserData("Missions")
    menuItem3 = new cc.MenuItemFont("Monsters", @.onMenuClick)
    menuItem3.setUserData("Monsters")

    menuItem1.setPosition(cc.p(0, 0))
    menuItem2.setPosition(cc.p(0, 30))
    menuItem3.setPosition(cc.p(0, 60))

    menu = new cc.Menu(menuItem1, menuItem2, menuItem3)
    menu.setPosition(cc.p(size.width / 2, size.height / 2))
    @.addChild(menu)

    true

  onMenuClick: (e)->
    console.log e.getUserData()
)

MainScene = cc.Scene.extend(
  onEnter: ->
    @._super()

    layer = new MainLayer()
    layer.init()

    @.addChild(layer, 0)
)

module.exports = MainScene
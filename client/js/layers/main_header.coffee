config = require("../config.coffee")

MainHeader = cc.Layer.extend(
  init: ->
    @._super()
    size = cc.winSize

    @.setPosition(cc.p(0, size.height - 10));

    level_text = new ccui.Text("Level: 1", config.font, 14);
    level_text.attr({
      textAlign: cc.TEXT_ALIGNMENT_LEFT
    })
    level_text.setAnchorPoint(cc.p(0, 1))
    level_text.setPosition(cc.p(10, 0))

    basic_money_text = new ccui.Text("Gold: 100", config.font, 14);
    basic_money_text.attr({
      textAlign: cc.TEXT_ALIGNMENT_LEFT
    })
    basic_money_text.setAnchorPoint(cc.p(0, 1))
    basic_money_text.setPosition(cc.p(100, 0))

    vip_money_text = new ccui.Text("Gems: 100", config.font, 14);
    vip_money_text.attr({
      textAlign: cc.TEXT_ALIGNMENT_LEFT
    })
    vip_money_text.setAnchorPoint(cc.p(0, 1))
    vip_money_text.setPosition(cc.p(200, 0))

    health_text = new ccui.Text("Health: 100/100", config.font, 14);
    health_text.attr({
      textAlign: cc.TEXT_ALIGNMENT_LEFT
    })
    health_text.setAnchorPoint(cc.p(0, 1))
    health_text.setPosition(cc.p(300, 0))

    energy_text = new ccui.Text("Energy: 100/100", config.font, 14);
    energy_text.attr({
      textAlign: cc.TEXT_ALIGNMENT_LEFT
    })
    energy_text.setAnchorPoint(cc.p(0, 1))
    energy_text.setPosition(cc.p(450, 0))

    stamina_text = new ccui.Text("Stamina: 100/100", config.font, 14);
    stamina_text.attr({
      textAlign: cc.TEXT_ALIGNMENT_LEFT
    })
    stamina_text.setAnchorPoint(cc.p(0, 1))
    stamina_text.setPosition(cc.p(600, 0))

    @.addChild(level_text)
    @.addChild(basic_money_text)
    @.addChild(vip_money_text)
    @.addChild(health_text)
    @.addChild(energy_text)
    @.addChild(stamina_text)
)

module.exports = MainHeader
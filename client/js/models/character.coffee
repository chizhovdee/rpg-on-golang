Spine = require("spine")

class Character extends Spine.Model
  @configure "Character", "level", "experience", "ep", "energy", "hp", "health",
    "basic_money", "vip_money"


module.exports = Character


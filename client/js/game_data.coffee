# Этот файл создан автоматически с помощью gofer game_data:export


mission = require("./game_data/mission.coffee")

mission_group = require("./game_data/mission_group.coffee")


module.exports =
  define: ->
    
    mission.define({"keys":["id","key","mission_group_key"],"values":[[224668150,"mission_1","mission_group_1"],[2490202188,"mission_2","mission_group_1"]]})
    
    mission_group.define({"keys":["id","key"],"values":[[2065960017,"mission_group_1"],[3794623979,"mission_group_2"]]})
    

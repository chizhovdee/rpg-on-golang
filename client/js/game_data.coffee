# Этот файл создан автоматически с помощью gofer game_data:export


quest = require("./game_data/quest.coffee")

quest_group = require("./game_data/quest_group.coffee")


module.exports =
  define: ->
    
    quest.define({"keys":["id","key","quest_group_key"],"values":[[2613178097,"quest_1","quest_group_1"],[46710603,"quest_2","quest_group_1"]]})
    
    quest_group.define({"keys":["id","key"],"values":[[3191809367,"quest_group_1"],[657880301,"quest_group_2"],[1345406075,"quest_group_3"],[3461730776,"quest_group_4"],[3109216590,"quest_group_5"],[542879988,"quest_group_6"],[1465679970,"quest_group_7"]]})
    

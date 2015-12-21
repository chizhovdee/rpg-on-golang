Scene = require("./base.coffee")
sceneManager = require("../lib/scene_manager.coffee")
MainHeaderLayer = require("../layers/main_header.coffee")
Character = require("../models/character.coffee")
MissionGroup = require("../game_data/mission_group.coffee")
Mission = require("../game_data/mission.coffee")
Pagination = require("../lib/pagination.coffee")

class MissionsScene extends Scene
  className: "missions scene"

  MISSION_GROUPS_PER_PAGE = 4
  MISSIONS_PER_PAGE = 3

  hide: ->
    MainHeaderLayer.hide()

    super

  show: ->
    super

    @missionGroupsPagination = new Pagination(MISSION_GROUPS_PER_PAGE)
    #@missionsPagination = new Pagination(MISSIONS_PER_PAGE)

    @missionGroups = MissionGroup.all()

    @paginatedMissionGroups = @missionGroupsPagination.paginate(@missionGroups, initialize: true)

    @currentMissionGroupKey = @paginatedMissionGroups[0].key

    @missions = Mission.findAllByAttribute("mission_group_key", @currentMissionGroupKey)

    @.render()

    MainHeaderLayer.show(el: @el.find("#main_header"))

  render: ->
    @html(@.renderTemplate("missions/index"))

  renderTabs: ->
    @el.find('.tabs').html(@.renderTemplate("missions/tabs"))

  renderMissionList: ->
    @el.find('.mission_list').html(@.renderTemplate("missions/mission_list"))

  bindEventListeners: ->
    super

    @el.on('click', '.tabs .paginate:not(.disabled)', (e)=> @.onTabsPaginateButtonClick(e))
    @el.on("click", ".tab:not(.current)", (e)=> @.onTabClick(e))

  unbindEventListeners: ->
    super

    @el.off('click', '.tabs .paginate:not(.disabled)', (e)=> @.onTabsPaginateButtonClick(e))
    @el.off("click", ".tab:not(.current)", (e)=> @.onTabClick(e))

  onTabsPaginateButtonClick: (e)->
    @paginatedMissionGroups = @missionGroupsPagination.paginate(@missionGroups,
      back: $(e.currentTarget).data('type') == 'back'
    )

    @.renderTabs()

  onTabClick: (e)->
    @el.find(".tab").removeClass("current")

    tabEl = $(e.currentTarget)
    tabEl.addClass("current")

    @currentMissionGroupKey = tabEl.data('tab')

    @missions = Mission.findAllByAttribute("mission_group_key", @currentMissionGroupKey)

    @.renderMissionList()



module.exports = MissionsScene

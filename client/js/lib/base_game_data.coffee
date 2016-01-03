class BaseGameData extends Spine.Model
  @define: (data)->
    @refresh(
      for values in data.values
        obj = {}

        for value, index in values
          obj[data.keys[index]] = value

        obj
    )

module.exports = BaseGameData
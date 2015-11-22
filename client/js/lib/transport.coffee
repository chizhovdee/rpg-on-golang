
transport =
  bind: (eventName, callback)->


  dispatch: (eventName, data)->


  send: (name, data)->
    @[name](data)

  ajax: (type, url, data)->
    $.ajax(
      type: type
      url: url
      data: data
      dataType: "json"
      success: (data)->
        console.log data
        # TO DO
        # dispatch event

      error: (xhr, type)->
        console.error("Ajax error!", xhr, type)
    )

  get: (url, data)->
    @.ajax("GET", url, data)

  post: (url, data)->
    @.ajax("POST", url, data)

  loadShop: ->
    @.get("/shop")


module.exports = transport
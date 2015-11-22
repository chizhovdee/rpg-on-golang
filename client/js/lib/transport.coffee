Zepto = require("Zepto")

transport =
  bind: (eventName, callback)->
    cc.eventManager.addCustomListener(eventName, (data)-> callback(data.getUserData()))

  dispatch: (eventName, data)->
    cc.eventManager.dispatchCustomEvent(eventName, data)

  send: (name, data)->
    @[name](data)

  ajax: (type, url, data)->
    Zepto.ajax(
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
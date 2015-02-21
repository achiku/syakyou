Promise = require 'Promise'

getResources = () ->
  return new Promise (resolve, reject) ->
    setTimeout () ->
      list = []
      for i in [1..10]
        list.push({num: i})
      resolve(list)
    ,Math.floor(Math.random() * 1000)

getAdditionalResource = (item) ->
  return new Promise (resolve, reject) ->
    setTimeout () ->
      if item.num % 2 == 0
        item.isEven = true
      else
        item.isEven = false
      resolve(item)
    ,Math.floor(Math.random() * 1000)


getResources()
  .then (values) ->
    console.log 'getResources done'
    return values
  .then (values) ->
    list = (getAdditionalResource(val) for val in values)
    Promise.all(list)
      .then (values) ->
        for val in values
          console.log val

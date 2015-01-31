Promise = require 'Promise'


promiseMultiply = (x, y, callback) ->
  promise = new Promise (resolve, reject) ->
    setTimeout () ->
      callback(x * y)
    ,Math.floor(Math.random() * 1000)
  return promise


promiseAdd = (x, y, callback) ->
  promise = new Promise (resolve, reject) ->
    setTimeout () ->
      callback(x * y)
    ,Math.floor(Math.random() * 1000)
  return promise

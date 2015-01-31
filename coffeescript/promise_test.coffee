Promise = require 'Promise'


asyncHello = (i) ->
  return new Promise (resolve, reject) ->
    setTimeout () ->
      resolve("Async Hello World n=#{i}")
    , Math.floor(Math.random() * 1000)

array = []
for i in [1..10]
  array.push(asyncHello(i))

console.log "All sync executions"
Promise.all(array).then(
  (values) ->
    for i in values
      console.log i
)

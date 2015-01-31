Promise = require 'Promise'


promiseMultiply = (x, y) ->
  return new Promise (resolve, reject) ->
    setTimeout () ->
      resolve(x * y)
    ,Math.floor(Math.random() * 1000)


promiseAdd = (x, y) ->
  return new Promise (resolve, reject) ->
    setTimeout () ->
      resolve(x + y)
    ,Math.floor(Math.random() * 1000)


multiplyArray = []
for i in [1..10]
  multiplyArray.push(promiseMultiply(2, i))

addArray = []
for i in [1..10]
  addArray.push(promiseAdd(2, i))

allArray = []
for i in [1..10]
  allArray.push(promiseMultiply(2, i))
  allArray.push(promiseAdd(2, i))


Promise.all(multiplyArray).then(
  (values) ->
    console.log values
)

Promise.all(addArray).then(
  (values) ->
    console.log values
)

Promise.all(allArray).then(
  (values) ->
    console.log values
)

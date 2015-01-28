asyncMultiply = (x, y, callback) ->
  setTimeout () ->
    callback(x * y)
  ,Math.floor(Math.random() * 1000)


for i in [1..10]
  asyncMultiply(i, 3,
    (result) -> console.log "result: #{result}"
  )

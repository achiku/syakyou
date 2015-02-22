uniq = require 'uniq'
utils = require './utils.coffee'
Person = require './person.coffee'

nums = [ 5, 2, 1, 3, 2, 5, 4, 2, 0, 1 ]
console.log uniq(nums)

achiku = new Person('achiku', 29)
console.log utils.pprintPerson(achiku)

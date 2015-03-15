$ = require 'jquery'
uniq = require 'uniq'
utils = require './components/utils.coffee'
Person = require './components/person.coffee'

nums = [ 5, 2, 1, 3, 2, 5, 4, 2, 0, 1 ]
console.log uniq(nums)
console.log uniq(nums)

achiku = new Person('achiku', 29)
console.log utils.pprintPerson(achiku)
console.log $('body').html()

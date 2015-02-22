# string
line = "test line\n"
line += "another line\n"
console.log line


# array operation
li = ['spam', 'egg', 'ham']
for i, idx in li
  console.log i, idx

li.push 'big mac'
for i in li
  console.log i

for i in li.reverse()
  console.log i

li = [
  'spam'
  'egg'
  'ham'
]
console.log li

numbers = [1, 2, 3, 4, 5, 6, 7, 8]
console.log numbers[0..2]
console.log numbers[3..-2]
console.log numbers[-2..]
console.log numbers[..]

numbers.pop()
console.log numbers

numbers.splice(1, 1)
console.log numbers

# objects operation
people = [
  {name: 'achiku', age: 29}
  ,{name: 'moqada', age: 29}
  ,{name: '8maki', age: 29}
]

for p in people
  console.log "#{p.name} is #{p.age} years old."

for key, val of people[0]
  console.log key, val

kids =
  brother:
    name: 'bob'
    age: 11
  sister:
    name: 'alice'
    age: 11

console.log kids


# functions
sum = (x, y) ->
  x + y
console.log(sum 1, 2)

square = (x) ->
  x * x
console.log(square 7)

console.log(square sum(1, 2))


# class
class Person
  constructor: (@name, @age) ->

  hello: () =>
    console.log "Hello! I'm #{@name}."

achiku = new Person 'achiku', 29
achiku.hello()
console.log achiku.name, achiku.age


print("-- as always")
print("Hello, World!")

print("-- stdout with newline")
print("I said,")

print("-- stdout without newline")
io.write("Hello, ")
io.write("World! \n")


print("-- string variable assingment")
local str1 = "spam "
local str2 = "egg "
local str3 = "ham "


print("-- string concat (part of arguments can be int/float type)")
print(str1 .. str2 .. str3)
local str1, str2, str3 = "i ", "love ", "lua"
print(str1 .. str2 .. str3 .. "!!")
print(str1 .. str2 .. str3 .. 5.2)


print("-- dict like table")
user = {name="achiku", age=29}
print(user["name"], user["age"])
print("user.name: " .. user.name, "user.age: " .. user.age)
for k, v in pairs(user) do
    print(k, v)
end


print("-- list like table")
nums = {'a', 'b', 'c', 'd'}
for i, val in ipairs(nums) do
    print(i, val)
end


print("-- list-ish table of dict-ish table")
users = {
    {name="8maki", age=29},
    {name="moqada", age=29},
    {name="ide", age=26},
    {name="achiku", age=29}
}
for i, u in ipairs(users) do
    print(i, u.name, u.age)
end
print(#users .. " members in users table")


print("-- function definition without return val")
function greet(name)
    print("Hello, " .. name)
end
greet(user.name)


print("-- function with return val")
function add(x, y)
    return x + y
end
print("result of add(3, 4): " .. add(3, 4))


print("-- function with typed arguments (sort of)")
function add(x, y)
    if type(x) ~= "number" or type(y) ~= "number" then
        error("x and y have to be numbers")
    end
    return x + y
end

local success, val = pcall(add, 3, 4)
if success then
    print("Sum of x, y: " .. val)
else
    print("Error: " .. val)
end

local success, val = pcall(add, "hey", 4)
if success then
    print("Sum of x, y: " .. val)
else
    print("Error: " .. val)
end


print("-- function with possible runtime error")
function div(x, y)
    if y == 0 then
        error("zero division")
    else
        return x / y
    end
end

-- normal call
local val = div(6, 3)
print("result of div(6, 3): " .. val)

-- with pcall error handling
local success, val = pcall(div, 4, 2)
if success then
    print("result of div(4, 2): " .. val)
else
    print("Error with div(4, 2): " .. val)
end

local success, val = pcall(div, 4, 0)
if success then
    print("result of div(4, 0): " .. val)
else
    print("Error with div(4, 0): " .. val)
end

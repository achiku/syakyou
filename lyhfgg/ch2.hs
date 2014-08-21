removeNonUppercase :: [Char] -> [Char]
removeNonUppercase st = [ c | c <- st, c `elem` ['A'..'Z'] ]

removeNonUppercase' :: [Char] -> [Char]
removeNonUppercase' st = [ c | c <- st, elem c ['A'..'Z'] ]

addThree :: Int -> Int -> Int -> Int
addThree x y z = x + y + z

length' xs = sum [1 | _ <- xs]

triangles = [(a, b, c) | a <- [1..10], b <- [1..10], c <- [1..10]]
righttriangles = [(a, b, c) | a <- [1..10], b <- [1..10], c <- [1..10], a^2 + b^2 == c^2]

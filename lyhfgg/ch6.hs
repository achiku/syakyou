multThree :: (Num a) => a -> a -> a -> a
multThree x y z = x * y * z

compareWithHundred :: (Num a, Ord a) => a -> Ordering
compareWithHundred x = compare 100 x

compareWithHundred' :: (Num a, Ord a) => a -> Ordering  
compareWithHundred' = compare 100  

divideByTen :: (Floating a) => a -> a
divideByTen = (/10)

isUpperAlphanum :: Char -> Bool
isUpperAlphanum = (`elem` ['A'..'Z'])

{-
*Main> :t elem
elem :: Eq a => a -> [a] -> Bool
*Main> 'a' `elem` ['A'..'Z']
False
*Main> 'A' `elem` ['A'..'Z']
True
*Main> elem 'a' ['A'..'Z']
False
*Main> elem 'A' ['A'..'Z']
True
-}

--  They indicate that the first parameter is a function 
--  that takes something and returns that same thing. 
applyTwice :: (a -> a) -> a -> a
applyTwice f x = f (f x)

zipWith' :: (a -> b -> c) -> [a] -> [b] -> [c]
zipWith' _ [] _ = []
zipWith' _ _ [] = []
zipWith' f (x:xs) (y:ys) = f x y :zipWith' f xs ys

flip' :: (a -> b -> c) -> (b -> a -> c)
flip' f = g
    where g x y = f y x

map' :: (a -> b) -> [a] -> [b]
map' _ [] = []
map' f (x:xs) = f x : map' f xs


filter' :: (a -> Bool) -> [a] -> [a]
filter' _ [] = []
filter' p (x:xs)
    | p x        = x : filter p xs
    | otherwise  = filter p xs

largetstDivisible :: (Integral a) => a
largetstDivisible = head (filter p [100000, 99999..])
    where p x = x `mod` 3829 == 0

chain :: (Integral a) => a -> [a]
chain 1 = [1]
chain n
    | even n = n:chain (n `div` 2)
    | odd n  = n:chain (n*3 + 1)

numLongChain :: Int
numLongChain = length (filter isLong (map chain [1..100]))
    where isLong xs = length xs > 15

numLongChain' :: Int
numLongChain' = length (filter (\xs -> length xs > 15) (map chain [1..100]))

addThree :: (Num a) => a -> a -> a -> a
addThree x y z = x + y + z

addThree' :: (Num a) => a -> a -> a -> a
addThree' = \x -> \y -> \z -> x + y + z

flip'' :: (a -> b -> c) -> b -> a -> c
flip'' f = \x y -> f y x

sum' :: (Num a) => [a] -> a
sum' xs = foldl (\acc x -> acc + x) 0 xs

sum'' :: (Num a) => [a] -> a
sum'' = foldl (+) 0

elem' :: (Eq a) => a -> [a] -> Bool
elem' y ys = foldl (\acc x -> if x == y then True else acc) False ys

mapr' :: (a -> b) -> [a] -> [b]
mapr' f xs = foldr (\x acc -> f x : acc) [] xs


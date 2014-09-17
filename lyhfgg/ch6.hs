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

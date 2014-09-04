lucky :: (Integral a) => a -> String
lucky 7 = "lucky number seven!"
lucky n = "sorry.."

sayMe :: (Integral a) => a -> String
sayMe 1 = "One!"  
sayMe 2 = "Two!"  
sayMe 3 = "Three!"  
sayMe 4 = "Four!"  
sayMe 5 = "Five!"  
sayMe x = "Not between 1 and 5"

factorial :: (Integral a) => a -> a
factorial 0 = 1
factorial n = n * factorial (n-1)

charName :: Char -> String  
charName 'a' = "Albert"  
charName 'b' = "Broseph"  
charName 'c' = "Cecil"  

addVectors :: (Num a) => (a, a) -> (a, a) -> (a, a)
addVectors a b = (fst a + fst b, snd a + snd b)

head' :: [a] -> a
head' [] = error "Can't head empty list"
head' (x:_) = x

tell :: (Show a) => [a] -> String
tell [] = "This list is empty"
tell (x:[]) = "This list has one element: " ++ show x
tell (x:y:[]) = "This list has two elements: " ++ show x ++ " and " ++ show y

bmiTell :: (RealFloat a) => a -> a -> String
bmiTell weight height
    | bmi <= 18.5 = "You're underweight, you emo, you!"
    | bmi <= 25.0 = "You're supposedly normal."
    | bmi <= 30.0 = "You're fat! Lose some weight, fatty!"
    | otherwise = "other"
    where bmi = weight / height ^ 2
          (skinny, normal, fat) = (18.5, 25.0, 30.0)

calcBmi :: (RealFloat a) => [(a, a)] -> [a]
calcBmi xs = [bmi w h | (w, h) <- xs]
    where bmi weight height = weight / height ^ 2

calcBmi' :: (RealFloat a) => [(a, a)] -> [a]
calcBmi' xs = [bmi | (w, h) <- xs, let bmi = w / h ^ 2, bmi >= 25.0]

initial :: String -> String -> String
initial firstname lastname = [f] ++ ". " ++ [l] ++ "."
    where (f:_) = firstname
          (l:_) = lastname

cylinder :: (RealFloat a) => a -> a -> a
cylinder r h =
    let sideArea = 2 * pi * r * h
        topArea = pi * r^2
    in  sideArea + 2 * topArea

describeList :: [a] -> String
describeList xs = "This List is " ++ what xs
    where what [] = "empty."
          what [x] = "a singleton list."
          what xs = "a longer list."

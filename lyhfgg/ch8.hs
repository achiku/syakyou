import qualified Data.Map as Map 


data Point = Point Float Float deriving (Show)

data Shape = 
    Circle Point Float 
    | Rectangle Point Point
    deriving (Show)

surface :: Shape -> Float
surface (Circle _ r) = pi * r ^ 2
surface (Rectangle (Point x1 y1) (Point x2 y2)) = (abs $ x1 - x2) * (abs $ y1 - y2)

nudge :: Shape -> Float -> Float -> Shape
nudge (Circle (Point x y) r) a b = Circle (Point (x + a) (y + b)) r
nudge (Rectangle (Point x1 y1) (Point x2 y2)) a b = 
    Rectangle (Point (x1 + a) (y1 + b)) (Point (x2 + a) (y2 + b))


data Person = Person { firstName :: String
                      ,lastName :: String
                      ,age :: Int
                      ,height :: Float
                      ,phoneNumber :: String
                      ,flavor :: String
                    } deriving (Show)

-- data Car = Car {company :: String, model :: String, year :: Int} deriving (Show)

-- data Maybe a = Nothing | Just a

data Car a b c = Car { company :: a  
                     , model :: b  
                     , year :: c   
                     } deriving (Show) 

data Vector a = Vector a a a deriving(Show)

vplus :: (Num t) => Vector t -> Vector t -> Vector t
(Vector i j k) `vplus` (Vector l m n) = Vector (i+l) (j+m) (k+n)

vectMult :: (Num t) => Vector t -> t -> Vector t
(Vector i j k) `vectMult` m = Vector (i*m) (j*m) (k*m)

scalarMult :: (Num t) => Vector t -> Vector t -> t
(Vector i j k) `scalarMult` (Vector l m n) = i*l + j*m + k*n



type PhoneNumber = String
type UserName = String
type PhoneBook = [(UserName, PhoneNumber)]

phoneBook =
    [("betty","555-2938")
    ,("bonnie","452-2928")
    ,("patsy","493-2928")
    ,("lucille","205-2928")
    ,("wendy","939-8282")
    ,("penny","853-2492")
    ]

inPhoneBook :: UserName -> PhoneNumber -> PhoneBook -> Bool
inPhoneBook name number pbook = (name, number) `elem` pbook

type AssocList k v = [(k, v)]


data LockerState = Taken | Free deriving(Show, Eq)
type Code = String
type LockerMap = Map.Map Int (LockerState, Code)

lockerLookup :: Int -> LockerMap -> Either String Code
lockerLookup lockerNumber map =
    case Map.lookup lockerNumber map of
        Nothing -> Left $ "Locker number " ++ show lockerNumber ++ " doesn't exist."
        Just (state, code) -> if state /= Taken
            then Right code
            else Left $ "Locker " ++ show lockerNumber ++ " is already taken!"

lockers :: LockerMap  
lockers = Map.fromList   
    [(100,(Taken,"ZD39I"))  
    ,(101,(Free,"JAH3I"))  
    ,(103,(Free,"IQSA9"))  
    ,(105,(Free,"QOTSA"))  
    ,(109,(Taken,"893JJ"))  
    ,(110,(Taken,"99292"))  
    ]

-- data List a = Empty | Cons a (List a) deriving (Show, Read, Eq, Ord)  
infixr 5 :-:  
data List a = Empty | a :-: (List a) deriving (Show, Read, Eq, Ord) 

infixr 5 .++
(.++) :: List a -> List a -> List a
Empty .++ ys =  ys
(x :-: xs) .++ ys = x :-: (xs .++ ys)  

data Tree a = EmptyTree | Node a (Tree a) (Tree a) deriving(Show, Read, Eq)

singleton :: a -> Tree a
singleton x = Node x EmptyTree EmptyTree

treeInsert :: (Ord a) => a -> Tree a -> Tree a
treeInsert x EmptyTree = singleton x
treeInsert x (Node a left right)
    | x == a = Node x left right
    | x < a = Node a (treeInsert x left) right
    | x > a = Node a left (treeInsert x right) 

treeElem ::(Ord a) => a -> Tree a -> Bool
treeElem x EmptyTree = False
treeElem x (Node a left right)  
    | x == a = True
    | x < a = treeElem x left
    | x > a = treeElem x right

data TrafficLight = Red | Yellow | Green

instance Eq TrafficLight where
    Red == Red = True
    Green == Green = True
    Yellow == Yellow = True
    _ == _ = False

instance Show TrafficLight where
    show Red = "Red light"
    show Green = "Green light"
    show Yellow = "Yellow light"

class AEq a where
    (==) :: a -> a -> Bool  
    (/=) :: a -> a -> Bool  
    x == y = not (x /= y)  
    x /= y = not (x == y)   

instance (AEq m) => AEq (Maybe m) where
    Just x == Just y = x == y
    Nothing == Nothing = True
    _ == _ = False

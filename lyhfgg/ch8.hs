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

data Maybe a = Nothing | Just a

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
type PhoneName = String
type PhoneBook = [(PhoneName, PhoneNumber)]

phoneBook =
    [("betty","555-2938")
    ,("bonnie","452-2928")
    ,("patsy","493-2928")
    ,("lucille","205-2928")
    ,("wendy","939-8282")
    ,("penny","853-2492")
    ]

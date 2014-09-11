multThree :: (Num a) => a -> a -> a -> a
multThree x y z = x * y * z

compareWithHundred :: (Num a, Ord a) => a -> Ordering
compareWithHundred x = compare 100 x


compareWithHundred' :: (Num a, Ord a) => a -> Ordering  
compareWithHundred' = compare 100  

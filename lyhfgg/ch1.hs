doubleMe x = x + x
doubleUs x y = x*2 + y*2

doubleUs' x y = doubleMe x + doubleMe y

doubleSmallNumber x = if x > 100
                        then x
                        else x*2
doubleSmallNumber' x = (if x > 100 then x else x*2) + 1


lostNumbers = [4,8,15,16,23,42] 

triangles = [ (a, b, c) | c <- [1..10], b <- [1..10], a <- [1..10] ]
rightTriangles = [ (a, b, c) | c <- [1..10], b <- [1..10], a <- [1..10], a^2 + b^2 == c^2]
rightTriangles' = [ (a, b, c) | c <- [1..10], b <- [1..10], a <- [1..10], a^2 + b^2 == c^2, a+b+c == 24]

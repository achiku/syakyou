import Data.Char

main :: IO()
main = do
    contents <- readFile "haiku.txt"
    writeFile "haiku_upper.txt" (map toUpper contents)

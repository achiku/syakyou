import System.IO
import Data.Char

main = do
    contents <- readFile "haiku.txt"
    writeFile "haiku_upper.txt" (map toUpper contents)


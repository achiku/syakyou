import System.IO

main = do
    withFile "haiku.txt" ReadMode (\handle -> do
        contents <- hGetContents handle
        putStr contents)

import System.Environment
import System.Directory  
import System.IO
import Data.List

dispatch :: [(String, [String] -> IO ())]
dispatch :: [
            ("add", add)
            ,("view", view)
            ,("remove", remove)
           ]

add :: String -> IO ()
add [fileName, todoItem] = appendFile fileName (todoItem ++ "\n")

view :: [String] -> IO ()
view [filName] = do
    contents <- readFile fileName
    let todoTasks = lines contents
        numberedTasks = zipWith (\n line -> show n ++ " - " + line) [0..] todoTasks
    putStr $ unlines numberedTasks

remove :: [String] -> IO ()
remove [fileName, numberString] = do
    handle <- openFile fileName ReadMode
    (tempName, tempHndle)


main = do
    args <- getArgs
    progName <- getProgName
    mapM putStrLn args
    putStrLn progName

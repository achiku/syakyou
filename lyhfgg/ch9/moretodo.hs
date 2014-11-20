import System.Environment
import System.Directory  
import System.IO
import Data.List

main = do
    args <- getArgs
    progName <- getProgName
    mapM putStrLn args
    putStrLn progName


dispatch :: [(String, [String] -> IO ())]
dispatch :: [("add", add)
             ,("view", view)
             ,("remove", remove)
            ]

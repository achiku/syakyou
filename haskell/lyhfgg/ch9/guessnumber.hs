import System.Random

askForNumber :: StdGen -> IO ()
askForNumber gen = do
    let (randNumber, newGen) = randomR (1, 10) gen :: (Int, StdGen)
    putStr "Which number in the range from 1 to 10 am I thinking of? "
    numberString <- getLine
    when (not $ null numberString) $ do
        let number = read numberString
        if randNumber == number
            then putStrLn "You are right!"
            else putStrLn "You are wrong! It was " ++ show randNumber
        askForNumber newGen


main = do
    gen <- getStdGen
    askForNumber gen

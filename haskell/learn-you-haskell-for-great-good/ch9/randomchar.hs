import System.Random

main = do
    putStrLn "-- global value returned by getStdGen will not be updated"
    gen <- getStdGen
    putStrLn $ take 20 (randomRs ('a', 'z') gen)
    gen2 <- getStdGen
    putStrLn $ take 20 (randomRs ('a', 'z') gen2)

    putStrLn "-- workaround"
    gen <- getStdGen
    let randomChars = (randomRs ('a', 'z') gen)
        (first20, rest) = splitAt 20 randomChars
        (second20, _) = splitAt 20 rest
    putStrLn first20
    putStrLn second20

    putStrLn "-- use newStdGen"
    gen <- getStdGen
    putStrLn $ take 20 (randomRs ('a', 'z') gen)
    gen2 <- newStdGen
    putStrLn $ take 20 (randomRs ('a', 'z') gen2)

main = do
    return ()
    return "HAHAHA"
    line <- getLine
    return "BLAH BLAH"
    return 4
    putStrLn line

    a <- return "hell"
    b <- return "yeah!"
    putStrLn $ a ++ " " ++ b

    let a = "hell"
        b = "yeah!!"
    putStrL $ a ++ " " ++ b

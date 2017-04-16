import Test.Hspec
import Test.QuickCheck

main :: IO ()
main = hspec $
  describe "Prelude.head" $

    it "returns the first element of an *arbitrary* list" $
      property $ \x xs -> head (x:xs) == (x :: Int)

    it "returns the first element of a list" $
      head [23 ..] `shouldBe` (23 :: Int)

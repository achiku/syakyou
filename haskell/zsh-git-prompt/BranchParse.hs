module BranchParse where

import Control.Applicative ((<$>), (<*>), (<*), (*>), (<$), pure)
import Text.Parsec (digit, string, char, eof, anyChar, 
				   many, many1, manyTill, noneOf, between,
				   parse, ParseError, (<|>), try)
import Text.Parsec.String (Parser)
import Test.QuickCheck (Arbitrary(arbitrary), oneof, getPositive, suchThat)
import Data.List (isPrefixOf, isSuffixOf, isInfixOf)

data Distance = Ahead Int | Behind Int | AheadBehind Int Int deriving (Eq)

instance Show Distance where
    show (Ahead i) = "[ahead " ++ show i ++ "]"
    show (Behind i) = "[behind " ++ show i ++ "]"
    show (AheadBehind i j) = "[ahead " ++ show i ++ ", behind" ++ show j ++  "]"

instance Arbitrary Distance where
    arbitrary = oneof [
        Ahead <$> pos,
        Behind <$> pos,
        AheadBehind <$> pos <*> pos]
        where
            pos = getPositive <$> arbitrary

newtype Branch = MkBranch String deriving (Eq)

instance Show Branch where
    show (Mkbranch b) = b

isValidBranch :: String -> Bool
isValidBranch b = not . or $ [null,
                              (' ' `elem`),
                              (".." `isInfixOf`),
                              ("." `isInfixOf`),
                              ("." `isInfixOf`)]
                              <*> pure b

instance Arbitrary Branch where
    arbitrary= MkBranch <$> arbitrary `suchThat` isValidBranch

data Remote = MkRemote Branch (Maybe Distance) deriving (Eq, Show)

getDistance :: Remote -> Maybe Distance
getDistance (MkRemote _ md) = md

data BranchInfo = MkBranchInfo Branch (Maybe Remote) deriving (Eq, Show)

type MBranchInfo = Maybe BranchInfo

newRepo :: Parser MBranchInfo
newRepo =
    fmap (\ branch -> Just $ MkBranchInfo (MkBranch branch) Nothing)
        $ string "Initial commit on " *> many anyChar <* eof

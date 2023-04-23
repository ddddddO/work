module FizzBuzz
  ( fizzBuzz
  , fizzBuzz'
  , fizzBuzz''
  , fizzBuzz'''
  ) where

fizzBuzz :: [Int] -> [String]
fizzBuzz nums =
  let fizzBuzzNums = [a | a <- fizz nums, b <- buzz nums, a == b]  -- [0,15,30,45]
      fizzNums = [a | a <- fizz nums, not (a `elem` fizzBuzzNums)] -- [3,6,9,12,18,21,24,27,33,36,39,42,48]
      buzzNums = [a | a <- buzz nums, not (a `elem` fizzBuzzNums)] -- [5,10,20,25,35,40,50]
  in
    [if a `elem` fizzBuzzNums
      then "fizzbuzz"
    else if a `elem` fizzNums
      then "fizz"
    else if a `elem` buzzNums
      then "buzz"
    else
      show a | a <- nums]

fizz nums = [a | a <- nums, a `mod` 3 == 0]
buzz nums = [a | a <- nums, a `mod` 5 == 0]


fizzBuzz' :: [Int] -> [String]
fizzBuzz' nums =
  map (\a -> if a `mod` 15 == 0 then "fizzbuzz" else if a `mod` 3 == 0 then "fizz" else if a `mod` 5 == 0 then "buzz" else show a) nums


fizzBuzz'' :: [Int] -> [String]
fizzBuzz'' nums =
  let initTupls = zip nums (cycle [""])
      fizzBuzz = zipFizzBuzz (zipFizzBuzz initTupls (3, "fizz")) (5, "buzz")
  in
    map (\tupl -> if snd tupl == "" then show (fst tupl) else snd tupl) fizzBuzz

zipFizzBuzz :: [(Int, String)] -> (Int, String) -> [(Int, String)]
zipFizzBuzz numTupls fb =
  let fbNums = take (length numTupls) (cycle [fst fb])
  in
    zipWith (\tupl -> \a -> (fst tupl, (snd tupl) ++ (if (fst tupl) `mod` a == 0 then snd fb else ""))) numTupls fbNums


fizzBuzz''' :: [Int] -> [String]
fizzBuzz''' nums =
  let initNums = map (\a -> FBNum a "") nums
      fizzBuzz = toFizzBuzz (toFizzBuzz initNums (3, "fizz")) (5, "buzz")
  in
    toStrings fizzBuzz

data FBNum = FBNum Int String
instance Show FBNum where
  show (FBNum num str) = if str == "" then show num else str
toString :: FBNum -> String
toString (FBNum _ str) = str
toInt :: FBNum -> Int
toInt (FBNum num _) = num
toFizzBuzz :: [FBNum] -> (Int, String) -> [FBNum]
toFizzBuzz fbNums fb =
  map (\a -> if toInt a `mod` (fst fb) == 0 then FBNum (toInt a) ((toString a) ++ (snd fb)) else FBNum (toInt a) (toString a)) fbNums
toStrings :: [FBNum] -> [String]
toStrings fbNums = map (\a -> show a) fbNums

-- $ ghci
-- ghci> :l a.hs
-- ghci> main
main = do
  let max = 50
      nums = [0..max]

  print "----------------"
  print "   Fizz Buzz!   "
  print "----------------"

  let ret1 = fizzBuzz nums  -- ["fizzbuzz","1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz"]
      ret2 = fizzBuzz' nums -- ["fizzbuzz","1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz"]

  print (ret1)
  print (ret2)

  print "----------------"
  print "      Test      "
  print "----------------"
  print (if (ret1 == ret2) then "Ok!" else "Fail...")

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

fizz nums =
  [a | a <- nums, a `mod` 3 == 0]

buzz nums =
  [a | a <- nums, a `mod` 5 == 0]


fizzBuzz' :: [Int] -> [String]

fizzBuzz' nums =
  map (\a -> if a `mod` 15 == 0 then "fizzbuzz" else if a `mod` 3 == 0 then "fizz" else if a `mod` 5 == 0 then "buzz" else show a) nums

-- etc
square nums =
  map (\a -> a*a) nums

square' nums =
  [a^2 | a <- nums]

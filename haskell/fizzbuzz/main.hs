import FizzBuzz

-- $ ghci
-- ghci> :l main.hs fizzbuzz.hs
-- ghci> main
main = do
  let max = 50
      nums = [0..max]

  print "----------------"
  print "   Fizz Buzz!   "
  print "----------------"

  -- ["fizzbuzz","1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz"]
  let ret1 = fizzBuzz nums
      ret2 = fizzBuzz' nums
      ret3 = fizzBuzz'' nums
      ret4 = fizzBuzz''' nums

  print ret1
  print ret2
  print ret3
  print ret4

  print "----------------"
  print "      Test      "
  print "----------------"
  print (if (ret1 == ret2) && (ret3 == ret2) && (ret4 == ret3) then "Ok!" else "Fail...")


-- etc
square nums = map (\a -> a*a) nums
square' nums = [a^2 | a <- nums]

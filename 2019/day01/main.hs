module Main(main) where


import qualified System.IO as IO


main = do
    s <- IO.readFile "input1.txt"
    print (solve1 (masses s))
    print (solve2 (masses s))
    where
        masses s = map (\x -> read x :: Int) (words s)


solve1 input =
    sum (map fuel input)


solve2 input =
    sum (map fullFuel input)


fuel :: Int -> Int
fuel mass =
    (quot mass 3) - 2


fullFuel mass =
    (fuel mass) + (fuelForFuel (fuel mass))
    where
        fuelForFuel n =
            if res <= 0 then
                0
            else
                res + (fuelForFuel res)
            where
                res = fuel n

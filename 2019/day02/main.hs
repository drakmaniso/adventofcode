module Main(main) where


import qualified System.IO as IO
import qualified Data.List.Split as Split
import qualified Data.Char as Char


main = do
    source <- IO.readFile "input1.txt"
    -- print (solve1 test1)
    -- print (solve1 test2)
    -- print (solve1 test3)
    -- print (solve1 test4)
    -- print (solve1 test5)
    print ((solve1 . (prepare 12 02) . input) source)
    print (((solve2 19690720) . input) source)
    where
        test1 = [1,9,10,3,2,3,11,0,99,30,40,50]
        test2 = [1,0,0,0,99]
        test3 = [2,3,0,3,99]
        test4 = [2,4,4,5,99,0]
        test5 = [1,1,1,4,99,5,6,0,99]
        input source = map (read :: String -> Int)
            (Split.wordsBy (not . Char.isDigit) source)


solve1 :: [Int] -> Int
solve1 input =
    walk 0 input
    where
        walk :: Int -> [Int] -> Int
        walk pos program =
            case snd (splitAt pos program) of
                opcode : _ | opcode == 99 ->
                    head program
                opcode : x : y : res : _ ->
                    if opcode == 99
                        then head program
                        else walk (pos + 4) (run opcode x y res program)


solve2 :: Int -> [Int] -> Int
solve2 expected program =
    walk 0 0
    where
        walk i j =
            if (solve1 . (prepare i j)) program == expected
                then 100 * i + j
                else if i == 99
                    then if j == 99
                        then -1
                        else walk 0 (j+1)
                    else walk (i+1) j


prepare :: Int -> Int -> [Int] -> [Int]
prepare noun verb (hd : _ : _ : tl) =
    hd : noun : verb : tl


run :: Int -> Int -> Int -> Int -> [Int] -> [Int]
run opcode x y res program =
    case opcode of
        1 ->
            before ++ (((at x program) + (at y program)) : after)

        2 ->
            before ++ (((at x program) * (at y program)) : after)

    where
        (before, _ : after) = splitAt res program

        at pos program =
            (head . snd) (splitAt pos program)

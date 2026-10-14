local test = require "src.util.test"
local day01 = require "src.day02.main"

local input = [[7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9]]

local function TestPart1()
    test.runTest(day01.Part1, input, 2)
end

local function TestPart2()
    test.runTest(day01.Part2, input, 4)
end

return { TestPart1 = TestPart1, TestPart2 = TestPart2 }

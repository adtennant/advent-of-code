local test = require "src.util.test"
local day01 = require "src.day01.main"

local input = [[3   4
4   3
2   5
1   3
3   9
3   3]]

local function TestPart1()
    test.runTest(day01.Part1, input, 11)
end

local function TestPart2()
    test.runTest(day01.Part2, input, 31)
end

return { TestPart1 = TestPart1, TestPart2 = TestPart2 }

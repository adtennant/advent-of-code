local test = require "src.util.test"
local day01 = require "src.day02.main"

local input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

local function TestPart1()
    test.runTest(day01.Part1, input, 161)
end

local function TestPart2()
    test.runTest(day01.Part2, input, 4)
end

return { TestPart1 = TestPart1, TestPart2 = TestPart2 }

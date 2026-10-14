local luaunit = require 'lib.luaunit'

--- @generic T
--- @param f fun(): T
--- @return T
--- @return number
local function time(f)
    local start = os.clock()
    local result = f()
    local elapsed = os.clock() - start
    return result, elapsed
end

--- @param day integer
local function printHeader(day)
    print(string.format("# Day %d", day))
end

--- @param part integer
--- @param result  number
--- @param time number
local function printResult(part, result, time)
    print(string.format("Part %d = %d (int %.3fs)", part, result, time))
end

--- @alias Solution fun(input: file*): number

---@param solution Solution
---@param input file*
local function runPart(part, solution, input)
    input:seek("set", 0)

    local result, t = time(function() return solution(input) end)
    printResult(part, result, t)
end

--- @param day integer
--- @param module { Part1: Solution, Part2: Solution }
--- @param filename string
local function run(module, filename)
    local input = assert(io.open(filename, "r"))

    runPart(1, module.Part1, input)
    runPart(2, module.Part2, input)

    input:close()
end

local function runSolution(day)
    printHeader(day)

    local filename = string.format("day%02d", day)
    local ok, module = pcall(require, "src." .. filename .. ".main")

    if not ok then
        print("No solution found.")
        return
    end

    run(module, "input/" .. filename .. ".txt")
end

local function runTests(day)
    printHeader(day)

    local filename = string.format("day%02d", day)
    local ok, module = pcall(require, "src." .. filename .. ".tests")

    if not ok then
        print("No tests found.")
        return
    end

    local runner = luaunit.LuaUnit.new()
    runner:runSuiteByInstances({ { filename, module } })
end

local args = { ... }

if #args == 0 then
    for day = 1, 25 do
        runSolution(day)
    end
elseif #args == 1 then
    if arg[1] == "tests" then
        for day = 1, 25 do
            runTests(day)
        end
    else
        local day = arg[1]
        runSolution(day)
    end
elseif #args == 2 then
    if arg[1] == "tests" then
        local day = arg[2]
        runTests(day)
    end
end

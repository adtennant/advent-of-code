local util = require "src.util.list"

---@param file file*
---@return {x:integer, y:integer}[]
local function parseInstructions(file)
    local instructions = {}

    for line in file:lines() do
        local instruction = {}

        for x, y in string.gmatch(line, "mul%((%d+)%,(%d+)%)") do
            instruction = { x = tonumber(x), y = tonumber(y) }
            instructions[#instructions + 1] = instruction
        end
    end

    return instructions
end

--- @param instruction {x:integer, y:integer}
--- @return integer
local function calculate(instruction)
    return instruction.x * instruction.y
end

local function Part1(input)
    local instructions = parseInstructions(input)
    local total = util.foldLeft(
        instructions,
        function(total, instruction)
            return total + calculate(instruction)
        end,
        0
    )
    return total
end

local function Part2(input)
    local reports = parseReports(input)
    return count(
        reports,
        function(report)
            return isSafe(report) or IsTollerable(report)
        end
    )
end

return { Part1 = Part1, Part2 = Part2 }

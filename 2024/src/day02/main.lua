local util = require "src.util.list"

---@param file file*
---@return integer[][]
local function parseReports(file)
    local reports = {}

    for line in file:lines() do
        local report = {}

        for level in string.gmatch(line, "([^ ]+)") do
            report[#report + 1] = tonumber(level)
        end

        reports[#reports + 1] = report
    end

    return reports
end

--- @param report integer[]
local function allIncreasing(report)
    return util.forAll(
        util.sliding(report, 2),
        function(pair)
            return pair[1] < pair[2]
        end
    )
end

--- @param report integer[]
local function allDecreasing(report)
    return util.forAll(
        util.sliding(report, 2),
        function(pair)
            return pair[1] > pair[2]
        end
    )
end

--- @param report integer[]
local function maxDifference(report, min, max)
    return util.forAll(
        util.sliding(report, 2),
        function(pair)
            local diff = math.abs(pair[1] - pair[2])
            return diff >= min and diff <= max
        end
    )
end

local function isSafe(report)
    return (allIncreasing(report) or allDecreasing(report)) and maxDifference(report, 1, 3)
end

local function isTollerable(report)
    for i = 1, #report, 1 do
        local newReport = { table.unpack(report) }
        table.remove(newReport, i)

        if isSafe(newReport) then
            return true
        end
    end

    return false
end

local function Part1(input)
    local reports = parseReports(input)
    return util.count(reports, isSafe)
end

local function Part2(input)
    local reports = parseReports(input)
    return util.count(
        reports,
        function(report)
            return isSafe(report) or isTollerable(report)
        end
    )
end

return { Part1 = Part1, Part2 = Part2 }

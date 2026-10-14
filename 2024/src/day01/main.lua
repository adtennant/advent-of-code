local util = require "src.util.list"

---@param file file*
---@return integer[], integer[]
local function parseLocations(file)
    local left = {}
    local right = {}

    for line in file:lines() do
        for l, r in string.gmatch(line, "(%w+)   (%w+)") do
            left[#left + 1] = tonumber(l)
            right[#right + 1] = tonumber(r)
        end
    end

    return left, right
end

--- @class Pair
--- @field left integer
--- @field right integer

--- @param left integer[]
--- @param right integer[]
--- @return Pair[]
local function createPairs(left, right)
    table.sort(left)
    table.sort(right)

    local pairs = {}

    for i, l in ipairs(left) do
        local r = right[i]
        pairs[#pairs + 1] = { left = l, right = r }
    end

    return pairs
end

--- @param pairs Pair[]
--- @return integer[]
local function getDistances(pairs)
    return util.foldLeft(
        pairs,
        function(distances, pair)
            local distance = math.abs(pair.left - pair.right)
            distances[#distances + 1] = distance
            return distances
        end,
        {}
    )
end

local function countByLocation(locations)
    return util.groupMapReduce(
        locations,
        util.identity,
        function() return 1 end,
        function(count, value) return count + value end,
        0
    )
end

local function calculateScore(locations, counts)
    return util.foldLeft(
        locations,
        function(score, location)
            local count = counts[location] or 0
            return score + (location * count)
        end,
        0
    )
end

local function Part1(input)
    local left, right = parseLocations(input)
    local pairs = createPairs(left, right)
    local distances = getDistances(pairs)
    local totalDistance = util.sum(distances)
    return totalDistance
end

local function Part2(input)
    local left, right = parseLocations(input)
    local counts = countByLocation(right)
    local score = calculateScore(left, counts)
    return score
end

return { Part1 = Part1, Part2 = Part2 }

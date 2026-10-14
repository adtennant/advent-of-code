local luaunit = require "lib.luaunit"

---@param filename string?
---@return file*
local function readInput(filename)
    local current_path = string.sub(debug.getinfo(1, "S").source, 2, string.len("/main.lua") * -1)
    local filename = filename or "input.txt"

    local input = assert(io.open(current_path .. filename, "r"))
    return input
end

--- @generic T
--- @param value T
--- @return T
local function identity(value)
    return value
end

--- @generic K
--- @generic A
--- @generic B
--- @param list A[]
--- @param key fun(key: A): K
--- @param value fun(value: A): B
--- @param op fun(result: B, value: B): B
--- @param default B
--- @return table<K, B>
local function groupMapReduce(list, key, value, op, default)
    local result = {}

    for _, item in ipairs(list) do
        local k = key(item)
        local v = value(item)
        result[k] = op(result[k] or default, v)
    end

    return result
end

--- @generic A
--- @generic B
--- @param list A[]
--- @param op fun(result: B, value: A): B
--- @param default B
--- @return B
local function foldLeft(list, op, default)
    local result = default

    for _, item in ipairs(list) do
        result = op(result, item)
    end

    return result
end

--- @param list number[]
--- @return number
local function sum(list)
    return foldLeft(
        list,
        function(result, value) return result + value end,
        0
    )
end

--- @generic T
--- @param list T[]
--- @param p fun(value: T): boolean
--- @return boolean
local function forAll(list, p)
    for _, value in ipairs(list) do
        if not p(value) then
            return false
        end
    end

    return true
end

--- @generic T
--- @param list T[]
--- @param size integer
--- @return T[][]
local function sliding(list, size)
    local windowed = {}

    for i = 1, #list - size + 1, 1 do
        local window = {}

        for j = 0, size, 1 do
            window[#window + 1] = list[i + j]
        end

        windowed[#windowed + 1] = window
    end

    return windowed
end

--- @generic T
--- @param list T[]
--- @param p fun(value: T): boolean
--- @return integer
local function count(list, p)
    local total = 0

    for _, value in ipairs(list) do
        if p(value) then
            total = total + 1
        end
    end

    return total
end

return {
    readInput = readInput,
    identity = identity,
    groupMapReduce = groupMapReduce,
    foldLeft = foldLeft,
    sum = sum,
    forAll = forAll,
    sliding = sliding,
    count = count
}

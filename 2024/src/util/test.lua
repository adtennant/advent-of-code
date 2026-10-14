local luaunit = require "lib.luaunit"

--- @param solution Solution
--- @param input string
local function runTest(solution, input, expected)
    local file = io.tmpfile()
    file:write(input)
    file:seek("set", 0)

    local result = solution(file)
    file:close()

    luaunit.assertEquals(result, expected)
end

return {
    runTest = runTest
}

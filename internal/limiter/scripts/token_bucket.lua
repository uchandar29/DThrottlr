local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local requested = tonumber(ARGV[3])

local current_time = redis.call('TIME')
local current_timestamp = tonumber(current_time[1]) * 1000 + math.floor(tonumber(current_time[2]) / 1000)

local data = redis.call('HMGET', key, 'tokens', 'last_refill')
local tokens = tonumber(data[1])
local last_refill = tonumber(data[2])

if tokens == nil or last_refill == nil then
    tokens = capacity
    last_refill = current_timestamp
end

local elapsed_time = math.max(0, current_timestamp - last_refill)
local refill_tokens = elapsed_time * refill_rate / 1000
tokens = math.min(capacity, tokens + refill_tokens)

local allowed = 0
if tokens >= requested then
    tokens = tokens - requested
    allowed = 1
else
    allowed = 0
end

redis.call('HSET', key, 'tokens', tokens, 'last_refill', current_timestamp)
redis.call("EXPIRE", key, math.ceil(capacity / refill_rate) + 10)

return {allowed, math.floor(tokens)}
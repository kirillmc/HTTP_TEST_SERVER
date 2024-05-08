counter = 0
request = function()
   path = "/programs/1"
   counter = counter + 1
   return wrk.format(nil, path)
end

response = function(status, headers, body)
   if counter == 3000 then
      wrk.thread:stop()
   end
end
require "json"
require "crambda"

def handler(event : JSON::Any, context : Crambda::Context)
  pp context
  pp "Ale gil"
end

Crambda.run_handler(->handler(JSON::Any, Crambda::Context))

require "json"
require "crambda"

def handler(event : JSON::Any, context : Crambda::Context)
  pp context
  { statusCode: 200, headers: { "Content-Type" => "application/json" }, body: { ale: "gil" }.to_json }
end

Crambda.run_handler(->handler(JSON::Any, Crambda::Context))

require "sinatra"
require "sinatra/base"
require "json"
require "thin"

set :bind, "0.0.0.0"

before do
  request.path_info.sub! %r{/$}, ""
end

get "/" do
  "hello world"
end

get "/status" do
  content_type :json
  {
    :status  => 200
  }.to_json
end

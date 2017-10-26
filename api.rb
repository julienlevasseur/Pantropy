#!/usr/bin/env ruby

require 'sinatra'
require 'net/http'
require 'net/https'
require 'uri'
require 'json'

set :bind, '0.0.0.0'

class Provider

  def initialize(name, arguments)
    @name = name
    @arguments = arguments
  end
end

class Resource

  def initialize(type, name, arguments)
    @type = type
    @name = name
    @arguments = arguments
  end
end

class Argument:

  def initialize(name, required, description)
    @name = name
    @required = required
    @description = description
  end
end

def index
  data = [
    'aws',
    'chef',
    'openstack'
  ]
  return data.to_json
end

get '/' do
  haml :index, :locals => {
    :title => 'Pantropy',
    :bookmarks => index()
  }
end

get '/openstack' do
  
end

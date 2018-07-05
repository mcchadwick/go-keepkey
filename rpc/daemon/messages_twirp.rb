# Code generated by protoc-gen-twirp_ruby 1.0.0, DO NOT EDIT.
require 'twirp'
require_relative 'messages_pb.rb'

module Solipsis
  module Gkk
    module Daemon
      class DaemonService < Twirp::Service
        package 'solipsis.gkk.daemon'
        service 'Daemon'
        rpc :Ping, PingRequest, PingResponse, :ruby_method => :ping
      end

      class DaemonClient < Twirp::Client
        client_for DaemonService
      end
    end
  end
end

# rubocop:disable all
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hal9000/hal9000.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "hal9000.Message" do
    optional :text, :string, 1
    optional :user, :message, 2, "hal9000.User"
    optional :room, :string, 3
  end
  add_message "hal9000.User" do
    optional :email, :string, 1
    optional :name, :string, 2
  end
  add_message "hal9000.Response" do
    optional :match, :bool, 1
  end
  add_message "hal9000.CreateRepfixErrorRequest" do
    optional :hostname, :string, 1
    optional :error, :string, 2
    optional :mysql_last_error, :string, 3
  end
  add_message "hal9000.CreateRepfixErrorResponse" do
    optional :status, :int64, 1
    optional :body, :string, 2
  end
end

module Hal9000
  Message = Google::Protobuf::DescriptorPool.generated_pool.lookup("hal9000.Message").msgclass
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("hal9000.User").msgclass
  Response = Google::Protobuf::DescriptorPool.generated_pool.lookup("hal9000.Response").msgclass
  CreateRepfixErrorRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("hal9000.CreateRepfixErrorRequest").msgclass
  CreateRepfixErrorResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("hal9000.CreateRepfixErrorResponse").msgclass
end

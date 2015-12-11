path = require "path"
grpc = require "grpc"
protobuf = require "protobufjs"

protodir = path.resolve(__dirname + "/../../proto")
proto = protobuf.loadProtoFile(root: protodir, file: "openflights.proto")
openflights = grpc.loadObject(proto.ns).openflights
client = new openflights.API("0.0.0.0:1747", grpc.Credentials.createInsecure())

module.exports = (robot) ->
  robot.respond /airport (\w+)/i, (msg) ->
    code = msg.match[1]
    request = new openflights.GetAirportByCodeRequest(code)
    client.getAirportByCode request, (err, response) ->
      msg.reply "#{response.id} #{response.iata_faa}, #{response.city}, #{response.country}"
  robot.respond /distance (\w+) (\w+)/i, (msg) ->
    source = msg.match[1]
    destination = msg.match[2]
    request = new openflights.GetDistanceByIDRequest(source, destination)
    client.getDistanceById request, (err, response) ->
      msg.reply "#{response.value}"
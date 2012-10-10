# Description:
#   Send messages to channels via hubot
#
# Dependencies:
#   None
#
# Configuration:
#   HUBOT_CAT_PORT
#
# Commands:
#   None
#
# Notes:
#    echo "hello everyone" | nc -w1 127.0.0.1 7890
#    
# Author:
#   Berg

net = require "net"
util = require "../lib/util"

module.exports = (robot) ->
  try
    server = net.createServer((c) ->
      c.on "data", (data) ->
        msg  = data.toString().trim()
        user = { room: process.env.HUBOT_IRC_ROOMS }

        recordRelease msg
        robot.send user, msg
    )

    server.listen parseInt(process.env.HUBOT_CAT_PORT), ->
      console.log "cat listening"
  catch e
    console.log e
    console.log "CAT FAILED!"

# Record the relase infos
recordRelease = (msg) ->
  match = msg.match(/^(\w\w)\sjust\supdated\sPROD\son\semail-d1\.pardot\.com\sfrom\sRevision:\s(\d+)/i)

  if match isnt null
    util.getReleaseDBConn().query 'INSERT INTO sync (releaser, revision, date) VALUES(?, ?, NOW())', [match[1], match[2], new Date()], (err,r,f) ->
        return console.log err if err

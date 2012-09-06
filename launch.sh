#!/bin/sh

# Irc connection settings
export HUBOT_IRC_SERVER="supportbot.pardot.com"
export HUBOT_IRC_ROOMS="#testroom"
export HUBOT_IRC_PORT="7000"
export HUBOT_IRC_USESSL="true"
export HUBOT_IRC_NICK="Parbot"
export HUBOT_IRC_SERVER_FAKE_SSL="true"

# Ports to listen on
export HUBOT_CAT_PORT="7890"
export PORT="7891"

# Quote database settings
export QUOTE_DB_USER="pardot"
export QUOTE_DB_PASSWORD="pardot"
export QUOTE_DB_DATABASE="pardot_quotes"

export HUBOT_LOG_LEVEL="debug"  # This helps to see what Hubot is doing
export HUBOT_IRC_DEBUG="true"

# Supportbot settings
export SUPPORTBOT_EXECUTABLE="php /var/www/supportq/symfony"

# Finally run:
./bin/hubot -a irc
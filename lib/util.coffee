mysql = require "mysql"

request = require 'request'

# Get all nicknames from the robot
module.exports.getAllNicks = (message) ->
  channel = message.robot.adapter.bot.chans[message.message.user.room]
  return [] if not channel
  Object.keys channel.users

# Get connection to release DB
module.exports.getReleaseDBConn = () ->
    try
        client = mysql.createClient
            user: process.env.DB_USER,
            password: process.env.DB_PASSWORD,
            database: process.env.RELEASE_DATABASE,
            host: '127.0.0.1'

        client
    catch e
        console.log e
        false

# Get connection to quote DB
module.exports.getQuoteDBConn = () ->
    try
        client = mysql.createClient
            user: process.env.DB_USER,
            password: process.env.DB_PASSWORD,
            database: process.env.QUOTE_DATABASE,
            host: '127.0.0.1'

        client
    catch e
        console.log e
        false

# Get connection to kpi DB
module.exports.getKPIDBConn = () ->
    try
        client = mysql.createClient
            user: process.env.DB_USER
            password: process.env.DB_PASSWORD
            database: process.env.KPI_DATABASE,
            host: '127.0.0.1'

        client
    catch e
        console.log e
        false

# Get account information from internal API
module.exports.apiGetAccountInfo = (account_id, msg) ->
    options =
        url: "https://tools.pardot.com/accounts/#{account_id}"
        headers:
            api_key: '1b424ca119675c1c712957531498ac5af4646afb'

    request.get options, (error, response, body) ->
        if error
            console.log 'ERROR'
            console.log error.message
            msg.send "Sorry, unable to fetch info for account ID: #{account_id}"
        else
            # console.log body
            account_info = JSON.parse(body)
            if account_info.error
                msg.send "Sorry, unable to fetch info for account ID: #{account_id}"
            else
                msg.send "Company: #{account_info.company} (#{account_info.id}) | Shard: #{account_info.shard_id}"

module.exports.apiGetAccountsLike = (search_text, msg) ->
    # escaped_text = escape search_text
    escaped_text = encodeURIComponent search_text
    console.log escaped_text
    options =
        url: "https://tools.pardot.com/accounts/like/#{escaped_text}"
        headers:
            api_key: '1b424ca119675c1c712957531498ac5af4646afb'

    request.get options, (error, response, body) ->
        if error
            console.log 'ERROR'
            console.log error.message
            msg.send "Sorry, unable to fetch accounts matching \'#{search_text}\'"
        else
            console.log body
            try
                info_for_accounts = JSON.parse(body)
            catch e
                console.log e
                msg.send "Sorry, unable to fetch accounts matching \'#{search_text}\'"
                return

            if info_for_accounts.error
                msg.send "Sorry, unable to fetch accounts matching \'#{search_text}\'"
            else
                if info_for_accounts.length == 0
                    msg.send "No accounts found that matched \'#{search_text}\'"
                else
                    accounts_to_show = info_for_accounts.length
                    if accounts_to_show > 8
                        msg.send "#{accounts_to_show} accounts matched. Showing the first five."
                        accounts_to_show = 5

                    account_list = ''
                    for account, idx in info_for_accounts
                        if idx >= accounts_to_show
                            break
                        account_list = account_list + "Company: #{account.company} (#{account.id}) | Shard: #{account.shard_id}\n"

                    msg.send account_list

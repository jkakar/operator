class PullRequestTableMaker

  createPullRequestTable: (pullRequests) ->
    tableMarkup = "<table><tr><td><strong>Title</strong></td><td><strong>User</strong></td><td><strong>URL</strong></td><td><strong>Created Date</strong></td></tr>"
    for pr in pullRequests
      tableMarkup = tableMarkup + @tableRow(pr)
    return tableMarkup + "</table>"

  tableRow: (pullRequest) ->
    "<tr><td>#{pullRequest.title}</td><td>#{pullRequest.user.login}</td><td><a href = \"#{pullRequest.html_url}\">#{pullRequest.number}</a></td><td>#{pullRequest.created_at.substring(0, 10)}</td></tr>"

module.exports = PullRequestTableMaker
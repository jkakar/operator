class Ticket < ApplicationRecord
  TRACKER_JIRA = "jira".freeze

  has_one :ticket_reference
  has_one :multipass

  def self.synchronize_jira_ticket(issue_key)
    payload = Changeling.config.jira_client.Issue.find(issue_key).attrs
    issue = JIRAIssue.new(payload)

    ticket = Ticket.find_or_initialize_by(
      external_id: issue_key,
      tracker: Ticket::TRACKER_JIRA
    )
    ticket.summary = issue.summary
    ticket.status = issue.status
    ticket.open = issue.open?
    ticket.url = issue.url
    ticket.save!

    ticket
  rescue JIRA::HTTPError => e
    if e.code.to_i != 404
      raise
    end

    ticket = Ticket.find_by(
      external_id: issue_key,
      tracker: Ticket::TRACKER_JIRA
    )

    if ticket
      ticket.open = false
      ticket.save!
      ticket
    end
  end
end
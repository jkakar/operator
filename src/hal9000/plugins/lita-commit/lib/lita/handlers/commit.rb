module Lita
  module Handlers
    class Commit < Handler
      # The <project> group is optional in this regex, the commit method
      # will default the project to 'pardot' if no <project> is provided
      route /^commit(?:\s+(?<project>[a-z0-9\-]+))?\s+(?<sha>[a-f0-9]+)$/i, :commit, command: true, help: {
        "commit (project)? <commit sha>" => "Responds with the commit url for any repo, defaults to pardot"
      }

      # The <sha2> group is optional in this regex. The diff method will
      # default to comparing just <sha1> to master if no <sha2> provided
      route /^diff\s+(?<sha1>[^\s]+)(?:\s+(?<sha2>[^\s]+))?$/i, :diff, command: true, help: {
        "diff <sha1> <sha2>" => "Responds with the compare url for the Pardot repo"
      }

      def commit(response)
        sha = response.match_data["sha"]
        project = response.match_data["project"] || "pardot"

        # TODO: figure out a way to post html using hipchat room notification
        # msg = "Commit: #{project} - #{sha}"
        url = "https://git.dev.pardot.com/Pardot/" + project + "/commit/" + sha
        # html = "<a href=\"#{url}\">#{msg}</a>"

        response.reply(url)
      end

      def diff(response)
        sha2 = response.match_data["sha2"]
        diff = response.match_data["sha1"].tr("/", ";")
        diff += "..." + sha2.tr("/", ";") if sha2

        # msg = "Diff: #{diff}"
        url = "https://git.dev.pardot.com/Pardot/pardot/compare/" + diff + "?w=1"
        # html = "<a href=\"#{url}\">#{msg}</a>"

        response.reply(url)
      end

      Lita.register_handler(self)
    end
  end
end

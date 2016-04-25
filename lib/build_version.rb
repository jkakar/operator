# BuildVersion loads and writes a local file that tracks the latest deployment. Both
# sync_scripts and pull_agent use this concept.
#
# An example looks something like:
#
#     build1234
#     bfa9aac
#     https://artifactory.dev.pardot.com/some/url/path.tar.gz
#
# The first line is a build number (possibly prefixed with `build`)
# The second line is the git sha
# The third line is the artifact url

BuildVersion = Struct.new(:build_number, :sha, :artifact_url) do
  # Loads the current build version information from a file
  def self.load(filename)
    File.open(filename, "r") do |f|
      lines = f.readlines.map { |line| line.strip != "" && line.strip }
      new(lines[0].sub(/^build/, "").to_i, lines[1], lines[2])
    end
  rescue => e
    Logger.log(:warn, "Couldn't load sync state from #{filename}: #{e}")
    nil
  end

  def instance_of_deploy?(deploy)
    artifact_url && deploy.artifact_url == artifact_url
  end

  def to_s
    ["build#{build_number}", sha, artifact_url].join("\n")
  end
end

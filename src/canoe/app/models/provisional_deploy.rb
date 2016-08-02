require "base64"
require "json"

class ProvisionalDeploy
  attr_reader :artifact_url, :branch, :build_number, :sha, :passed_ci, :created_at
  attr_reader :options_validator, :options

  def self.from_artifact_url(project, artifact_url)
    hash = Artifactory.client.get(artifact_url, properties: nil)
    artifact = Artifactory::Resource::Artifact.from_hash(hash)
    properties = artifact.properties.each_with_object({}) { |(k, v), h| h[k] = v.first }

    from_artifact_url_and_properties(project, artifact_url, properties)
  end

  def self.from_artifact_hash(hash)
    artifact_url = Artifactory.client.build_uri(:get, "/" + ["api", "storage", hash["repo"], hash["path"], hash["name"]].join("/")).to_s
    properties = hash["properties"].each_with_object({}) { |p, h| h[p["key"]] = p["value"] }

    from_artifact_url_and_properties(hash["repo"], artifact_url, properties)
  end

  def self.from_artifact_url_and_properties(project, artifact_url, properties)
    return nil unless
      properties["gitBranch"] && \
      properties["buildNumber"] && \
      properties["gitSha"] && \
      properties["buildTimeStamp"]

    options_validator = \
      begin
        properties["optionsValidator"] && JSON.parse(Base64.decode64(properties["optionsValidator"]))
      rescue JSON::ParseError
        Instrumentation.log_exception($!, {
          at: "ProvisionalDeploy",
          fn: "from_artifact_url_and_properties",
        })
        nil
      end

    new(
      project: project,
      artifact_url: artifact_url,
      branch: properties["gitBranch"],
      build_number: properties["buildNumber"].to_i,
      sha: properties["gitSha"],
      passed_ci: !!(properties["passedCI"] && properties["passedCI"] == "true"),
      created_at: Time.parse(properties["buildTimeStamp"]),
      options_validator: options_validator,
      options: {}
    )
  end

  def self.from_previous_deploy(project, deploy)
    new(
      project: project,
      artifact_url: deploy.artifact_url,
      branch: deploy.branch,
      build_number: deploy.build_number,
      sha: deploy.sha,
      passed_ci: deploy.passed_ci,
      options_validator: deploy.options_validator,
      options: deploy.options
    )
  end

  def initialize(project:, artifact_url:, branch:, build_number:, sha:, passed_ci:, created_at: nil, options_validator: nil, options: {})
    @project = project
    @artifact_url = artifact_url
    @branch = branch
    @build_number = build_number
    @sha = sha
    @passed_ci = passed_ci
    @created_at = created_at
    @options_validator = options_validator
    @options = options
  end

  def project_name
    @project.name
  end

  def show_single_servers?
    !@project.all_servers_default
  end

  def valid?
    @sha.present?
  end

  def passed_ci?
    @passed_ci
  end
end

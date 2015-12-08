require "uri"
require "fetch_strategy_base"
require 'artifactory'

class FetchStrategyArtifactory < FetchStrategyBase
  include Artifactory::Resource
  attr_accessor :environment

  def initialize(environment)
    self.environment = environment

    Artifactory.configure do |config|
      config.username = environment.artifactory_user
      config.password = environment.artifactory_token

      proxy = URI(config.endpoint).find_proxy
      if proxy
        # If we connect through a proxy, we use an HTTP (non-SSL) URL. The proxy
        # intercepts this request and still talks to the upstream over SSL, but
        # it allows the response to be cached since it's decrypted by the
        # _proxy_ instead of _this host_.
        config.endpoint = 'http://artifactory.dev.pardot.com'
        config.proxy_address = proxy.hostname
        config.proxy_port = proxy.port
      else
        config.endpoint = 'https://artifactory.dev.pardot.com'
        config.ssl_verify = true
      end
    end
  end

  def valid?(deploy)
    return false unless deploy.artifact_url
    artifact = Artifact.from_url(deploy.artifact_url)
    artifact && artifact.properties["gitSha"]
  rescue Artifactory::Error::HTTPError
    false
  end

  def fetch(deploy)
    # returns path to fetched asset (file or directory)
    artifact = Artifact.from_url(deploy.artifact_url)
    artifact.download(environment.payload.artifacts_path)
  end
end

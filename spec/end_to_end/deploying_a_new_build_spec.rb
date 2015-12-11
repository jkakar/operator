require "cli"

describe "deploying a new build" do
  let(:current_build_number) { 1111 }
  let(:build_number) { 1234 }

  let(:current_sha) { "abc123" }
  let(:sha) { "bcd234" }

  let(:current_artifact_url) { "https://artifactory.dev.pardot.com/artifactory/api/storage/build1111.tar.gz" }

  let(:artifact_url) { "https://artifactory.dev.pardot.com/artifactory/api/storage/build1234.tar.gz" }
  let(:artifact_download_url) { "https://artifactory.dev.pardot.com/artifactory/build1234.tar.gz" }

  let(:tempdir) { Dir.mktmpdir("pull-agent") }

  after { FileUtils.rm_rf(tempdir) }

  before do
    stub_request(:get, "http://canoe.test/api/targets/test/deploys/latest?repo_name=pardot&server=#{ShellHelper.hostname}")
      .to_return(body: %({"id":445,"what":"branch","what_details":"master","artifact_url":"#{artifact_url}","build_number":#{build_number},"servers":{"#{ShellHelper.hostname}":{"stage":"pending","action":"deploy"}}}))

    bootstrap_repo_path(tempdir)
    current_version = BuildVersion.new(build_number, sha, current_artifact_url)
    File.write(File.join(tempdir, "current", "build.version"), current_version.to_s)
  end

  it "downloads the artifact, unpacks it, and switches over the symlink" do
    # API request for the Artifact
    stub_request(:get, %r{#{Regexp.escape(artifact_url)}(\?properties=)?}).
      to_return(
        status: 200,
        body: JSON.dump(
          uri: artifact_url,
          downloadUri: artifact_download_url,
          properties: {"gitSha" => [sha]},
        ),
        headers: {"Content-Type" => "application/json"}
      )

    # Download request for the Artifact
    stub_request(:get, artifact_download_url)
      .to_return(
        status: 200,
        body: empty_tar_gz_contents,
        headers: {"Content-Type" => "application/x-gzip"}
      )

    canoe_request = stub_request(:put, "http://canoe.test/api/targets/test/deploys/445/results/#{ShellHelper.hostname}")
      .to_return(status: 200)

    cli = CLI.new(%w[test pardot])
    cli.parse_arguments!
    cli.environment.payload.options[:repo_path] = tempdir

    expect(File.readlink(File.join(tempdir, "current"))).to match(%r{releases/A$})
    output = capturing_stdout { cli.checkin }

    expect(canoe_request).to have_been_made
    expect(File.readlink(File.join(tempdir, "current"))).to match(%r{releases/B$})
  end
end

require_relative "base"
require "helpers/salesedge"

module Environments
  class Staging < Base
    include SalesEdgeEnvModule

    restart_task :restart_autojobs,
      :restart_old_style_jobs,
      :restart_redis_jobs,
      only: :pardot

    after_deploy :restart_pithumbs_service, only: :pithumbs

    after_deploy :restart_salesedge, only: :'realtime-frontend'

    after_deploy :restart_workflowstats_service, only: :'workflow-stats'

    after_deploy :deploy_topology, only: :murdoc
  end

  register(:staging, Staging)
end

module Api
  class ProjectsController < Controller
    def index
      render json: Project.enabled.order(name: "asc").limit(100)
    end
  end
end

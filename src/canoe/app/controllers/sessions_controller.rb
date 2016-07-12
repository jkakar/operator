class SessionsController < ApplicationController
  skip_before_action :require_oauth_authentication, only: [:new, :create, :failure]

  # OmniAuth doesn't send us a CSRF token
  skip_before_action :verify_authenticity_token, only: [:create]

  def new
  end

  def create
    session.destroy

    user = AuthUser.find_or_create_by_omniauth(request.env["omniauth.auth"])
    if user && user.persisted?
      self.current_user = user
      redirect_to root_url
    else
      @errors = user.errors
      render action: "new"
    end
  end

  def destroy
    session.destroy
    redirect_to root_url
  end

  def failure
    session.destroy
    @auth_hash = request.env["omniauth.auth"]
    @failure_message = params[:message]
  end
end

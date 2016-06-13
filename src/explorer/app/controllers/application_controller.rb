class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception
  before_action :require_oauth_authentication

  rescue_from Exception do |exception|
    Instrumentation.log_exception(exception)

    if !Rails.env.development?
      render file: "public/500.html", layout: false, status: 500
    else
      raise exception
    end
  end

  around_action :log_context

  protected

  def log_context
    data = { request_id: Instrumentation.request_id }

    if current_user
      data[:user_email] = current_user.email
    end

    Instrumentation.context(data) do
      yield
    end
  end

  def datacenter
    @datacenter ||= DataCenter.new(
      Rails.application.config.x.datacenter,
      DatabaseConfigurationFile.load
    )
  end

  private

  def append_info_to_payload(payload)
    if !current_user
      return
    end

    payload[:context] = { user_email: current_user.email }
  end

  def require_oauth_authentication
    unless current_user.present?
      return redirect_to oauth_path
    end

    unless current_user.access_authorized?
      return redirect_to "/auth/unauthorized"
    end
  end

  def current_user
    return @current_user if defined?(@current_user)
    @current_user = load_current_user
  end
  helper_method :current_user

  def load_current_user
    if !session[:user_id]
      return nil
    end

    created_at = session[:created_at]
    if created_at && Time.zone.at(created_at) >= Rails.application.config.x.session_ttl.ago
      return User.find_by_id(session[:user_id])
    end

    session.destroy
    nil
  end

  def current_user=(user)
    @current_user = user
    session[:user_id] = user.id
    session[:created_at] = Time.now.to_i
  end

  def oauth_path
    case Rails.env
    when "development" then "/auth/developer"
    when "test" then "/auth/developer"
    else "/auth/ldap"
    end
  end
end

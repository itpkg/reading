class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception

  # locale
  before_action :set_locale

  def set_locale
    I18n.locale = params[:locale] || I18n.default_locale
  end

  def default_url_options
    {locale: I18n.locale}
  end

  def authenticate_admin_user!
    puts '#'*8, current_user.is_admin?
    raise SecurityError unless (current_user && current_user.is_admin?)
  end

  def after_sign_in_path_for(user)
    user.is_admin? ? admin_dashboard_path : root_path
  end

  def after_sign_out_path_for(resource_or_scope)
    root_path
  end


end

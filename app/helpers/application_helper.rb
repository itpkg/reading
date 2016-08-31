module ApplicationHelper
  def site_info(key)
    Setting["#{I18n.locale}://site//#{key}"] || "site.#{key}"
  end

  def dashboard_path
    Rails.env.production? ? '/dashboard' : 'http://localhost:4200/dashboard'
  end
end

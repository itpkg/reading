module ApplicationHelper
  def site_info(key)
    Setting["#{I18n.locale}://site//#{key}"] || "site.#{key}"
  end
end

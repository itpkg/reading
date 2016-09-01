class Api::SiteController < ApplicationController
  def info
    rst = {}
    %w(title subTitle keywords description copyright).each do |key|
      rst[key] = Setting["#{I18n.locale}://site//#{key}"] || "site.#{key}"
    end
    author = {}
    %w(name email).each do |key|
      author[key] = Setting["://site//author/#{key}"] || "site.author.#{key}"
    end
    rst[:author] = author

    rst[:lang] = I18n.locale

    render json: rst
  end
end
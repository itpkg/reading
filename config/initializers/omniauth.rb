Rails.application.config.middleware.use OmniAuth::Builder do
  provider :github, ENV['GITHUB_KEY'], ENV['GITHUB_SECRET'], scope: 'email,profile'

  provider :facebook, ENV['FACEBOOK_KEY'], ENV['FACEBOOK_SECRET']

  provider :google_oauth2, ENV['GOOGLE_KEY'], ENV['GOOGLE_SECRET']

  provider :weibo, ENV['WEIBO_KEY'], ENV['WEIBO_SECRET']

  provider :wechat, ENV['WECHAT_APP_ID'], ENV['WECHAT_APP_SECRET']

  provider :tqq, ENV['TQQ_KEY'], ENV['TQQ_SECRET']
  provider :qq_connect, ENV['QQ_CONNECT_KEY'], ENV['QQ_CONNECT_SECRET']
end
Rails.application.routes.draw do

  scope '(:locale)', locale: /en|zh-CN/ do
    get 'home', to: 'home#index'

    devise_for :users
  end

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root to: 'home#index'
end

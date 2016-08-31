Rails.application.routes.draw do

  scope '(:locale)', locale: /en|zh-CN/ do
    get 'home', to: 'home#index'

    namespace :cms do
      resources :tags
      resources :articles
      resources :comments
    end

    devise_for :users
  end

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root to: 'home#index'
end

Rails.application.routes.draw do


  scope '(:locale)', locale: /en|zh-CN/ do
    get 'home', to: 'home#index'
    get 'home/about'
    get 'home/help'
    get 'home/faq'

    namespace :cms do
      resources :tags
      resources :articles
      resources :comments
    end

    namespace :epub do
      resources :books, only: [:index, :show, :destroy]
      get 'pages/:bid/*name', to:'pages#show', as: :page
    end

    mount_devise_token_auth_for 'User', at: 'auth'
  end

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root to: 'home#index'
end

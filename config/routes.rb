Rails.application.routes.draw do


  get 'notices/index'

  scope '(:locale)', locale: /en|zh-CN/ do
    get 'home', to: 'home#index'
    get 'home/about'
    get 'home/help'
    get 'home/faq'

    resources :notices, only: [:index] #TODO

    #todo
    namespace :cms do
      resources :tags, only: [:index, :show]
      resources :articles, only: [:index, :show]
      resources :comments, only: [:index, :show]
    end

    namespace :epub do
      resources :books, only: [:index, :show]
      get 'pages/:bid/*name', to:'pages#show', as: :page
    end

    namespace :api do
      resources :notices #todo
    end

    mount_devise_token_auth_for 'User', at: 'api/auth'

  end

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root to: 'home#index'
end

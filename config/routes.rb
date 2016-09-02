# require 'sidekiq/web'
# Sidekiq::Web.set :session_secret, Rails.application.secrets[:secret_key_base]

Rails.application.routes.draw do

  #admin
  devise_for :users, ActiveAdmin::Devise.config
  ActiveAdmin.routes(self)


  scope '(:locale)', locale: /en|zh-CN/ do

    #home todo
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
      get 'pages/:bid/*name', to: 'pages#show', as: :page
    end


    namespace :api do
      get 'site/info'
      resources :notices #todo
      resources :leave_words, only: [:index, :create, :destroy] #todo

      mount_devise_token_auth_for 'User', at: 'auth', controllers: {
          sessions: 'auth/sessions',
          registrations: 'auth/registrations',
      }
    end


  end

  # sidekiq monitoring
  # authenticate :user, lambda { |u| u.admin? } do
  #   mount Sidekiq::Web => '/sidekiq'
  # end

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  root to: 'home#index'
end

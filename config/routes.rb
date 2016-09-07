Reading::Engine.routes.draw do
  get 'members'=> 'members#index'

  resources :books, only: [:index, :show, :destroy]
  resources :notes

  get 'home/index'

  get 'page/:bid/*name', to: 'home#page', as: :page

  root to: 'home#index'
end

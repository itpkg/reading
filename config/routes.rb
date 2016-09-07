Reading::Engine.routes.draw do
  get 'members'=> 'members#index'

  resources :books, only: [:index, :destroy]
  resources :notes

  get 'home/index'

  root to: 'home#index'
end

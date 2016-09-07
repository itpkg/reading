Reading::Engine.routes.draw do

  resources :members, only: [:new, :create, :show]
  resources :books, only: [:index, :show, :destroy]
  resources :notes

  get 'page/:book_id/*name', to: 'home#page', as: :page

  root to: 'home#index'
end

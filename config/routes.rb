Reading::Engine.routes.draw do

  resources :members, only: [:new, :create, :show]
  resources :books, only: [:index, :show, :destroy]
  resources :notes

  get 'page/:book_id/*name', to: 'home#page', as: :page

  %w(dict search).each do |a|
    get a => "home##{a}"
    post a => "home##{a}"
  end


  root to: 'home#index'
end

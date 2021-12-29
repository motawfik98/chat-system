Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
  post "/applications/:token/chats/:number/search", to: 'messages#index'
end

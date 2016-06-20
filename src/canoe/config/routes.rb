Rails.application.routes.draw do
  resources :sessions, only: [:new] do
    delete :destroy, on: :collection
  end
  get "/auth/:provider/callback", to: "sessions#create"
  post "/auth/:provider/callback", to: "sessions#create"
  get "/auth/failure", to: "sessions#failure"

  resources :projects, param: :name, only: [:index, :show] do
    resources :tags, param: :name, only: [:index] do
      get :latest, on: :collection
    end
    resources :branches, param: :name, constraints: { name: /.*/ }, only: [:index] do
      resources :builds, only: [:index]
    end

    resources :deploys, only: [:new, :create, :show] do
      get :select_target, on: :collection
      post :force_to_complete, on: :member
      post :pick_new_restart_servers, on: :member
      post :cancel, on: :member
    end
  end

  resources :targets, param: :name, only: [:show] do
    resources :projects, param: :name, only: [] do
      resources :deploys, only: [:index]

      post :lock, on: :member
      post :unlock, on: :member
    end
  end
  resources :servers, only: [:index, :new, :create, :edit, :update]

  namespace :api, defaults: { format: "json" } do
    # legacy
    post "deploy/:id/completed_server" => "deploys#completed_server"

    resources :targets, param: :name, only: [] do
      resources :projects, param: :name, only: [] do
        resources :deploys, only: [:index]
      end

      resources :deploys, only: [:show] do
        resources :results, param: :hostname, constraints: { hostname: /.*/ }, only: [:update]
        get :latest, on: :collection
        post :completed_server, on: :member
      end
    end

    resources :projects, param: :name, only: [] do
      resources :branches, param: :name, constraints: { name: /.*/ }, only: [] do
        resources :builds, only: [:index]
      end
    end
  end

  post "/api/chef/checkin", to: "api/chef_deploys#checkin"
  post "/api/chef/complete_deploy", to: "api/chef_deploys#complete_deploy"

  get "/_boomtown", to: "projects#boomtown"
  root to: "projects#index"
end

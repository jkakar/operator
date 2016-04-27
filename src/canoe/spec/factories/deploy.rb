FactoryGirl.define do
  factory :deploy do
    what "branch"
    what_details "master"
    sha "abc123"
    completed true
    repo_name "pardot"

    association :auth_user, factory: :user
    association :deploy_target, factory: :deploy_target
  end
end
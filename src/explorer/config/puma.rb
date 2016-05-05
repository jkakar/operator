workers Integer(ENV.fetch("WEB_CONCURRENCY", 2))
threads_count = Integer(ENV.fetch("MAX_THREADS", 5))
threads threads_count, threads_count

preload_app!

rackup      DefaultRackup
port        ENV.fetch("PORT", 4000)
environment ENV.fetch("RACK_ENV", "development")

log = File.expand_path("../../log/#{ENV.fetch("RACK_ENV")}.log", __FILE__)
stdout_redirect log, log, true

on_worker_boot do
  # Worker specific setup for Rails 4.1+
  # See: https://devcenter.heroku.com/articles/deploying-rails-applications-with-the-puma-web-server#on-worker-boot
  ActiveRecord::Base.establish_connection
end

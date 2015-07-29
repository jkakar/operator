# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20140904164525) do

  create_table "auth_users", force: :cascade do |t|
    t.string   "email",      limit: 255
    t.string   "name",       limit: 255
    t.string   "uid",        limit: 255
    t.string   "token",      limit: 255
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "auth_users", ["email"], name: "index_auth_users_on_email", using: :btree

  create_table "deploy_targets", force: :cascade do |t|
    t.string   "name",            limit: 255
    t.string   "script_path",     limit: 255
    t.string   "lock_path",       limit: 255
    t.boolean  "locked"
    t.integer  "locking_user_id", limit: 4
    t.datetime "created_at"
    t.datetime "updated_at"
    t.boolean  "lockable",                    default: false
  end

  add_index "deploy_targets", ["name"], name: "index_deploy_targets_on_name", using: :btree

  create_table "deploys", force: :cascade do |t|
    t.integer  "deploy_target_id",  limit: 4
    t.integer  "auth_user_id",      limit: 4
    t.string   "repo_name",         limit: 255
    t.string   "what",              limit: 255
    t.string   "what_details",      limit: 255
    t.boolean  "completed",                       default: false
    t.datetime "created_at"
    t.datetime "updated_at"
    t.string   "process_id",        limit: 255
    t.boolean  "canceled",                        default: false
    t.integer  "server_count",      limit: 4,     default: 0
    t.text     "servers_used",      limit: 65535
    t.text     "specified_servers", limit: 65535
    t.text     "completed_servers", limit: 65535
    t.text     "sha",               limit: 65535
  end

  add_index "deploys", ["deploy_target_id"], name: "index_deploys_on_deploy_target_id", using: :btree

  create_table "locks", force: :cascade do |t|
    t.integer  "deploy_target_id", limit: 4
    t.integer  "auth_user_id",     limit: 4
    t.boolean  "locking",                    default: false
    t.boolean  "forced",                     default: false
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "locks", ["auth_user_id"], name: "index_locks_on_auth_user_id", using: :btree
  add_index "locks", ["deploy_target_id"], name: "index_locks_on_deploy_target_id", using: :btree

  create_table "target_jobs", force: :cascade do |t|
    t.integer  "deploy_target_id", limit: 4
    t.integer  "auth_user_id",     limit: 4
    t.string   "job_name",         limit: 255
    t.string   "command",          limit: 255
    t.string   "process_id",       limit: 255
    t.boolean  "completed",                    default: false
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "target_jobs", ["auth_user_id"], name: "index_target_jobs_on_auth_user_id", using: :btree
  add_index "target_jobs", ["deploy_target_id"], name: "index_target_jobs_on_deploy_target_id", using: :btree

end

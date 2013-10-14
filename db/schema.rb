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

ActiveRecord::Schema.define(version: 20131010212003) do

  create_table "auth_users", force: true do |t|
    t.string   "email"
    t.string   "name"
    t.string   "uid"
    t.string   "token"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "auth_users", ["email"], name: "index_auth_users_on_email", using: :btree

  create_table "deploy_targets", force: true do |t|
    t.string   "name"
    t.string   "script_path"
    t.string   "lock_path"
    t.boolean  "locked"
    t.integer  "locking_user_id"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "deploy_targets", ["name"], name: "index_deploy_targets_on_name", using: :btree

  create_table "deploys", force: true do |t|
    t.integer  "deploy_target_id"
    t.integer  "auth_user_id"
    t.string   "repo_name"
    t.string   "what"
    t.string   "what_details"
    t.boolean  "completed"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "deploys", ["deploy_target_id"], name: "index_deploys_on_deploy_target_id", using: :btree

end

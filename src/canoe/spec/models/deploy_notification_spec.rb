require "rails_helper"

RSpec.describe DeployNotification do
  let(:deploy) { FactoryGirl.create(:deploy) }
  let(:notification) { FactoryGirl.create(:deploy_notification) }
  let(:notifier) { FakeHipchatNotifier.new }

  before do
    notification.notifier = notifier
  end

  describe "#notify_deploy_start" do
    it "notifies a HipChat room" do
      notification.notify_deploy_start(deploy)
      expect(notifier.messages.size).to eq(1)
      expect(notifier.messages[0].room_id).to eq(notification.hipchat_room_id)
      expect(notifier.messages[0].message).to match(/just began syncing/)
    end
  end

  describe "#notify_deploy_complete" do
    it "notifies a HipChat room" do
      notification.notify_deploy_complete(deploy)
      expect(notifier.messages.size).to eq(1)
      expect(notifier.messages[0].room_id).to eq(notification.hipchat_room_id)
      expect(notifier.messages[0].message).to match(/just finished syncing/)
    end
  end

  describe "#notify_deploy_cancelled" do
    it "notifies a HipChat room" do
      notification.notify_deploy_cancelled(deploy)
      expect(notifier.messages.size).to eq(1)
      expect(notifier.messages[0].room_id).to eq(notification.hipchat_room_id)
      expect(notifier.messages[0].message).to match(/CANCELLED syncing/)
    end
  end

  describe "#notify_untested_deploy" do
    it "notifies a HipChat room" do
      notification.notify_untested_deploy(deploy)
      expect(notifier.messages.size).to eq(1)
      expect(notifier.messages[0].room_id).to eq(notification.hipchat_room_id)
      expect(notifier.messages[0].message).to match(/UNTESTED deploy/)
    end
  end
end
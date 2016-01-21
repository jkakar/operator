require 'ipaddr'

TRUSTED_NETWORKS = [
  '127.0.0.1/32',
  '204.14.236.0/24',   # aloha-east
  '204.14.239.0/24',   # aloha-west
  '174.37.191.2/32',   # proxy.dev
  '169.45.0.88/32',    # squid-d4
  '136.147.104.20/30'  # pardot-proxyout1-{1,2,3,4}-dfw
].map { |i| IPAddr.new(i) }.freeze

Rack::Attack.whitelist('ip whitelist') do |req|
  remote_ip = IPAddr.new(req.ip)
  req.path == '/_ping' || TRUSTED_NETWORKS.any? { |net| net.include?(remote_ip) }
end

Rack::Attack.blacklist('deny from internet') do |req|
  true
end

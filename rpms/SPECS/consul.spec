%define consul_ver 0.7.0
%define data_dir /opt/consul
%define bin_dir /usr/local/bin/

Summary: Consul is a tool for service discovery and configuration. Consul is distributed, highly available, and extremely scalable.
Name: consul
Version: %{consul_ver}
Release: 2.pardot%{?dist}
Group: System Environment/Daemons
License: MPLv2.0
URL: http://www.consul.io
Source0: https://releases.hashicorp.com/%{name}/%{version}/%{name}_%{version}_linux_amd64.zip
Source1: https://releases.hashicorp.com/%{name}/%{version}/%{name}_%{version}_web_ui.zip
BuildRoot: %(mktemp -ud %{_tmppath}/%{name}-%{version}-%{release}-XXXXXX)

%package ui
Summary: Consul Web UI
Requires: consul = %{version}
BuildArch: noarch

%description
Consul is a tool for service discovery and configuration. Consul is distributed, highly available, and extremely scalable.

Consul provides several key features:
 - Service Discovery - Consul makes it simple for services to register themselves and to discover other services via a DNS or HTTP interface. External services such as SaaS providers can be registered as well.
 - Health Checking - Health Checking enables Consul to quickly alert operators about any issues in a cluster. The integration with service discovery prevents routing traffic to unhealthy hosts and enables service level circuit breakers.
 - Key/Value Storage - A flexible key/value store enables storing dynamic configuration, feature flagging, coordination, leader election and more. The simple HTTP API makes it easy to use anywhere.
 - Multi-Datacenter - Consul is built to be datacenter aware, and can support any number of regions without complex configuration.

%description ui
Consul comes with support for a beautiful, functional web UI. The UI can be used for viewing all services and nodes, viewing all health checks and their current status, and for reading and setting key/value data. The UI automatically supports multi-datacenter.

%prep
%setup -q -c -b 1

%install
mkdir -p %{buildroot}/%{bin_dir}
cp consul %{buildroot}/%{bin_dir}
mkdir -p %{buildroot}/%{data_dir}
mkdir -p %{buildroot}/%{data_dir}/%{name}-ui
cp -r index.html static %{buildroot}/%{data_dir}/%{name}-ui

%clean
rm -rf %{buildroot}

%files
%defattr(-,root,root,-)
%dir %attr(755, root, root) %{data_dir}
%attr(755, root, root) %{bin_dir}/consul

%files ui
%defattr(-,root,root,-)
%config(noreplace) %attr(-, root, root) %{data_dir}/%{name}-ui

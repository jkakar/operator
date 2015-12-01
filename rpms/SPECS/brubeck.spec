Name: brubeck
Version: 1.0.0
Release: 5%{?dist}
Summary: A Statsd-compatible metrics aggregator
Group: Applications/Internet
License: MIT
URL: https://github.com/github/brubeck
Source0: https://git.dev.pardot.com/Pardot/brubeck/archive/3bd41eda71a0988f3ab58f2af4a4fb56f6d95e47.tar.gz
BuildRequires: gcc
BuildRequires: git
BuildRequires: jansson-devel >= 2.5
BuildRequires: libmicrohttpd-devel
BuildRequires: openssl-devel
Requires: jansson >= 2.5
Requires: libmicrohttpd
Requires: openssl
BuildRoot: %{_tmppath}/%name-root

%description
A Statsd-compatible metrics aggregator developed by GitHub.

%prep
%setup -q -n brubeck-3bd41eda71a0988f3ab58f2af4a4fb56f6d95e47
git clone -b 1a84d49c3ca794356f015a3391e6d10be98f6a6a https://github.com/concurrencykit/ck vendor/ck

%build
make brubeck

%install
rm -rf $RPM_BUILD_ROOT
install -d -m 755 $RPM_BUILD_ROOT/usr/bin
install -m 755 brubeck $RPM_BUILD_ROOT/usr/bin/brubeck

%files
%defattr(-,root,root,-)
/usr/bin/brubeck

%clean
rm -rf $RPM_BUILD_ROOT

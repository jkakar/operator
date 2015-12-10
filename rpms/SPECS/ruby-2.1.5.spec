%global gem_version 2.1.0

Name: ruby-2.1.5
Version: 2.1.5
Release: 2%{?dist}
Summary: An interpreter of object-oriented scripting language
Group: Development/Languages
License: BSD-2-Clause
URL: https://www.ruby-lang.org/
Source0: http://cache.ruby-lang.org/pub/ruby/2.1/ruby-%{version}.tar.gz
BuildRequires: autoconf
BuildRequires: gdbm-devel
BuildRequires: ncurses-devel
BuildRequires: libffi-devel
BuildRequires: openssl-devel
BuildRequires: libyaml-devel
BuildRequires: readline-devel
BuildRequires: tk-devel
BuildRoot: %{_tmppath}/%name-root

%description
Ruby is the interpreted scripting language for quick and easy
object-oriented programming.  It has many features to process text
files and to do system management tasks (as in Perl).  It is simple,
straight-forward, and extensible.

%prep
%setup -q -n ruby-%{version}

%build
./configure --prefix=/opt/rubies/ruby-%{version} --exec-prefix=/opt/rubies/ruby-%{version}
make

%install
rm -rf $RPM_BUILD_ROOT
make install DESTDIR=$RPM_BUILD_ROOT

# install bundler
GEM_HOME=$RPM_BUILD_ROOT/opt/rubies/ruby-%{version}/lib/ruby/gems/%{gem_version} \
  $RPM_BUILD_ROOT/opt/rubies/ruby-%{version}/bin/ruby \
  -I $RPM_BUILD_ROOT/opt/rubies/ruby-%{version}/lib/ruby/%{gem_version} \
  -I $RPM_BUILD_ROOT/opt/rubies/ruby-%{version}/lib/ruby/%{gem_version}/%{_arch}-linux \
  $RPM_BUILD_ROOT/opt/rubies/ruby-%{version}/bin/gem install bundler -v 1.10.6

%files
%defattr(-,root,root,-)
/opt/rubies/ruby-%{version}

%clean
rm -rf $RPM_BUILD_ROOT

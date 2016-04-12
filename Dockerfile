FROM teampass/teampass:2.1.25.2 
# You'll need to update docker-entrypoint.sh if this version is changed

COPY docker-entrypoint.sh /
COPY apache-default.conf /etc/apache2/sites-available/000-default.conf

EXPOSE 80

ENTRYPOINT ["/docker-entrypoint.sh"]

# This doesn't seem to be working, so launching from ENTRYPOINT instead
#CMD ["/bin/sh -c \"/usr/sbin/apache2ctl -D FOREGROUND &; tail -f /var/log/apache2/*log\""]
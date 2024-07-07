FROM nginx:latest
RUN echo Europe/Berlin >/etc/timezone && ln -sf /usr/share/zoneinfo/Europe/Berlin /etc/localtime && dpkg-reconfigure -f noninteractive tzdata && apt-get update && apt-get --no-install-recommends -y install logrotate
RUN apt-get -yqq install openssl
RUN openssl dhparam -out /etc/ssl/private/dhparam.pem 2048
RUN mkdir -p /var/cache/nginx/upload; chown -R www-data:www-data /var/cache/nginx/
COPY sites-enabled.nginx /etc/nginx/sites-enabled/custom-host
COPY ssl/cert.crt /etc/nginx/ssl/cert.crt
COPY ssl/cert.key /etc/nginx/ssl/cert.key
COPY nginx.conf /etc/nginx/nginx.conf
COPY fastcgi.conf /etc/nginx/conf.d/fastcgi.conf
COPY proxy.conf /etc/nginx/conf.d/proxy.conf
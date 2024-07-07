FROM php:8.3-fpm

RUN apt-get update && apt-get install -y pkg-config patch

# Install developer dependencies
RUN echo Europe/Berlin > /etc/timezone && ln -sf /usr/share/zoneinfo/Europe/Berlin /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
RUN apt-get update -yqq && apt-get install -y git curl zip unzip libzip-dev libsqlite3-dev libxml2-dev libicu-dev libfreetype6-dev libmcrypt-dev libjpeg62-turbo-dev libcurl4-gnutls-dev libbz2-dev libssl-dev imagemagick -yqq

# Install php extensions

# Configure gd package
RUN docker-php-ext-configure gd
RUN docker-php-ext-install mysqli pdo pdo_mysql curl gd opcache sockets xml zip

COPY php.ini /usr/local/etc/php/
COPY www.conf /usr/local/etc/php-fpm.d/www.conf

RUN curl -sSL https://getcomposer.org/installer | php -- --install-dir=/usr/bin --filename=composer

WORKDIR /var/www
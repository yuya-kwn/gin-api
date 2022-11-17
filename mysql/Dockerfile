FROM mysql:latest

WORKDIR /var/lib/mysql

EXPOSE 3306

ADD ./my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
# Environment Setup

### Update Apt-Get
```shell
sudo apt-get update  
sudo apt-get upgrade
```

### Install Node
Only if you want to run the Express server to be sure that things are working
```shell
curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
sudo apt-get install -y nodejs
```

### Install Redis
```shell
sudo apt-get -y install redis-server
sudo nano /etc/redis/redis.conf
    Change default port (I used port 6400)
sudo service redis-server restart
    Restart Redis
redis-cli -p 6400 ping
    Should recieve 'PONG' response
```

### Install Apache
```shell
sudo apt-get install apache2
sudo a2enmod rewrite
sudo nano /etc/apache2/apache2.conf
    AllowOverride All
    Require all granted
  Create '.htaccess' in '/var/www/html'
    RewriteEngine on
    RewriteRule ^i$ ingest.php [NC]
sudo service apache2 restart
    Create 'php.log' in '/var/www/html'
sudo chmod 777 php.log
    This sets permissions for Apache to write to the log
```

### Install PHP CLI
```shell
sudo apt-get install php5-cli
```
This isn't needed, but nice to test file and make sure it works

### Install PHP
```shell
sudo apt-get install php5 libapache2-mod-php5
sudo nano /etc/apache2/mods-enabled/dir.conf
    Move index.php to be the first file listed after DirectoryIndex
```

### Install PHP Redis
```shell
sudo apt-get install php5-redis
sudo service apache2 restart
```

### Test PHP Install
```shell
sudo nano /var/www/html/info.php
    '<?php
    phpinfo();
    ?>'
    Visit 'http://localhost/info.php' in browser to check if it works
```

### Install GO
```shell
sudo curl -O https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
sudo tar -xf go1.6.2.linux-amd64.tar.gz
sudo mv go /usr/local
sudo nano ~/.profile
    Add the line 'export PATH=$PATH:/usr/local/go/bin'
source ~/.profile
mkdir $HOME/work
  export GOPATH=$HOME/work
go get gopkg.in/redis.v4  
```

### Install Git
```shell
apt-get install git
git config --global user.name 'Your Name'
git config --global user.email 'youremail@domain.com'
```
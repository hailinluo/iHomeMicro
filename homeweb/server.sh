redis-server ./conf/redis.conf

fdfs_trackerd  /home/vagrant/go/src/iHomeMicroMicro/homeweb/conf/tracker.conf restart

fdfs_storaged  /home/vagrant/go/src/iHomeMicroMicro/homeweb/conf/storage.conf restart

sudo nginx

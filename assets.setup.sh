sudo unlink /etc/systemd/system/assets.service
sudo ln -s /home/ueda/assets/assets.service /etc/systemd/system
sudo mkdir -p /var/log/atelierueda
sudo mkdir -p /var/log/assets

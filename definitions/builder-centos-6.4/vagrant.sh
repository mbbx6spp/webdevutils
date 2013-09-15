USERNAME=builder
GROUPNAME=builder
PASSWORD=vagrant
# Vagrant specific
date > /etc/vagrant_box_build_time

# Add builder user
/usr/sbin/groupadd $GROUPNAME
/usr/sbin/useradd $USERNAME -g $GROUPNAME -G wheel
echo "$PASSWORD"|passwd --stdin $USERNAME
echo "$USERNAME        ALL=(ALL)       NOPASSWD: ALL" >> /etc/sudoers.d/builder
chmod 0440 /etc/sudoers.d/builder

# Installing vagrant keys
mkdir -pm 700 /home/$USERNAME/.ssh
wget --no-check-certificate 'https://raw.github.com/mitchellh/vagrant/master/keys/vagrant.pub' -O /home/$USERNAME/.ssh/authorized_keys
chmod 0600 /home/$USERNAME/.ssh/authorized_keys
chown -R $USERNAME /home/$USERNAME/.ssh

# Customize the message of the day
echo 'Welcome to your Vagrant-built virtual machine.' > /etc/motd

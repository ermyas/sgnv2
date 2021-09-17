USERNAME="ubuntu"

sudo mkdir -p /var/log/crdb
sudo chmod 755 /var/log/crdb
sudo chown $USERNAME.$USERNAME /var/log/crdb
sudo -u $USERNAME touch /var/log/crdb/out.log
sudo -u $USERNAME touch /var/log/crdb/err.log

sudo tee /etc/systemd/system/crdb.service << EOF
[Unit]
Description=CockroachDB single node
After=network-online.target
[Service]
WorkingDirectory=/home/ubuntu/cbridge/cockroach
ExecStart=/usr/local/bin/cockroach start-single-node --insecure --listen-addr=localhost:26257 \
  --http-addr=localhost:18080 --store=path=/home/ubuntu/cbridge/cockroach
StandardOutput=append:/var/log/crdb/out.log
StandardError=append:/var/log/crdb/err.log
Restart=always
RestartSec=10
User=$USERNAME

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl enable crdb.service
sudo systemctl start crdb.service
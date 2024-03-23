# Setting up this service on your local machine
- Open the project and add your *peronal Access token.*
- Build the project using `go build` command.
- Now, Open a terminal and create a new systemd unit file for your service:
   ```
  sudo nano /etc/systemd/system/github-notifications.service
  ```
  Add the following content to the file:
  ```
  [Unit]
  Description=GitHub Notifications Service
  After=network.target

  [Service]
  Type=simple
  Restart=always
  RestartSec=5
  ExecStart="/usr/bin/github-notifications"

  [Install]
  WantedBy=multi-user.target
  ```
  Note: Adjust the ExecStart path if necessary to point to the location of your compiled Go program.
  Save the file and exit the text editor.
- Then, enable the service to start automatically at boot:
   ```
   sudo systemctl enable github-notifications.service
   ```
  ```
   sudo systemctl start github-notifications.service
  ```
- verify the status
  ```
  sudo systemctl status github-notifications
  ```

- Additionally, for monitoring the service logs, use this command:
  ```
   journalctl -u github-notifications -f
  ```
  

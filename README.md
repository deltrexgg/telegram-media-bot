<h2>telegram-media-bot</h2>
Media storage bot with directory management leveraging telegram storage via chat.

<h3>Project Structure</h3>
telegram-media-bot/
├── cmd/
├── internal/
├── migrations/
├── data/               # Persistent directory (SQLite DB lives here)
├── Dockerfile
├── .env
├── go.mod
├── go.sum
├── README.md

<h3>Build this app</h3>

The data/ directory is critical. It is mounted into the container so the database survives container restarts and deletion.

Environment Variables

Create a .env file in the project root:

TELEGRAM_API=your_telegram_bot_token
DATABASE_FILE=/data/database.db


Important notes:

DATABASE_FILE must point to /data/database.db

/data is mounted from the host using a Docker volume

Build Docker Image

Run this from the project root directory:

sudo docker build -t telegram-media-bot .


Verify the image exists:

sudo docker images | grep telegram-media-bot

Run the Container With Persistence
sudo docker run -d \
  --name telegram-media-bot \
  --restart unless-stopped \
  --env-file .env \
  -v $(pwd)/data:/data \
  telegram-media-bot

<h3>Commands Available</h3>

• Make a folder:
  /makefolder foldername

• Upload files to a folder:
  /addfiles
  (send files, then type /stop)

• View all files in a folder:
  /getfiles

• View latest uploaded files:
  /latest

• Share folder access:
  /share

• Delete folder:
  /delete


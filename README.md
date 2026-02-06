<h2>telegram-media-bot</h2>

<p>
Media storage bot with directory management, leveraging Telegram chat storage as the backend.
</p>

<hr>

<h3>Project Structure</h3>

<pre>
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
</pre>

<p>
The <code>data/</code> directory is critical. It is mounted into the container so the database survives
container restarts and deletion.
</p>

<hr>

<h3>Environment Variables</h3>

<p>Create a <code>.env</code> file in the project root:</p>

<pre>
TELEGRAM_API=your_telegram_bot_token
DATABASE_FILE=/data/database.db
</pre>

<p><strong>Notes:</strong></p>
<ul>
  <li><code>DATABASE_FILE</code> must point to <code>/data/database.db</code></li>
  <li><code>/data</code> is mounted from the host using a Docker volume</li>
</ul>

<hr>

<h3>Build the Application</h3>

<p>Run this from the project root directory:</p>

<pre>
sudo docker build -t telegram-media-bot .
</pre>

<p>Verify the image exists:</p>

<pre>
sudo docker images | grep telegram-media-bot
</pre>

<hr>

<h3>Run the Container With Persistence</h3>

<pre>
sudo docker run -d \
  --name telegram-media-bot \
  --restart unless-stopped \
  --env-file .env \
  -v $(pwd)/data:/data \
  telegram-media-bot
</pre>

<p>This setup ensures:</p>
<ul>
  <li>Environment variables are loaded securely</li>
  <li>SQLite database persists outside the container</li>
  <li>The container restarts automatically unless stopped manually</li>
</ul>

<hr>

<h3>Available Commands</h3>

<h4>Folder Management</h4>
<ul>
  <li>
    <strong>Create a folder</strong><br>
    <code>/makefolder foldername</code>
  </li>
  <li>
    <strong>Delete a folder</strong><br>
    <code>/delete</code>
  </li>
</ul>

<h4>File Management</h4>
<ul>
  <li>
    <strong>Upload files to a folder</strong><br>
    <code>/addfiles</code><br>
    Send files, then type <code>/stop</code>
  </li>
  <li>
    <strong>View all files in a folder</strong><br>
    <code>/getfiles</code>
  </li>
  <li>
    <strong>View latest uploaded files</strong><br>
    <code>/latest</code>
  </li>
</ul>

<h4>Sharing</h4>
<ul>
  <li>
    <strong>Share folder access</strong><br>
    <code>/share</code>
  </li>
</ul>



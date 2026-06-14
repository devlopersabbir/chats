const messagesDiv = document.getElementById("messages");
const messageInput = document.getElementById("messageInput");
const sendBtn = document.getElementById("sendBtn");
const statusSpan = document.getElementById("status");

// REPLACE THIS with your actual WebSocket server address
const WS_URL = "ws://localhost:9090/ws";

let socket;

function connect() {
  try {
    socket = new WebSocket(WS_URL);

    socket.onopen = function (event) {
      updateStatus("Connected");
      messageInput.disabled = false;
      sendBtn.disabled = false;
      addSystemMessage("Connected to server");
      messageInput.focus();
    };

    socket.onmessage = function (event) {
      console.log(event);
      // When a message arrives, append it to the chat
      addMessage(event.data, "received");
    };

    socket.onclose = function (event) {
      updateStatus("Disconnected");
      messageInput.disabled = true;
      sendBtn.disabled = true;
      addSystemMessage("Connection lost");
    };

    socket.onerror = function (error) {
      console.error("WebSocket Error:", error);
      addSystemMessage("Error occurred");
    };
  } catch (e) {
    console.error("Connection failed:", e);
    addSystemMessage("Failed to connect");
  }
}

// Send message function
function sendMessage() {
  const text = messageInput.value.trim();
  if (text && socket && socket.readyState === WebSocket.OPEN) {
    socket.send(text);
    // Add my own message to the UI immediately
    addMessage(text, "sent");
    messageInput.value = "";
  }
}

// UI Helpers
function addMessage(text, type) {
  const msgDiv = document.createElement("div");
  msgDiv.classList.add("message", type);
  msgDiv.textContent = text;
  messagesDiv.appendChild(msgDiv);
  scrollToBottom();
}

function addSystemMessage(text) {
  const msgDiv = document.createElement("div");
  msgDiv.classList.add("message", "system");
  msgDiv.textContent = text;
  messagesDiv.appendChild(msgDiv);
  scrollToBottom();
}

function updateStatus(text) {
  statusSpan.textContent = text;
  statusSpan.style.color = text === "Connected" ? "#4caf50" : "#f44336";
}

function scrollToBottom() {
  messagesDiv.scrollTop = messagesDiv.scrollHeight;
}

// Event Listeners
sendBtn.addEventListener("click", sendMessage);

messageInput.addEventListener("keypress", function (e) {
  if (e.key === "Enter") {
    sendMessage();
  }
});

// Initiate connection on load
connect();

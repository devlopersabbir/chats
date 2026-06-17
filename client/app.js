const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = () => {
  document.getElementById("output").innerText += "Connected!\n";
};

socket.onmessage = (event) => {
  console.log(event);
  document.getElementById("output").innerText += "Server: " + event.data + "\n";
};

function sendMsg() {
  const msg = document.getElementById("msg").value;
  socket.send(msg);
}

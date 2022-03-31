import React from 'react';
// import io from 'socket.io-client';

const App = () => {

  const [message, setMessage] = React.useState("");
  const [socket, setSocket] = React.useState(null);

  React.useEffect(() => {
    const _socket = new WebSocket("ws://localhost:8000/ws");
    setSocket(_socket)
  }, []);

  const sendMessage = (e) => {
    e.preventDefault();
    socket.send(message);
    setMessage("");
  }

  return (
    <div>
      <h1>App</h1>
      <form onSubmit={(e) => sendMessage(e)}>
        <input value={message} onChange={e => setMessage(e.target.value)} />
      </form>
    </div>
  )

}

export default App;
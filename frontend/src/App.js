import React from 'react';
import MessageScreen from './MessageScreen';
import JoinScreen from './JoinScreen';

const SCREENS = Object.freeze({
	JOIN: "join",
	MESSAGE: "message"
});

const App = () => {
	
	const [socket, setSocket] = React.useState(null);
	const [screen, setScreen] = React.useState(SCREENS.JOIN);
	const [users, setUsers] = React.useState([]);
	const [messages, setMessages] = React.useState([]);

	React.useEffect(() => {
		const ws = new WebSocket("ws://127.0.0.1:8000/ws");
		setSocket(ws);
	}, []);

	React.useEffect(() => {
		if (!socket) return;
		socket.addEventListener("message", ({data}) => {
			const { message, users: _users } = JSON.parse(data);
			if (_users && _users.length) {
				setUsers(_users);
			}
			console.log({ message })
			setMessages([...messages, message]);
		})
	}, [socket]);

	const renderScreen = () => {
		switch (screen) {
			case SCREENS.MESSAGE:
				return <MessageScreen 
					users={users}
					onSend={(username, message) => {
						socket.send(JSON.stringify({ username, type: "direct", message }));
					}}
				/>;
			default:
				return <JoinScreen onJoin={(username) => {
					socket.send(JSON.stringify({ username, type: "broadcast" }));
					setScreen(SCREENS.MESSAGE);
				}}/>
		}
	}

	return <div className="container">
		{renderScreen()}
		<h5>Messages</h5>
		<ul>
		{messages.length > 0 && messages.map(m => <li key={m}>{m}</li>)}
		</ul>
	</div>
}

export default App;


// import React from 'react';
// // import io from 'socket.io-client';

// const App = () => {

//   const [message, setMessage] = React.useState("");
//   const [socket, setSocket] = React.useState(null);

//   React.useEffect(() => {
//     // const _socket = new WebSocket("ws://localhost:8000/ws");
//     // setSocket(_socket)
//   }, []);

//   const sendMessage = (e) => {
//     e.preventDefault();
//     socket.send(message);
//     setMessage("");
//   }

//   return (
//     <div>
//       <h1>App</h1>
//       <form onSubmit={(e) => sendMessage(e)}>
//         <input value={message} onChange={e => setMessage(e.target.value)} />
//       </form>
//     </div>
//   )

// }

// export default App;
import React from 'react';

const App = () => {
	
	const [ state, setState ] = React.useState({
		username: "",
		message: "",
	});

	const submitHandler = (e) => {
		e.preventDefault();
		console.log(state)
	}
	return (
		<div className="container my-5 py-5">
			<h3>Go Chat!</h3>
			<form id="input-form" className=" my-4" onSubmit={submitHandler}>
				<div className="form-group">
					<input
						type="text"
						className="form-control"
						placeholder="Enter username"
						value={state.username}
						onChange={(e) => setState(old => ({ ...old, username: e.target.value }))}
					/>
				</div>
				<div className="form-group">
					<input
						type="text"
						className="form-control"
						placeholder="Enter chat text here"
						value={state.message}
						onChange={(e) => setState(old => ({ ...old, message: e.target.value }))}
					/>
				</div>
				<button className="btn btn-primary btn-block" type="submit">Send</button>
	  		 </form>
  		 </div>
	)
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
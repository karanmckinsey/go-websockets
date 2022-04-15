import React from 'react'; 

const MessageScreen = ({ users, onSend }) => {
	const [selectedUser, setSelectedUser] = React.useState("");
	const [message, setMessage] = React.useState("");

	const submitHandler = (e) => {
		if (!message || !selectedUser) return;
		e.preventDefault();
		console.log({ selectedUser, message });
		onSend(selectedUser, message);
	}

	return (<ul>
		{ users.map(u => (
			<li 
				key={u} 
				style={{ 
					fontWeight: `${selectedUser === u} && 800`,
					cursor: 'pointer'
				}}
				onClick={() => setSelectedUser(u)}
			>
				{u}
			</li>
		))}
		{
			selectedUser && <form id="input-form" className=" my-4" onSubmit={submitHandler}>
				<div className="form-group">
					<input
						type="text"
						className="form-control"
						placeholder="Enter your message"
						value={message}
						onChange={(e) => setMessage(e.target.value)}
					/>
				</div>
				<button className="btn btn-primary btn-block" type="submit">Send</button>
			</form>
		}
	</ul>)
}

export default MessageScreen;
import React from 'react'; 

const JoinScreen = ({ onJoin }) => {
	const [username, setUsername] = React.useState("");

	const submitHandler = (e) => {
		e.preventDefault();
		onJoin(username)
	} 
	return (
		<form id="input-form" className=" my-4" onSubmit={submitHandler}>
			<div className="form-group">
				<input
					type="text"
					className="form-control"
					placeholder="Enter your username"
					value={username}
					onChange={(e) => setUsername(e.target.value)}
				/>
			</div>
			<button className="btn btn-primary btn-block" type="submit">JOIN</button>
		</form>
	)
}

export default JoinScreen;
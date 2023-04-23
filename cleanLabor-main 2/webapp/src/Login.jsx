import React, {useState} from "react";

// Login Page [Yeshan Li]
export const Login = (props) => {
    // User information variables [Yeshan Li]
    const [username, setUsername] = useState('');
    const [pass, setPass] = useState('');
    const handleSubmit = (e) => {
        e.preventDefault();
        console.log(username);
        console.log(pass);
    }

    // HTTP POST login request to the backend go server [Yeshan Li]
    function loginPOST()  {
        // Simple POST request with a JSON body using fetch [Yeshan Li]
        const requestOptions = {
            method: 'POST',
            body: JSON.stringify({UserName: username, Desc: "", WorkContent: "", Phone: "", Email: "", Password: pass})
        };
        fetch('http://localhost:8080/loginTaskerUser', requestOptions)
            .then(response => response.json())
            .then(data => {
                // Decode the response data and display the corresponding result [Yeshan Li]
                console.log(data)
                console.log(data["status"])
                console.log(data["reason"])
                if (data["status"] == "success") {
                    alert("Successfully login the account! ")
                    console.log("Successfully login the account! ")
                } else{
                    alert("Failed to login the account, because " + data["reason"])
                    console.log("Failed to login the account, because " + data["reason"])
                }
            });
    }
    return (
        <div>
            {/*Login form [Yeshan Li]*/}
            <p>Login: </p>
            <form onSubmit={handleSubmit}>
                <label htmlFor="username">Username: </label>
                <br/><input value={username} onChange={(e) => setUsername(e.target.value)} type="username" placeholder="your user name" id="username" name="username"/>
                <br/>
                <label htmlFor="password">Password: </label>
                <br/><input value={pass} onChange={(e) => setPass(e.target.value)} type="password" placeholder="your password" id="password" name="password"/>
                <br/><button type={"submit"} onClick={() => loginPOST()}>Log in</button>
            </form>
            <button onClick={() => props.onFormSwitch('register')}>Register here</button>
        </div>
    )
}
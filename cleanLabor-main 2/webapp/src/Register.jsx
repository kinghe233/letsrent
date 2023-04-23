import React, {useState, Component} from "react";

// Register Page [Yeshan Li]
export const Register = (props) => {
    // User information variables [Yeshan Li]
    const [username, setUsername] = useState('');
    const [desc, setDesc] = useState('');
    const [workcontent, setWorkcontent] = useState('');
    const [phone, setPhone] = useState('');
    const [email, setEmail] = useState('');
    const [pass, setPass] = useState('');
    const handleSubmit = (e) => {
        e.preventDefault();
        console.log(username);
        console.log(pass);
    }
    // HTTP POST register request to the backend go server [Yeshan Li]
    function registerPOST()  {
        // Simple POST request with a JSON body using fetch [Yeshan Li]
        const requestOptions = {
            method: 'POST',
            body: JSON.stringify({UserName: username, Desc: desc, WorkContent: workcontent, Phone: phone, Email: email, Password: pass})
        };
        fetch('http://localhost:8080/registerTaskerUser', requestOptions)
            .then(response => response.json())
            .then(data => {
                // Decode the response data and display the corresponding result [Yeshan Li]
                console.log(data)
                console.log(data["status"])
                console.log(data["reason"])
                if (data["status"] == "success") {
                    alert("Successfully created the account! ")
                    console.log("Successfully created the account! ")
                } else{
                    alert("Failed to created the account, because " + data["reason"])
                    console.log("Failed to created the account, because " + data["reason"])
                }
            });
    }
    return (
        <div>
            {/*Register form [Yeshan Li]*/}
            <p>Register: </p>
            <form onSubmit={handleSubmit}>
                <label htmlFor="username">Username: </label>
                <br/><input value={username} onChange={(e) => setUsername(e.target.value)} type="username" placeholder="your user name" id="username" name="username"/>
                <br/>
                <label htmlFor="desc">Desc: </label>
                <br/><input value={desc} onChange={(e) => setDesc(e.target.value)} type="desc" placeholder="your desc" id="desc" name="desc"/>
                <br/>
                <label htmlFor="workcontent">Work Content: </label>
                <br/><input value={workcontent} onChange={(e) => setWorkcontent(e.target.value)} type="workcontent" placeholder="your workcontent" id="workcontent" name="workcontent"/>
                <br/>
                <label htmlFor="email">Email: </label>
                <br/><input value={email} onChange={(e) => setEmail(e.target.value)} type="email" placeholder="your email" id="email" name="email"/>
                <br/>
                <label htmlFor="phone">Phone: </label>
                <br/><input value={phone} onChange={(e) => setPhone(e.target.value)} type="phone" placeholder="your phone" id="phone" name="phone"/>
                <br/>
                <label htmlFor="password">Password: </label>
                <br/><input value={pass} onChange={(e) => setPass(e.target.value)} type="password" placeholder="your password" id="password" name="password"/>
                <br/><button type={"submit"} onClick={() => registerPOST()}>Register</button>
            </form>
            <button onClick={() => props.onFormSwitch('login')}>Already have an account? Login here</button>
        </div>
    )
}
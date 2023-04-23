import React, {useEffect, useState} from 'react';
import './App.css';
import Logo from './Logo';
import {Tech} from "./tech/Tech";
import {Login} from "./Login";
import {Register} from "./Register";


export function App() {
    // Login/Register Switch [Yeshan Li]
    const [currentForm, setCurrentForm] = useState('login');
    const toggleForm = (formName) => {
        setCurrentForm(formName);
    }
    return (
        <div className="app">
            <h2 className="title">my-app</h2>
            <div className="logo"><Logo/></div>
            {/*Login/Register Switch [Yeshan Li]*/}
            <div>
                {
                    currentForm === "login" ? <Login onFormSwitch={toggleForm}></Login> : <Register onFormSwitch={toggleForm}></Register>
                }
            </div>
        </div>
    );
}

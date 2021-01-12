import React from "react";
import {BrowserRouter as Router, Route} from "react-router-dom";
import {Container} from 'semantic-ui-react'
import 'semantic-ui-css/semantic.min.css';
import "./App.css";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import MenuBar from "./components/MenuBar";

const App = () => {
    return (
        <Router>
            <Container>
                <MenuBar/>
                <Route exact path="/" component={Home}></Route>
                <Route exact path="/login" component={Login}></Route>
                <Route exact path="/register" component={Register}></Route>
            </Container>
        </Router>
    );
}

export default App;

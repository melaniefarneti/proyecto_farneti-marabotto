import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './Components/Home.jsx'
import Register from "./Components/Register.jsx";


  const App = () => {

  return (
      <Router>
      <div>
        <h1>Mi aplicacion</h1>
          <Switch>
              <Route exact path="/" component={Home} />
              <Route path="/register" component={Register} />
          </Switch>
      </div>
      </Router>
  );
};

export default App;

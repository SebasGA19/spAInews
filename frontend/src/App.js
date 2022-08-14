import logo from './logo.svg';
import { Login } from './components/login';
import { Register } from './components/register';
import './App.css';
import { ConfirmRegister } from './components/confirm-register';

function App(){
  return(
    <div>
      <Register/>
      <ConfirmRegister/>
      <Login/>
    </div>
  )
}
/*
function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}
*/
export default App;

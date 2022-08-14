import { login } from "../api/login";

function handleLogin() {
    const username = document.getElementById("login-username").value;
    const password = document.getElementById("login-password").value;
    login(username, password)
    .then(
        session => {
            alert(`My session is ${session}`);
        }
    )
    .catch(
        error => {
            alert(`Something goes wrong: ${error}`);
        }
    );
}

export function Login(){
    return(
        <div>
            <h3>Login</h3>
            <form onSubmit={handleLogin}>
                <input required type="text" placeholder="username" id="login-username"/>
                <input required type="password" placeholder="password" id="login-password"/>
                <button type="submit">Submit</button>
            </form>
        </div>
    )
  }
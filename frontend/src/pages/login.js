import { Page } from "../components/page";
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
        <Page>
            <h3>Login</h3>
            <form onSubmit={handleLogin}>
                <div class="mb-3">
                    <label for="login-username" class="form-label">Username</label>
                    <input required type="text" placeholder="username" id="login-username" class="form-control"/>
                </div>
                <div class="mb-3">
                    <label for="login-password" class="form-label">Password</label>
                    <input required type="password" placeholder="password" id="login-password" class="form-control"/>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
        </Page>
    )
  }
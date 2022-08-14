import { Page } from "../components/page";
import { register } from "../api/register";

function handleRegister() {
    const username = document.getElementById("register-username").value;
    const email = document.getElementById("register-email").value;
    const password = document.getElementById("register-password").value;
    const confirmPassword = document.getElementById("register-confirm-password").value;
    if (password !== confirmPassword) {
        alert("Las contrasenas no coinciden");
        return false;
    }
    register(username, email, password)
    .then(
        () => {
            alert("A confirmation email was sent")
        }
    )
    .catch(
        error => {
            alert(`Error message is ${error}`)
        }
    )

}

export function Register() {
    return(
        <Page>
            <h3>Register</h3>
            <form onSubmit={handleRegister}>
                <div class="mb-3">
                    <label for="register-username" class="form-label">Username</label>
                    <input required type="text" placeholder="username" id="register-username" class="form-control"/>
                </div>
                <div class="mb-3">
                    <label for="register-email" class="form-label">Email</label>
                    <input required type="email" placeholder="email" id="register-email" class="form-control"/>
                </div>
                <div class="mb-3">
                    <label for="register-password" class="form-label">Password</label>
                    <input required type="password" placeholder="password" id="register-password" class="form-control"/>
                </div>
                <div class="mb-3">
                    <label for="register-confirm-password" class="form-label">Confirm Password</label>
                    <input required type="password" placeholder="confirm password" id="register-confirm-password" class="form-control"/>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
        </Page>
    );
}
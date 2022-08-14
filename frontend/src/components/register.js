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
        <div>
            <h3>Register</h3>
            <form onSubmit={handleRegister}>
                <input required type="text" placeholder="username" id="register-username"/>
                <input required type="email" placeholder="email" id="register-email"/>
                <input required type="password" placeholder="password" id="register-password"/>
                <input required type="password" placeholder="confirm password" id="register-confirm-password"/>
                <button type="submit">Submit</button>
            </form>
        </div>
    );
}
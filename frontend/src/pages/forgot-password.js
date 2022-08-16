import { Page } from "../components/page";
import { forgot_password} from "../api/forgot-password";

function handleForgotPassword(event) {
    event.preventDefault();
    const email = document.getElementById("user-email").value;
    forgot_password(email)
    .then(
        () => {
            alert("An email with the URL for password reset has been sent")
        }
    )
    .catch(
        error => {
            alert(`Error message is ${error}`)
        }
    )
}

export function ForgotPassword() {
    return(
        <Page>
            <div class="container text-center w-25">
                <h3>Reset your password</h3>

                <form onSubmit={handleForgotPassword}>
                    <div class="mb-3">
                        <label for="email-address" class="form-label">Email-address</label>
                        <input required type="text" placeholder="email" id="user-email" class="form-control"/>
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>

            </div>
        </Page>
    )
}
import { Page } from "../components/page";
import { forgot_password_confirm} from "../api/forgot-password";

function handleForgotPasswordConfirm() {
    const newPassword = document.getElementById("new-password").value;
    forgot_password_confirm(newPassword)
    .then(
        () => {
            alert("Your Password has been changed correctly")
        }
    )
    .catch(
        error => {
            alert(`Error message is ${error}`)
        }
    )
}

export function ForgotPasswordConfirm() {
    return(
        <Page>
            <div class="container text-center w-25">
                <h3>Insert your new password</h3>

                <form onSubmit={handleForgotPasswordConfirm}>
                    <div class="mb-3">
                        <label for="New password" class="form-label">New password</label>
                        <input required type="password" placeholder="password" id="new-password" class="form-control"/>
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>

            </div>
        </Page>
    )
}
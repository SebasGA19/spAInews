import { Page } from "../components/page"
import { confirm_change_email } from "../api/change-email";

function handleConfirm(event) {
    event.preventDefault();
    const confirmCode = document.getElementById("email-confirm-code").value;
    confirm_change_email(confirmCode)
    .then(
        () => {
            alert(`Email changed successfully`);
        }
    )
    .catch(
        error => {
            alert(`Something goes wrong: ${error}`);
        }
    );
}

export function ConfirmChangeEmail(){
    return(
        <Page>
            <div class="container text-center w-25">
            <h3>Confirm Change email</h3>
            <form onSubmit={handleConfirm}>
                <div class="mb-3">
                <label for="register-confirm-code" class="form-label">Confirmation code</label>
                    <input required type="text" placeholder="confirm-code" id="email-confirm-code" class="form-control"/>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
            </div>
        </Page>
    )
}
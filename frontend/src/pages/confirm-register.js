import { Page } from "../components/page";
import { confirmRegister } from "../api/confirm";

function handleConfirm() {
    const confirmCode = document.getElementById("register-confirm-code").value;
    confirmRegister(confirmCode)
    .then(
        () => {
            alert(`Account confirmed successfully`);
        }
    )
    .catch(
        error => {
            alert(`Something goes wrong: ${error}`);
        }
    );
}

export function ConfirmRegister(){
    return(
        <Page>
            <div class="container text-center w-25">
            <h3>Confirm registration</h3>
            <form onSubmit={handleConfirm}>
                <div class="mb-3">
                <label for="register-confirm-code" class="form-label">Confirmation code</label>
                    <input required type="text" placeholder="confirm-code" id="register-confirm-code" class="form-control"/>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
            </div>
        </Page>
    )
  }
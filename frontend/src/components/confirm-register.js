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
        <div>
            <h3>Confirm registration</h3>
            <form onSubmit={handleConfirm}>
                <input required type="text" placeholder="confirm-code" id="register-confirm-code"/>
                <button type="submit">Submit</button>
            </form>
        </div>
    )
  }
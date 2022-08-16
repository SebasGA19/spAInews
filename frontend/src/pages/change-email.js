import { Page } from "../components/page"
import { change_email } from "../api/change-email";

function handleChangeEmail(event){
    event.preventDefault();
    const password = document.getElementById("user-password").value;
    const newEmail = document.getElementById("user-new-email").value;

    change_email(password,newEmail)
    .then(
        () => {
            alert("Password changed succesfully")
        }
    )
    .catch(
        error => {
            alert(`Error message is ${error}`)
        }
    )
}

export function ChangeEmail(){
    return(
        <Page>
            <div class="container text-center w-25">
            <h3>Change Email</h3>
            <form onSubmit={handleChangeEmail}>
                <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                    <input required type="password" placeholder="password" id="user-password" class="form-control"/>
                <label for="change-email" class="form-label">New email</label>
                    <input required type="text" placeholder="new-email" id="user-new-email" class="form-control"/>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
            </div>
        </Page>
    )
}
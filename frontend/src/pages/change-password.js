import { Page } from "../components/page"
import { change_password } from "../api/change-password";


function handleChangePassword(){
    const oldPassword = document.getElementById("user-old-password").value;
    const newPassword = document.getElementById("user-new-password").value;

    change_password(oldPassword,newPassword)
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

export function ChangePassword(){
    return(
        <Page>
            <div class="container text-center w-25">
            <h3>Change Password</h3>
            <form onSubmit={handleChangePassword}>
                <div class="mb-3">
                <label for="change-password" class="form-label">Old password</label>
                    <input required type="text" placeholder="old-password" id="user-old-password" class="form-control"/>
                <label for="change-password" class="form-label">Old password</label>
                    <input required type="text" placeholder="new-password" id="user-new-password" class="form-control"/>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
            </div>
        </Page>
    )
}
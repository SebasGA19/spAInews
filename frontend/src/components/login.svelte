<script>
import {session, sessionUsername} from '../stores/session'
import {login} from '../api/login'


let username, password;
let errorMessage, successMessage;

function handleLogin() {
    login(username, password)
    .then(
        s => {
            sessionUsername.set(username);
            session.set(s);
            location.reload();
        }
    ).catch(
        error => {
            errorMessage = error;
        }
    );
}
</script>


<!-- Button trigger modal -->
<a class="nav-link" data-bs-toggle="modal" href="#" data-bs-target="#loginModal">
    Login
</a>
  
<!-- Modal -->
<div class="modal fade" id="loginModal" tabindex="-1" aria-labelledby="loginModalLabel" aria-hidden="{true}">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="loginModalLabel">Login</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                {#if errorMessage}
                <div class="alert alert-danger" role="alert">
                    {errorMessage}
                </div>
                {/if}
                <form on:submit|preventDefault={handleLogin}>
                    <div class="mb-3">
                        <label for="login-username" class="form-label">Username</label>
                        <input required type="text" placeholder="username" class="form-control" bind:value="{username}"/>
                    </div>
                    <div class="mb-3">
                        <label for="login-password" class="form-label">Password</label>
                        <input required type="password" placeholder="password" class="form-control" bind:value="{password}"/>
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                    <br>
                    <a href="/forgotpassword">Forgot password?</a>
                </form>
            </div>
        </div>
    </div>
</div>
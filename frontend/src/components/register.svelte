<script>
    import {register} from '../api/register'
    
    let username, email, password;
    let errorMessage, successMessage;
    
    function handleRegister() {
        register(username, email, password)
        .then(
            s => {
                successMessage="Check your email for authentication";
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
    <a class="nav-link" data-bs-toggle="modal" href="#" data-bs-target="#registerModal">
        Register
    </a>

    <!-- Modal -->
    <div class="modal fade" id="registerModal" tabindex="-1" aria-labelledby="registerModalLabel" aria-hidden="{true}">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="registerModalLabel">Register</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    {#if errorMessage}
                    <div class="alert alert-danger" role="alert">
                        {errorMessage}
                    </div>
                    {/if}
                    {#if successMessage}
                    <div class="alert alert-success" role="alert">
                        {successMessage}
                    </div>
                    {/if}
                    <form on:submit|preventDefault={handleRegister}>
                        <div class="mb-3">
                            <label for="register-username" class="form-label">Username</label>
                            <input required type="text" placeholder="username" class="form-control" bind:value="{username}"/>
                        </div>
                        <div class="mb-3">
                            <label for="register-email" class="form-label">email</label>
                            <input required type="text" placeholder="email" class="form-control" bind:value="{email}"/>
                        </div>
                        <div class="mb-3">
                            <label for="register-password" class="form-label">Password</label>
                            <input required type="password" placeholder="password" class="form-control" bind:value="{password}"/>
                        </div>
                        <button type="submit" class="btn btn-primary">Submit</button>
                    </form>
                </div>
            </div>
        </div>
    </div>
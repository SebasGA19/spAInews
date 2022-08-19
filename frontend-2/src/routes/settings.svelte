<script>
	import Page from '../components/page.svelte';
	import { change_password } from "../api/change-password";
  import { change_email } from "../api/change-email";
  import { session } from '../stores/session';
  import Footer from '../components/footer.svelte';

  let password, newEmail;
  let oldPassword, newPassword;
  let errorMessage, successMessage;
  let errorMessage_pass, successMessage_pass;

  function handleChangePassword(){
    change_password(oldPassword, newPassword, $session)
    .then(
      () => {
        successMessage_pass = "Check your email for authentication";
      }
    ).catch(
      error => {
        errorMessage_pass = error;
      }
    )
  }

  function handleChangeEmail(){
    change_email(password, newEmail, $session)
    .then(
      () => {
        successMessage = "Check your email for authentication";
      }
    ).catch( 
        error => {
        errorMessage = error;
        }
    )
  }
</script>

<Page>
    <div class="container text-center">
        <div class="accordion accordion-flush" id="accordionFlushExample">
            <div class="accordion-item">
              <h2 class="accordion-header" id="flush-headingOne">
                <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#flush-collapseOne" aria-expanded="false" aria-controls="flush-collapseOne">
                  Change Email
                </button>
              </h2>
              <div id="flush-collapseOne" class="accordion-collapse collapse" aria-labelledby="flush-headingOne" data-bs-parent="#accordionFlushExample">
                <div class="accordion-body">
                  <div class="container text-center w-25">
                    <h3>Change Email</h3>
                    {#if errorMessage}
                    <div class="alert alert-danger" role="alert">
                        {errorMessage}
                    </div>
                    {/if}
                    {#if successMessage}
                    <div class="alert alert-danger" role="alert">
                      {successMessage}
                    </div>
                    {/if}
                    <form on:submit|preventDefault={handleChangeEmail}>
                        <div class="mb-3">
                        <label for="password" class="form-label">Password</label>
                            <input required type="password" placeholder="password" bind:value="{password}" class="form-control"/>
                        <label for="change-email" class="form-label">New email</label>
                            <input required type="email" placeholder="new-email" bind:value="{newEmail}" class="form-control"/>
                        </div>
                        <button type="submit" class="btn btn-primary">Submit</button>
                    </form>
                    </div>
                </div>
              </div>
            </div>
            <div class="accordion-item">
              <h2 class="accordion-header" id="flush-headingTwo">
                <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#flush-collapseTwo" aria-expanded="false" aria-controls="flush-collapseTwo">
                  Change Password
                </button>
              </h2>
              <div id="flush-collapseTwo" class="accordion-collapse collapse" aria-labelledby="flush-headingTwo" data-bs-parent="#accordionFlushExample">
                <div class="accordion-body">
                  <div class="container text-center w-25">
                    <h3>Change Password</h3>
                    {#if errorMessage}
                    <div class="alert alert-danger" role="alert">
                        {errorMessage}
                    </div>
                    {/if}
                    {#if successMessage}
                    <div class="alert alert-danger" role="alert">
                      {successMessage}
                    </div>
                    {/if}
                    <form on:submit|preventDefault={handleChangePassword}>
                        <div class="mb-3">
                        <label for="change-password" class="form-label">Old password</label>
                            <input required type="text" placeholder="old-password" bind:value="{oldPassword}" class="form-control"/>
                        <label for="change-password" class="form-label">New password</label>
                            <input required type="text" placeholder="new-password" bind:value="{newPassword}" class="form-control"/>
                        </div>
                        <button type="submit" class="btn btn-primary">Submit</button>
                    </form>
                    </div>
                </div>
              </div>
            </div>
          </div>
    </div>
</Page>
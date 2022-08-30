<script>
	import { forgot_password_confirm } from '../api/forgot-password';
	import Page from '../components/page.svelte';
	let  code, newPassword;
	let errorMessage, successMessage;

	function handleForgotPasswordConfirm() {
		forgot_password_confirm(newPassword,code)
			.then((s) => {
				successMessage = 'Password changed succesfully';
				
			})
			.catch((error) => {
				errorMessage = error;
			});
	}
</script>

<Page>
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
   
	<div class="container text-center" style="max-width: 50vw;">
		<form on:submit|preventDefault={handleForgotPasswordConfirm}>
			<div class="mb-3">
				<label for="forgot-password" class="form-label">Code</label>
				<input
					required
					type="text"
					placeholder="Confirmation code"
					class="form-control"
					bind:value={code}
				/>
			</div>
            <div class="mb-3">
				<label for="forgot-password" class="form-label">New Password</label>
				<input
					required
					type="password"
					placeholder="password"
					class="form-control"
					bind:value={newPassword}
				/>
			</div>
			<button type="submit" class="btn btn-primary">Submit</button>
		</form>
	</div>
</Page>

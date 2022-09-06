<script>
	import { forgot_password } from '../api/forgot-password';
	import Page from '../components/page.svelte';
	let  email;
	let errorMessage, successMessage;

	function handleForgotPassword() {
		forgot_password(email)
			.then((s) => {
				successMessage = 'Check your email for authentication';
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
   
	<div class="container text-center">
		<form on:submit|preventDefault={handleForgotPassword}>
			<div class="mb-3">
				<label for="forgot-password" class="form-label">email</label>
				<input
					required
					type="text"
					placeholder="email"
					class="form-control"
					bind:value={email}
				/>
			</div>
			<button type="submit" class="btn btn-primary">Submit</button>
		</form>
	</div>
</Page>

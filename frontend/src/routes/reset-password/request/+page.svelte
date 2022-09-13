<script lang="ts">
	import { requestPasswordReset } from '../../../api/forgot-password';

	let email: string;
	let errorMessage: string | undefined;
	let successMessage: string | undefined;

	function handleForgotPassword() {
		requestPasswordReset(email)
			.then((s) => {
				errorMessage = undefined;
				successMessage = 'Check your email';
			})
			.catch((error) => {
				successMessage = undefined;
				errorMessage = error;
			});
	}
</script>

<div class="container" style="max-width: 500px;">
	<form on:submit|preventDefault={handleForgotPassword}>
		<h3>Request password reset</h3>
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
		<div class="mb-3">
			<input
				required
				type="text"
				placeholder="your.email@mail.provider.com"
				class="form-control"
				bind:value={email}
			/>
		</div>
		<button type="submit" class="btn btn-primary">Submit</button>
	</form>
</div>

<script lang="ts">
	import { resetPassword } from '../../../../api/forgot-password';

	export let data: any;
	let newPassword: string;
	let errorMessage: string | undefined;
	let successMessage: string | undefined;

	function handleForgotPasswordConfirm() {
		resetPassword(newPassword, data.code)
			.then(() => {
				successMessage = 'Password changed succesfully';
			})
			.catch((error) => {
				errorMessage = error;
			});
	}
</script>

<div class="container" style="max-width: 500px;">
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
	<form on:submit|preventDefault={handleForgotPasswordConfirm}>
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

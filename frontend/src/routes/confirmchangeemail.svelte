<script>
	import { change_email, confirm_change_email } from '../api/change-email';
	import Page from '../components/page.svelte';

	let confirmCode;
	let successMessage, errorMessage;

	function handleConfirm() {
		confirm_change_email(confirmCode)
			.then(() => {
				successMessage = 'Email changed succesfully';
			})
			.catch((error) => {
				errorMessage = error;
			});
	}
</script>

<Page>
	<div class="container text-center" style="max-width: 50vw;">
		<h3>Confirm Change email</h3>
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
		<form on:submit|preventDefault={handleConfirm}>
			<div class="mb-3">
				<label for="confirm-code" class="form-label">Confirmation code</label>
				<input
					required
					type="text"
					placeholder="confirm-code"
					bind:value={confirmCode}
					class="form-control"
				/>
			</div>
			<button type="submit" class="btn btn-primary">Submit</button>
		</form>
	</div>
</Page>

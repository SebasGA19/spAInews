<script>
	import { confirmRegister } from '../api/confirm';
	import Page from '../components/page.svelte';

	let confirmCode;
	let successMessage, errorMessage;

	function handleConfirm() {
		confirmRegister(confirmCode)
			.then(() => {
				successMessage = 'Password changed succesfully';
			})
			.catch((error) => {
				errorMessage = error;
			});
	}
</script>

<Page>
	<div class="container text-center" style="max-width: 50vw;">
		<h3>Confirm Register</h3>
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

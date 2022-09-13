<script lang="ts">
	import { register } from '../api/register';

	let username: string, email: string, password: string;
	let errorMessage: string | undefined, successMessage: string | undefined;

	function handleRegister() {
		register(username, email, password)
			.then((s) => {
				successMessage = 'Check your email for authentication';
				errorMessage = undefined;
			})
			.catch((error) => {
				successMessage = undefined;
				errorMessage = error;
			});
	}
</script>

<!-- Button trigger modal -->
<a class="nav-link" data-bs-toggle="modal" href="#" data-bs-target="#registerModal"> Register </a>

<!-- Modal -->
<div
	class="modal fade"
	id="registerModal"
	tabindex="-1"
	aria-labelledby="registerModalLabel"
	aria-hidden={true}
>
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<h5 class="modal-title" id="registerModalLabel">Register</h5>
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" />
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
					<div class="mb-1">
						<label for="register-username" class="form-label">Username</label>
						<input
							required
							type="text"
							placeholder="username"
							class="form-control"
							bind:value={username}
						/>
					</div>
					<div class="mb-1">
						<label for="register-email" class="form-label">email</label>
						<input
							required
							type="text"
							placeholder="email"
							class="form-control"
							bind:value={email}
						/>
					</div>
					<div class="mb-3">
						<label for="register-password" class="form-label">Password</label>
						<input
							required
							type="password"
							placeholder="password"
							class="form-control"
							bind:value={password}
						/>
					</div>
					<div class="text-center w-100">
						<button type="submit" class="btn btn-primary">Submit</button>
					</div>
				</form>
			</div>
		</div>
	</div>
</div>

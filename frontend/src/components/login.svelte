<script lang="ts">
	import { session, sessionUsername } from '../stores/session';
	import { login } from '../api/login';
	import { onMount } from 'svelte';

	let username: string;
	let password: string;
	let errorMessage: string | undefined;
	let loginModal: any;
	onMount(() => {
		loginModal = new bootstrap.Modal('#loginModal', {
			keyboard: false
		});
	});

	function handleLogin() {
		login(username, password)
			.then((s) => {
				sessionUsername.set(username);
				session.set(s);
				location.reload();
				loginModal.hide();
			})
			.catch((error) => {
				errorMessage = error;
			});
	}
</script>

<!-- Button trigger modal -->
<a class="nav-link" data-bs-toggle="modal" href="#" data-bs-target="#loginModal"> Login </a>

<!-- Modal -->
<div
	class="modal fade"
	id="loginModal"
	tabindex="-1"
	aria-labelledby="loginModalLabel"
	aria-hidden={true}
>
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<h5 class="modal-title" id="loginModalLabel">Login</h5>
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" />
			</div>
			<div class="modal-body">
				{#if errorMessage}
					<div class="alert alert-danger" role="alert">
						{errorMessage}
					</div>
				{/if}
				<form on:submit|preventDefault={handleLogin}>
					<div class="mb-1">
						<label for="login-username" class="form-label">Username</label>
						<input
							required
							type="text"
							placeholder="username"
							class="form-control"
							bind:value={username}
						/>
					</div>
					<div class="mb-1">
						<label for="login-password" class="form-label">Password</label>
						<input
							required
							type="password"
							placeholder="password"
							class="form-control"
							bind:value={password}
						/>
					</div>
					<div class="mb-3">
						<a href="/reset-password/request" class="text-reset" on:click={loginModal.hide()}
							>Forgot password?</a
						>
					</div>
					<div class="text-center w-100">
						<button type="submit" class="btn btn-primary">Submit</button>
					</div>
				</form>
			</div>
		</div>
	</div>
</div>

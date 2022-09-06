<script lang="ts">
	import { session, sessionUsername } from '../stores/session';
	import { darkMode } from '../stores/dark-mode';
	import { logout } from '../api/login';

	import Login from './login.svelte';
	import Register from './register.svelte';

	let darkModeChecked: boolean = $darkMode === 'dark';

	$: {
		darkModeChecked;
		$darkMode = darkModeChecked ? 'dark' : 'light';
	}

	function handle_logout() {
		logout($session)
			.then(() => {
				$sessionUsername = '';
				$session = '';
			})
			.catch(() => {
				$sessionUsername = '';
				$session = '';
			});
	}
</script>

<nav class="navbar navbar-expand-lg {darkModeChecked ? 'navbar-dark bg-dark' : 'bg-light'}">
	<div class="container-fluid">
		<button
			class="navbar-toggler"
			type="button"
			data-bs-toggle="collapse"
			data-bs-target="#navbarNav"
			aria-controls="navbarNav"
			aria-expanded="false"
			aria-label="Toggle navigation"
		>
			<span class="navbar-toggler-icon" />
		</button>
		<div class="collapse navbar-collapse" id="navbarNav">
			<ul class="navbar-nav me-auto mb-2 mb-lg-0">
				<li class="nav-item">
					<a href="/" class="nav-link">spAInews</a>
				</li>
			</ul>
			<div class="d-flex" role="search">
				<ul class="navbar-nav me-auto mb-2 mb-lg-0">
					{#if $session !== ''}
						<div class="dropdown-center">
							<button
								class="btn btn-primary dropdown-toggle"
								type="button"
								data-bs-toggle="dropdown"
								aria-expanded="false"
							>
								{$sessionUsername}
							</button>
							<ul class="dropdown-menu dropdown-menu-end">
                                <li><a class="dropdown-item" href="#" on:click={() => {darkModeChecked = !darkModeChecked}}>Theme: {darkModeChecked ? "Dark" : "Light"}</a></li>
								<li><a class="dropdown-item" href="/settings">Settings</a></li>
								<li><a class="dropdown-item" href="#" on:click={handle_logout}>Logout</a></li>
							</ul>
						</div>
					{:else}
						<li class="nav-item">
							<Login />
						</li>
						<li class="nav-item">
							<Register />
						</li>
					{/if}
				</ul>
			</div>
		</div>
	</div>
</nav>
